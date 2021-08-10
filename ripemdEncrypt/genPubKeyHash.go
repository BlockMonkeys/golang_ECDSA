package ripemdEncrypt

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func GenPubKeyHash(publicKey []byte) []byte{
	publicSha := sha256.Sum256(publicKey)

	ripHash := ripemd160.New()
	_, err := ripHash.Write(publicSha[:])
	if err != nil {
		fmt.Println(err)
	}
	ripPubKey := ripHash.Sum(nil)
	return ripPubKey
}
