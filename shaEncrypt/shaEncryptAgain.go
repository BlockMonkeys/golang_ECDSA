package shaEncrypt

import (
	"crypto/sha256"
)

func ShaEnCryptionAgain(key []uint8) []uint8 {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashedPubkey := hash.Sum(key)
	return hashedPubkey
}