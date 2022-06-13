package util

import "encoding/binary"

// MarshalUInt64 returns a byte array of length 8 in little-endian encoding
func MarshalUInt64(n uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, n)

	return buf
}

func UnmarshalUInt64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

func Build(Vb []byte) [][]byte {
	V := make([][]byte, len(Vb)/TsxSize)
	for i := 0; i < len(Vb); i += TsxSize {
		V[i/TsxSize] = Vb[i : i+TsxSize]
	}

	return V
}
