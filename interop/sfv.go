package interop

import (
	"fmt"
	"io"

	"github.com/willglynn/hatt/manifest"
)

type sfv struct{}

func (x sfv) Name() string {
	return "sfv"
}

func (x sfv) Description() string {
	return "filename and CRC32"
}

func (x sfv) Export(m *manifest.Manifest, w io.Writer) error {
	fmt.Fprintf(w, "; SFV written by hatt\n;\n")

	for path, file := range m.Files {
		if file.Hashes.CRC32 != nil {
			if _, err := fmt.Fprintf(w, "%s %08X\n", path, *file.Hashes.CRC32); err != nil {
				return err
			}
		}
	}

	return nil
}
