package bitcoinlib

/*
 * KeyPair = ecdsa(elliptic.P256())
 * Address = Version(0x00) + Public Key Hash + Checksum
 */

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/itchyny/base58-go"
	"golang.org/x/crypto/ripemd160"
)

const version = 0x00

// KeyPair public key and private key
type KeyPair struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// GenKeyPair new private key and public key
func GenKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(fmt.Sprintf("GenKeyPaire ecdsa generate key error: %s", err))
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

// HashPubKey get the hash of the public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(publicSHA256[:])
	if err != nil {
		panic(fmt.Sprintf("HashPubKey ripemd160Hasher Write error: %s", err))
	}
	publicRIPEMD160 := ripemd160Hasher.Sum(nil)

	return publicRIPEMD160
}

// GetAddressFromPubKeyHash get address from public key hash, reverse GetPubKeyHashFromAddress
func GetAddressFromPubKeyHash(pubKeyHash []byte) []byte {
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	encode := base58.BitcoinEncoding
	address, err := encode.Encode(fullPayload)
	if err != nil {
		panic(fmt.Sprintf("GetAddressFromPubKeyHash base58 encode error: %s", err))
	}

	return address
}

// GetPubKeyHashFromAddress get public key hash from address, reverse GetAddressFromPubKeyHash
func GetPubKeyHashFromAddress(address []byte) []byte {
	encode := base58.BitcoinEncoding
	versionedPayload, err := encode.Decode(address)
	if err != nil {
		panic(fmt.Sprintf("GetPubKeyHashFromAddress base58 decode error: %s", err))
	}
	pubKeyHash := versionedPayload[1 : len(versionedPayload)-4]
	return pubKeyHash
}

func checksum(payload []byte) []byte {
	firstSha := sha256.Sum256(payload)
	sencondSha := sha256.Sum256(firstSha[:])

	return sencondSha[:4]
}
