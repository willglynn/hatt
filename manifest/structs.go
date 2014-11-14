//go:generate msgp

package manifest

import "time"

type Manifest struct {
	Files map[string]File
}

type File struct {
	Size    int64
	ModTime time.Time
	Hashes  HashSet
}

type HashSet struct {
	MD5     []byte
	SHA1    []byte
	SHA256  []byte
	CRC32   uint32
	Adler32 uint32
}
