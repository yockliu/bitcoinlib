package bitcoinlib

import (
	"fmt"
	"strings"
)

// HashCode sha256 Code 32 byte
type HashCode [32]byte

// Compare { -1 : < } { 0 : = } { 1 : > }
func (hash *HashCode) Compare(anotherHash *HashCode) int {
	hexStr := fmt.Sprintf("%x", hash)
	anotherHexStr := fmt.Sprintf("%x", anotherHash)
	return strings.Compare(hexStr, anotherHexStr)
}
