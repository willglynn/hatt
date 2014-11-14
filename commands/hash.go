package commands

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/willglynn/hatt/manifest"
	"github.com/willglynn/hatt/tree"
	"github.com/willglynn/hatt/workgroup"
)

var HashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash files, creating (or updating) a manifest",
	Long: `This is the core function of hatt: hash all the things!

The 'hash' command generates a manifest, which describes all the files it
hashed. Manifests can be exported in a number of formats for comparison by
other tools, or they can be compared by hatt directly.`,

	Run: hash.run,
}

type hashOperation struct {
	manifestPath string
	threads      int
	checksum     bool

	m *manifest.Manifest
}

var hash hashOperation

func (hash *hashOperation) shouldQueue(entry tree.Entry) bool {
	if manifestEntry, exists := hash.m.Files[entry.Path]; exists {
		// path exists in manifest

		if hash.checksum {
			// force re-checking
			return true
		}

		if manifestEntry.Size == entry.Size && manifestEntry.ModTime == entry.ModTime {
			// size and mtime match
			// don't bother re-checking
			return false
		}

		// default: check
		return true
	} else {
		// path does not exist in manifest
		return true
	}
}

func (hash *hashOperation) run(cmd *cobra.Command, args []string) {
	if hash.manifestPath == "" {
		cmd.Usage()
		log.Fatal("manifest file must be specified")
		return
	}

	if hash.threads <= 0 {
		hash.threads = 1
	}

	// ensure we actually use the requested number of threads
	if runtime.GOMAXPROCS(0) < hash.threads {
		runtime.GOMAXPROCS(hash.threads)
	}

	// open the manifest, ignorning not found errors
	if m, err := manifest.ReadFromFile(hash.manifestPath); err != nil && !os.IsNotExist(err) {
		log.Fatalf("error reading manifest file %q: %v", hash.manifestPath, err)
	} else {
		hash.m = m
	}

	// create a new manifest if needed
	if hash.m == nil {
		hash.m = manifest.New()
	}

	// trap signals to close down gracefully (i.e. preserving work)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, syscall.SIGTERM)

	// set up the hash workgroup
	wg := workgroup.New(hash.threads)
	wgInput := wg.Input()
	wgOutput := wg.Output()

	// start walking the tree
	walk := tree.Walk(".")

	// do all the things!
	interrupted := false
	pending := make([]string, 0, 20)
loop:
	for {
		// close wgInput if we're done walking and there's nothing else to send
		if walk == nil && wgInput != nil && len(pending) == 0 {
			close(wgInput)
			wgInput = nil
		}

		// selectively enable reading from the "files" channel based on the length
		// of our pending queue
		recvWalk := walk
		if len(pending) >= 20 {
			recvWalk = nil
		}

		// selectively enable sending to the wgOutput channel if we have something to send
		sendWgInput := wgInput
		firstPending := ""
		if len(pending) == 0 {
			sendWgInput = nil
		} else {
			firstPending = pending[0]
		}

		select {
		case <-interrupt:
			interrupted = true
			log.Println("interrupted; aborting...")
			break loop

		case entry, ok := <-recvWalk:
			if !ok {
				// finished walking
				walk = nil

			} else {
				// consider the walk entry
				if entry.Error != nil {
					log.Printf("error walking %q: %v", entry.Path, entry.Error)
				} else if hash.shouldQueue(entry) {
					pending = append(pending, entry.Path)
				}
			}

		case sendWgInput <- firstPending:
			// sent to workgroup
			pending = pending[1:]

		case hashedFile, ok := <-wgOutput:
			if !ok {
				// workgroup is done
				break loop
			} else {
				if hashedFile.Error != nil {
					// failed to hash this file
					log.Printf("error hashing %q: %v", hashedFile.Path, hashedFile.Error)
				} else {
					// got a hashed file
					// add it to the manifest
					log.Printf("hashed %q (%d bytes)", hashedFile.Path, hashedFile.File.Size)
					hash.m.Files[hashedFile.Path] = *hashedFile.File
				}
			}
		}
	}

	// write the manifest
	if err := hash.m.WriteToFile(hash.manifestPath); err != nil {
		log.Fatalf("unable to write manifest to %q: %v\n", hash.manifestPath, err)
	}

	if interrupted {
		os.Exit(1)
	}
}

func init() {
	HashCmd.PersistentFlags().StringVarP(&hash.manifestPath, "manifest", "m", "", "path to manifest file")
	HashCmd.PersistentFlags().IntVarP(&hash.threads, "threads", "t", 0, "number of files to hash simultaneously")
	HashCmd.PersistentFlags().BoolVarP(&hash.checksum, "checksum", "c", false, "force checksum, even when size+mtime match")
}
