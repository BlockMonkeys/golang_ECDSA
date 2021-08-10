package shaEncrypt

import (
	"crypto/sha256"
)

func GenCheckSum(publicKey []byte) []byte{
	shaOne := sha256.Sum256(publicKey)
	shaTwo := sha256.Sum256(shaOne[:])
	return shaTwo[:]
}