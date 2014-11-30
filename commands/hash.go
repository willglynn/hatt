package commands

import (
	"fmt"
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
other tools, or they can be compared by hatt directly.

'hash' will:
  * create a new manifest if needed
  * add files to the manifest for each file found on disk
  * update files that are already in the manifest iff the size+mtime differ
    (modified by --checksum/-c)
  * silently ignore files that exist in the manifest but not on disk
    (modified by --missing=log or --missing=remove)
  * use one thread for hashing
    (modified by --threads/-t)
  * use nontrivial amounts of CPU time to calculate many different hashes
  	(modified by --32-bits-only/-c)
  * write out the work-in-progess manifest when interrupted (i.e. ^C)`,

	Run: hash.run,
}

// what do we do with missing things?
type hashMissingAction int

const (
	hashMissingIgnore hashMissingAction = iota
	hashMissingLog
	hashMissingRemove
)

func (hma hashMissingAction) String() string {
	switch hma {
	case hashMissingIgnore:
		return "ignore"
	case hashMissingLog:
		return "log"
	case hashMissingRemove:
		return "remove"
	default:
		return ""
	}
}

func (hma *hashMissingAction) Set(value string) error {
	switch value {
	case "ignore":
		*hma = hashMissingIgnore
	case "log":
		*hma = hashMissingLog
	case "remove":
		*hma = hashMissingRemove

	default:
		return fmt.Errorf("expected 'ignore', 'log', 'remove'; got %q", value)
	}
	return nil
}

func (hma *hashMissingAction) Type() string {
	return "hash missing action"
}

type hashOperation struct {
	// configuration
	manifestPath     string
	threads          int
	checksum         bool
	thirtyTwoBitOnly bool
	missingAction    hashMissingAction

	// runtime state
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
			// don't bother re-checking, but update the manifest indicating that we found it
			manifestEntry.Seen = true
			hash.m.Files[entry.Path] = manifestEntry
			return false
		}

		// default: check
		return true
	} else {
		// path does not exist in manifest
		// hash it
		return true
	}
}

func (hash *hashOperation) manifestHashOptions() manifest.HashOptions {
	hashOpts := manifest.HashOptions{}

	if hash.thirtyTwoBitOnly {
		hashOpts.DisableMD5 = true
		hashOpts.DisableSHA1 = true
		hashOpts.DisableSHA256 = true
	}

	return hashOpts
}

func (hash *hashOperation) handleUnseenEntries() {
	// bail early if there's nothing to do
	if hash.missingAction == hashMissingIgnore {
		return
	}

	// walk the manifest, removing any entries that aren't Seen
	for key, entry := range hash.m.Files {
		if !entry.Seen {
			if hash.missingAction == hashMissingLog {
				log.Printf("%q does not exist locally but exists in manifest", key)
			} else {
				delete(hash.m.Files, key)
				log.Printf("removed %q from manifest", key)
			}
		}
	}
}

func (hash *hashOperation) hashAllTheThings() bool {
	// trap signals to close down gracefully (i.e. preserving work)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, syscall.SIGTERM)

	// set up the hash workgroup
	wg := workgroup.New(hash.threads, hash.manifestHashOptions())
	wgInput := wg.Input()
	wgOutput := wg.Output()

	// start walking the tree
	walk := tree.Walk(".")

	// do all the things!
	pending := make([]string, 0, 20)
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
			log.Println("interrupted; aborting...")
			return true

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
				return false
			} else {
				if hashedFile.Error != nil {
					// failed to hash this file
					log.Printf("error hashing %q: %v", hashedFile.Path, hashedFile.Error)
				} else {
					// got a hashed file
					// add it to the manifest
					log.Printf("hashed %q (%d bytes)", hashedFile.Path, hashedFile.File.Size)

					manifestEntry := *hashedFile.File
					manifestEntry.Seen = true
					hash.m.Files[hashedFile.Path] = manifestEntry
				}
			}
		}
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

	// run the main loop
	interrupted := hash.hashAllTheThings()

	// handle unseen entries iff we finished normally
	if !interrupted {
		hash.handleUnseenEntries()
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
	HashCmd.PersistentFlags().IntVarP(&hash.threads, "threads", "t", 1, "number of files to hash simultaneously")
	HashCmd.PersistentFlags().BoolVarP(&hash.checksum, "checksum", "c", false, "force checksum, even when size+mtime match")
	HashCmd.PersistentFlags().BoolVar(&hash.thirtyTwoBitOnly, "32-bits-only", false, "compute only 32-bit checksums (fast)")
	HashCmd.PersistentFlags().Var(&hash.missingAction, "missing", "action to perform for missing files (ignore, log, remove)")
}
