package trie

import (
	"encoding/binary"
)

func Uint32ToBytes(buf []byte, num uint32) []byte {
	binary.LittleEndian.PutUint32(buf, num)
	switch {
	case num <= 0xFF:
		return buf[:1]
	case num <= 0xFFFF:
		return buf[:2]
	case num <= 0xFFFFFF:
		return buf[:3]
	}
	return buf
}
