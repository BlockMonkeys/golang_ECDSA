package ripemdEncrypt

import (
	"golang.org/x/crypto/ripemd160"
)

func RipemdEncryption(hashed []uint8) []uint8 {
	ripemd := ripemd160.New()
	ripemd.Write([]byte(hashed))
	ripemdHashPubKey := ripemd.Sum(hashed)
	return ripemdHashPubKey
}
