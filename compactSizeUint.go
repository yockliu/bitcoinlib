package bitcoinlib

// CompactSizeUint compactsize unsigned integers
// https://bitcoin.org/en/developer-reference#compactsize-unsigned-integers
type CompactSizeUint struct {
	Value uint64
}

// Size get the bytes count
func (csu *CompactSizeUint) Size() int {
	if csu.Value <= 252 {
		return 1
	}
	if csu.Value <= 0xffff {
		return 3
	}
	if csu.Value <= 0xffffffff {
		return 5
	}
	return 9
}

// Bytes get the bytes buffer
func (csu *CompactSizeUint) Bytes() []byte {
	return make([]byte, 9)
}
