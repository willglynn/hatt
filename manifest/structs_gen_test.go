package manifest

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"testing"
	"bytes"
	"github.com/philhofer/msgp/msgp"
)


func TestManifestMarshalUnmarshal(t *testing.T) {
	v := new(Manifest)
	bts, err := v.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func BenchmarkManifestMarshalMsg(b *testing.B) {
	v := new(Manifest)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkManifestAppendMsg(b *testing.B) {
	v := new(Manifest)
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkManifestUnmarshal(b *testing.B) {
	v := new(Manifest)
	bts, _ := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestManifestEncodeDecode(t *testing.T) {
	v := new(Manifest)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Maxsize() for %v is inaccurate", v)
	}

	vn := new(Manifest)
	err := msgp.Decode(&buf, vn)
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkManifestEncode(b *testing.B) {
	v := new(Manifest)
	en := msgp.NewWriter(msgp.Nowhere)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.EncodeMsg(en)
	}
	en.Flush()
}

func BenchmarkManifestDecode(b *testing.B) {
	v := new(Manifest)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	rd := bytes.NewReader(buf.Bytes())
	dc := msgp.NewReader(rd)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rd.Seek(0, 0)
		err := v.DecodeMsg(dc)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func TestFileMarshalUnmarshal(t *testing.T) {
	v := new(File)
	bts, err := v.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func BenchmarkFileMarshalMsg(b *testing.B) {
	v := new(File)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkFileAppendMsg(b *testing.B) {
	v := new(File)
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkFileUnmarshal(b *testing.B) {
	v := new(File)
	bts, _ := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestFileEncodeDecode(t *testing.T) {
	v := new(File)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Maxsize() for %v is inaccurate", v)
	}

	vn := new(File)
	err := msgp.Decode(&buf, vn)
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkFileEncode(b *testing.B) {
	v := new(File)
	en := msgp.NewWriter(msgp.Nowhere)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.EncodeMsg(en)
	}
	en.Flush()
}

func BenchmarkFileDecode(b *testing.B) {
	v := new(File)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	rd := bytes.NewReader(buf.Bytes())
	dc := msgp.NewReader(rd)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rd.Seek(0, 0)
		err := v.DecodeMsg(dc)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func TestHashSetMarshalUnmarshal(t *testing.T) {
	v := new(HashSet)
	bts, err := v.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func BenchmarkHashSetMarshalMsg(b *testing.B) {
	v := new(HashSet)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkHashSetAppendMsg(b *testing.B) {
	v := new(HashSet)
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkHashSetUnmarshal(b *testing.B) {
	v := new(HashSet)
	bts, _ := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestHashSetEncodeDecode(t *testing.T) {
	v := new(HashSet)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Maxsize() for %v is inaccurate", v)
	}

	vn := new(HashSet)
	err := msgp.Decode(&buf, vn)
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkHashSetEncode(b *testing.B) {
	v := new(HashSet)
	en := msgp.NewWriter(msgp.Nowhere)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.EncodeMsg(en)
	}
	en.Flush()
}

func BenchmarkHashSetDecode(b *testing.B) {
	v := new(HashSet)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	rd := bytes.NewReader(buf.Bytes())
	dc := msgp.NewReader(rd)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rd.Seek(0, 0)
		err := v.DecodeMsg(dc)
		if err != nil {
			b.Fatal(err)
		}
	}
}