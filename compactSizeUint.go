package bitcoinlib

// CompactSizeUint compactsize unsigned integers
// https://bitcoin.org/en/developer-reference#compactsize-unsigned-integers
type CompactSizeUint struct {
	value uint64
}

// Size get the bytes count
func (csu *CompactSizeUint) Size() int {
	return 8
}

// Bytes get the bytes buffer
func (csu *CompactSizeUint) Bytes() []byte {
	return make([]byte, 9)
}
