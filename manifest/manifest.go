package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"

	"io"
	"io/ioutil"
)

type Manifest struct {
	Files map[string]File
}

func New() *Manifest {
	return &Manifest{
		Files: make(map[string]File),
	}
}

func Read(r io.Reader) (*Manifest, error) {
	// we can use bufio to peek here later, dispatching on magic numbers
	// for now: assume JSON

	decoder := json.NewDecoder(r)
	m := &Manifest{}
	if err := decoder.Decode(m); err != nil {
		return nil, err
	} else {
		return m, nil
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

func (m Manifest) Write(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(m)
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
