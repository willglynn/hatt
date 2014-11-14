package manifest

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"io"
)

type HashSet struct {
	MD5     []byte
	SHA1    []byte
	SHA256  []byte
	CRC32   uint32
	Adler32 uint32
}

type writableHashSet struct {
	writer            io.Writer
	md5, sha1, sha256 hash.Hash
	crc32, adler32    hash.Hash32
}

func newWritableHashSet() *writableHashSet {
	whs := &writableHashSet{
		md5:     md5.New(),
		sha1:    sha1.New(),
		sha256:  sha256.New(),
		crc32:   crc32.NewIEEE(),
		adler32: adler32.New(),
	}

	whs.writer = io.MultiWriter(whs.md5, whs.sha1, whs.sha256, whs.crc32, whs.adler32)

	return whs
}

func (whs *writableHashSet) Write(p []byte) (int, error) {
	return whs.writer.Write(p)
}

func (whs *writableHashSet) Sum() HashSet {
	return HashSet{
		MD5:     whs.md5.Sum(make([]byte, 0, 16)),
		SHA1:    whs.sha1.Sum(make([]byte, 0, 20)),
		SHA256:  whs.sha256.Sum(make([]byte, 0, 32)),
		CRC32:   whs.crc32.Sum32(),
		Adler32: whs.adler32.Sum32(),
	}
}
