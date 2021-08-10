package genKey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

func GenKey()([]byte, ecdsa.PrivateKey){
	//ECDSA KEYPAIR 생성
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return pubKey, *privateKey
}


