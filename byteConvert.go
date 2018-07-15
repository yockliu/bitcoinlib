package bitcoinlib

import "encoding/binary"

// Uint32ToBytes convert uint64 to byte slice
func Uint32ToBytes(value uint32) []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, value)
	return result
}

// Uint64ToBytes convert uint64 to byte slice
func Uint64ToBytes(value uint64) []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, value)
	return result
}
