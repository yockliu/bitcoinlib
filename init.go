package bitcoinlib

import "encoding/binary"

// UvarintToBytes change uint64 to 9 max bytes uvarint
func UvarintToBytes(value uint64) ([]byte, int) {
	buf := make([]byte, 9)
	bitsSize := binary.PutUvarint(buf, value)
	result := make([]byte, bitsSize)
	copy(buf, result)
	return result, bitsSize
}
