package manifest

import (
	"fmt"
	"io"
	"os"
	"time"
)

type File struct {
	Size    int64
	ModTime time.Time
	Hashes  HashSet
}

func NewFileFromPath(path string) (*File, error) {
	// open the file
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		defer f.Close()

		// stat the file to determine what we expect
		file := File{}
		if stat, err := f.Stat(); err != nil {
			return nil, err
		} else {
			file.Size = stat.Size()
			file.ModTime = stat.ModTime()
		}

		// calculate hashes for everything
		// TODO: progress indicator
		whs := newWritableHashSet()
		n, err := io.Copy(whs, f)
		if err != nil {
			return nil, err
		}

		// see if we read a different number of bits
		if file.Size != n {
			return nil, fmt.Errorf("%q: expected %d bytes, read %d", path, file.Size, n)
		}

		// finalize the hashset
		file.Hashes = whs.Sum()

		return &file, nil
	}
}
