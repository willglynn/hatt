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

type HashOptions struct {
	DisableMD5    bool
	DisableSHA1   bool
	DisableSHA256 bool
}

type writableHashSet struct {
	writer            io.Writer
	md5, sha1, sha256 hash.Hash
	crc32, adler32    hash.Hash32
}

func newWritableHashSet(opts HashOptions) *writableHashSet {
	whs := &writableHashSet{
		crc32:   crc32.NewIEEE(),
		adler32: adler32.New(),
	}
	writers := []io.Writer{whs.crc32, whs.adler32}

	if !opts.DisableMD5 {
		whs.md5 = md5.New()
		writers = append(writers, whs.md5)
	}
	if !opts.DisableSHA1 {
		whs.sha1 = sha1.New()
		writers = append(writers, whs.sha1)
	}
	if !opts.DisableSHA256 {
		whs.sha256 = sha256.New()
		writers = append(writers, whs.sha256)
	}

	whs.writer = io.MultiWriter(writers...)

	return whs
}

func (whs *writableHashSet) Write(p []byte) (int, error) {
	return whs.writer.Write(p)
}

func (whs *writableHashSet) Sum() HashSet {
	crc32 := whs.crc32.Sum32()
	adler32 := whs.adler32.Sum32()
	hs := HashSet{
		CRC32:   &crc32,
		Adler32: &adler32,
	}

	if whs.md5 != nil {
		hs.MD5 = whs.md5.Sum(make([]byte, 0, 16))
	}
	if whs.sha1 != nil {
		hs.SHA1 = whs.sha1.Sum(make([]byte, 0, 20))
	}
	if whs.sha256 != nil {
		hs.SHA256 = whs.sha256.Sum(make([]byte, 0, 32))
	}

	return hs
}
