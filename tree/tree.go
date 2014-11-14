package tree

import (
	"io"
	"log"
	"os"
	"time"
)

type Entry struct {
	Path    string
	Size    int64
	ModTime time.Time
	Error   error
}

type treeWalker struct {
	c chan Entry
}

func (tw treeWalker) concat(a, b string) string {
	return a + "/" + b
}

func (tw treeWalker) handleFileInfo(path string, fi os.FileInfo) {
	mode := fi.Mode()
	modeType := mode & os.ModeType

	if (modeType & os.ModeDir) != 0 {
		// this is a directory; recurse
		tw.walkDirectory(path)

	} else if modeType == 0 {
		// this is a regular file
		tw.c <- Entry{
			Path:    path,
			Size:    fi.Size(),
			ModTime: fi.ModTime(),
		}

	} else {
		// not a regular file; ignore
		log.Printf("%q: not a regular file, ignoring")
	}
}

func (tw treeWalker) walk(path string) {
	if fi, err := os.Lstat(path); err != nil {
		// emit an error entry
		tw.c <- Entry{
			Path:  path,
			Error: err,
		}
		return
	} else {
		// not an error; handle normally
		// this invokes walkDirectory() as needed
		tw.handleFileInfo(path, fi)
	}

	// done walking
	// signal by closing the channel
	close(tw.c)
}

func (tw treeWalker) walkDirectory(path string) {
	// open the directory
	if f, err := os.Open(path); err != nil {
		// fail
		tw.c <- Entry{
			Path:  path,
			Error: err,
		}
		return
	} else {
		// enumerate all the things
		for {
			fis, err := f.Readdir(64)
			if err == io.EOF {
				break
			}

			if err != nil {
				tw.c <- Entry{
					Path:  path,
					Error: err,
				}
				break
			}

			if len(fis) == 0 {
				break
			}

			for _, fi := range fis {
				fullPath := tw.concat(path, fi.Name())
				tw.handleFileInfo(fullPath, fi)
			}
		}

		f.Close()
	}
}

func Walk(path string) <-chan Entry {
	tw := &treeWalker{
		c: make(chan Entry),
	}

	go tw.walk(path)

	return tw.c
}
