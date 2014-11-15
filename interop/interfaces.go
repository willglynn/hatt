package interop

import (
	"bufio"
	"io"

	"github.com/willglynn/hatt/manifest"
)

type Format interface {
	Name() string
	Description() string
}

type Importer interface {
	Import(*bufio.Reader) (*manifest.Manifest, error)
}

type Exporter interface {
	Export(*manifest.Manifest, io.Writer) error
}
