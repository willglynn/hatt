package interop

import (
	"fmt"
	"io"

	"github.com/willglynn/hatt/manifest"
)

type md5sum struct{}

func (m md5sum) Name() string {
	return "md5sum"
}

func (m md5sum) Description() string {
	return "MD5 and filename, separated by two spaces"
}

func (x md5sum) Export(m *manifest.Manifest, w io.Writer) error {
	for path, file := range m.Files {
		if len(file.Hashes.MD5) == 16 {
			if _, err := fmt.Fprintf(w, "%16x  %s\n", file.Hashes.MD5, path); err != nil {
				return err
			}
		}
	}

	return nil
}
