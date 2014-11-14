package manifest

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/philhofer/msgp/msgp"

	"io"
	"io/ioutil"
)

func New() *Manifest {
	return &Manifest{
		Files: make(map[string]File),
	}
}

func Read(r io.Reader) (*Manifest, error) {
	br := bufio.NewReader(r)
	prefix, err := br.Peek(16)
	if err != nil {
		return nil, err
	}

	if len(prefix) >= 1 && prefix[0] == '{' {
		// looks like JSON
		return readJSON(br)
	} else if len(prefix) >= 2 && prefix[0] == 0x1f && prefix[1] == 0x8b {
		// looks like gzip
		return readGzip(br)
	} else {
		return nil, errors.New("unable to identify format; is this a hatt manifest?")
	}
}

func readJSON(r io.Reader) (*Manifest, error) {
	m := &Manifest{}
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(m); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func readGzip(r io.Reader) (*Manifest, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	defer gzr.Close()

	switch gzr.Name {
	case "hatt-manifest.msgp.v0":
		m := &Manifest{}
		if err := msgp.Decode(gzr, m); err != nil {
			return nil, err
		} else {
			return m, nil
		}

	default:
		return nil, fmt.Errorf("unhandled gzip filename %q; is this a hatt manifest?", gzr.Name)
	}
}

func (m Manifest) Write(w io.Writer) error {
	if true {
		gzw := gzip.NewWriter(w)
		defer gzw.Close()

		gzw.Name = "hatt-manifest.msgp.v0"
		gzw.ModTime = time.Now()

		return msgp.Encode(gzw, &m)

	} else {
		// old-style JSON
		encoder := json.NewEncoder(w)
		return encoder.Encode(m)
	}
}

func ReadFromFile(filename string) (*Manifest, error) {
	if f, err := os.Open(filename); err != nil {
		return nil, err
	} else {
		defer f.Close()

		if m, err := Read(f); err != nil {
			return nil, err
		} else {
			return m, nil
		}
	}
}

func (m Manifest) WriteToFile(filename string) error {
	// make a temporary file in the same directory with the same basename
	f, err := ioutil.TempFile(filepath.Dir(filename), filepath.Base(filename)+".tmp.")
	if err != nil {
		return err
	}

	// ensure we (try to) close and clean up
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	// write the manifest
	if err := m.Write(f); err != nil {
		return err
	}

	// close the tempfile
	if err := f.Close(); err != nil {
		return err
	}

	// rename the tempfile over the target file
	if err := os.Rename(f.Name(), filename); err != nil {
		return err
	}

	// success
	return nil
}
