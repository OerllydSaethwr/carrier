package util

import "encoding/binary"

const (
	Suint64 uint8 = 8
	Suint32 uint8 = 4
)

// MarshalUInt64 returns a byte array of length 8 in little-endian encoding
func MarshalUInt64(n uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, n)

	return buf
}

func UnmarshalUInt64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

// MarshalUInt32 returns a byte array of length 8 in little-endian encoding
func MarshalUInt32(n uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, n)

	return buf
}

func UnmarshalUInt32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func Build(Vb []byte) [][]byte {
	V := make([][]byte, len(Vb)/TsxSize)
	for i := 0; i < len(Vb); i += TsxSize {
		V[i/TsxSize] = Vb[i : i+TsxSize]
	}

	return V
}

func Frame(buf []byte) []byte {
	// 0. Calculate the size of the whole struct, encode it as 8 bytes - uint64 and put it at the beginning of the whole buffer
	framedBuf := make([]byte, 0)
	lenBuf2 := uint32(len(buf))
	framedBuf = append(framedBuf, MarshalUInt32(lenBuf2)...)
	framedBuf = append(framedBuf, buf...)

	return framedBuf
}
