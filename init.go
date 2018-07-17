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

// Hashable something that can do a Hash method to get the hash code of the struct
type Hashable interface {
	Hash() *HashCode
}

// Serializable serialized to byte array and deserialized from the byte array
type Serializable interface {
	Serialize() []byte
	Deserialize([]byte)
}

// Cell basic type of the block chain
type Cell interface {
	Hashable
	Serializable
}
