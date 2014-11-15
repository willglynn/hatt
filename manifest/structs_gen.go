package manifest

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"github.com/philhofer/msgp/msgp"
)


// MarshalMsg implements the msgp.Marshaler interface
func (z *Manifest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 1)

	o = msgp.AppendString(o, "Files")

	o = msgp.AppendMapHeader(o, uint32(len(z.Files)))
	for xvk, bzg := range z.Files {
		o = msgp.AppendString(o, xvk)

		o, err = bzg.MarshalMsg(o)
		if err != nil {
			return
		}

	}

	return
}
// UnmarshalMsg unmarshals a Manifest from MessagePack, returning any extra bytes
// and any errors encountered
func (z *Manifest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Files":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Files == nil && msz > 0 {
				z.Files = make(map[string]File, int(msz))
			} else if len(z.Files) > 0 {
				for key, _ := range z.Files {
					delete(z.Files, key)
				}
			}
			for inx := uint32(0); inx < msz; inx++ {
				var xvk string
				var bzg File
				xvk, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}

				bts, err = bzg.UnmarshalMsg(bts)

				if err != nil {
					return
				}

				z.Files[xvk] = bzg
			}

		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// Msgsize implements the msgp.Sizer interface
func (z *Manifest) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 5

	s += msgp.MapHeaderSize
	if z.Files != nil {
		for xvk, bzg := range z.Files {
			_ = bzg
			s += msgp.StringPrefixSize + len(xvk)

			s += bzg.Msgsize()

		}
	}

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *Manifest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Files":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Files == nil && msz > 0 {
				z.Files = make(map[string]File, int(msz))
			} else if len(z.Files) > 0 {
				for key, _ := range z.Files {
					delete(z.Files, key)
				}
			}
			for inx := uint32(0); inx < msz; inx++ {
				var xvk string
				var bzg File
				xvk, err = dc.ReadString()
				if err != nil {
					return
				}

				err = bzg.DecodeMsg(dc)

				if err != nil {
					return
				}

				z.Files[xvk] = bzg
			}

		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}

	return
}

// EncodeMsg implements the msgp.Encodable interface
func (z *Manifest) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(1)
	if err != nil {
		return
	}

	err = en.WriteString("Files")
	if err != nil {
		return
	}

	err = en.WriteMapHeader(uint32(len(z.Files)))
	if err != nil {
		return
	}

	for xvk, bzg := range z.Files {
		err = en.WriteString(xvk)
		if err != nil {
			return
		}

		err = bzg.EncodeMsg(en)

		if err != nil {
			return
		}

	}

	return
}

// MarshalMsg implements the msgp.Marshaler interface
func (z *File) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 3)

	o = msgp.AppendString(o, "Size")

	o = msgp.AppendInt64(o, z.Size)

	o = msgp.AppendString(o, "ModTime")

	o = msgp.AppendTime(o, z.ModTime)

	o = msgp.AppendString(o, "Hashes")

	o, err = z.Hashes.MarshalMsg(o)
	if err != nil {
		return
	}

	return
}
// UnmarshalMsg unmarshals a File from MessagePack, returning any extra bytes
// and any errors encountered
func (z *File) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Size":

			z.Size, bts, err = msgp.ReadInt64Bytes(bts)

			if err != nil {
				return
			}

		case "ModTime":

			z.ModTime, bts, err = msgp.ReadTimeBytes(bts)

			if err != nil {
				return
			}

		case "Hashes":

			bts, err = z.Hashes.UnmarshalMsg(bts)

			if err != nil {
				return
			}

		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// Msgsize implements the msgp.Sizer interface
func (z *File) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 4

	s += msgp.Int64Size
	s += msgp.StringPrefixSize + 7

	s += msgp.TimeSize
	s += msgp.StringPrefixSize + 6

	s += z.Hashes.Msgsize()

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *File) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Size":

			z.Size, err = dc.ReadInt64()

			if err != nil {
				return
			}

		case "ModTime":

			z.ModTime, err = dc.ReadTime()

			if err != nil {
				return
			}

		case "Hashes":

			err = z.Hashes.DecodeMsg(dc)

			if err != nil {
				return
			}

		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}

	return
}

// EncodeMsg implements the msgp.Encodable interface
func (z *File) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}

	err = en.WriteString("Size")
	if err != nil {
		return
	}

	err = en.WriteInt64(z.Size)

	if err != nil {
		return
	}

	err = en.WriteString("ModTime")
	if err != nil {
		return
	}

	err = en.WriteTime(z.ModTime)

	if err != nil {
		return
	}

	err = en.WriteString("Hashes")
	if err != nil {
		return
	}

	err = z.Hashes.EncodeMsg(en)

	if err != nil {
		return
	}

	return
}

// MarshalMsg implements the msgp.Marshaler interface
func (z *HashSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 5)

	o = msgp.AppendString(o, "MD5")

	o = msgp.AppendBytes(o, z.MD5)

	o = msgp.AppendString(o, "SHA1")

	o = msgp.AppendBytes(o, z.SHA1)

	o = msgp.AppendString(o, "SHA256")

	o = msgp.AppendBytes(o, z.SHA256)

	o = msgp.AppendString(o, "CRC32")

	if z.CRC32 == nil {
		o = msgp.AppendNil(o)
	} else {

		o = msgp.AppendUint32(o, *z.CRC32)

	}

	o = msgp.AppendString(o, "Adler32")

	if z.Adler32 == nil {
		o = msgp.AppendNil(o)
	} else {

		o = msgp.AppendUint32(o, *z.Adler32)

	}

	return
}
// UnmarshalMsg unmarshals a HashSet from MessagePack, returning any extra bytes
// and any errors encountered
func (z *HashSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "MD5":

			z.MD5, bts, err = msgp.ReadBytesBytes(bts, z.MD5)

			if err != nil {
				return
			}

		case "SHA1":

			z.SHA1, bts, err = msgp.ReadBytesBytes(bts, z.SHA1)

			if err != nil {
				return
			}

		case "SHA256":

			z.SHA256, bts, err = msgp.ReadBytesBytes(bts, z.SHA256)

			if err != nil {
				return
			}

		case "CRC32":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.CRC32 = nil
			} else {
				if z.CRC32 == nil {
					z.CRC32 = new(uint32)
				}

				*z.CRC32, bts, err = msgp.ReadUint32Bytes(bts)

				if err != nil {
					return
				}

			}

		case "Adler32":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Adler32 = nil
			} else {
				if z.Adler32 == nil {
					z.Adler32 = new(uint32)
				}

				*z.Adler32, bts, err = msgp.ReadUint32Bytes(bts)

				if err != nil {
					return
				}

			}

		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// Msgsize implements the msgp.Sizer interface
func (z *HashSet) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 3

	s += msgp.BytesPrefixSize + len(z.MD5)
	s += msgp.StringPrefixSize + 4

	s += msgp.BytesPrefixSize + len(z.SHA1)
	s += msgp.StringPrefixSize + 6

	s += msgp.BytesPrefixSize + len(z.SHA256)
	s += msgp.StringPrefixSize + 5

	if z.CRC32 == nil {
		s += msgp.NilSize
	} else {

		s += msgp.Uint32Size

	}
	s += msgp.StringPrefixSize + 7

	if z.Adler32 == nil {
		s += msgp.NilSize
	} else {

		s += msgp.Uint32Size

	}

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *HashSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "MD5":

			z.MD5, err = dc.ReadBytes(z.MD5)

			if err != nil {
				return
			}

		case "SHA1":

			z.SHA1, err = dc.ReadBytes(z.SHA1)

			if err != nil {
				return
			}

		case "SHA256":

			z.SHA256, err = dc.ReadBytes(z.SHA256)

			if err != nil {
				return
			}

		case "CRC32":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.CRC32 = nil
			} else {
				if z.CRC32 == nil {
					z.CRC32 = new(uint32)
				}

				*z.CRC32, err = dc.ReadUint32()

				if err != nil {
					return
				}

			}

		case "Adler32":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Adler32 = nil
			} else {
				if z.Adler32 == nil {
					z.Adler32 = new(uint32)
				}

				*z.Adler32, err = dc.ReadUint32()

				if err != nil {
					return
				}

			}

		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}

	return
}

// EncodeMsg implements the msgp.Encodable interface
func (z *HashSet) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(5)
	if err != nil {
		return
	}

	err = en.WriteString("MD5")
	if err != nil {
		return
	}

	err = en.WriteBytes(z.MD5)

	if err != nil {
		return
	}

	err = en.WriteString("SHA1")
	if err != nil {
		return
	}

	err = en.WriteBytes(z.SHA1)

	if err != nil {
		return
	}

	err = en.WriteString("SHA256")
	if err != nil {
		return
	}

	err = en.WriteBytes(z.SHA256)

	if err != nil {
		return
	}

	err = en.WriteString("CRC32")
	if err != nil {
		return
	}

	if z.CRC32 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {

		err = en.WriteUint32(*z.CRC32)

		if err != nil {
			return
		}

	}

	err = en.WriteString("Adler32")
	if err != nil {
		return
	}

	if z.Adler32 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {

		err = en.WriteUint32(*z.Adler32)

		if err != nil {
			return
		}

	}

	return
}
