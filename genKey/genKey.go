package genKey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

func GenKey()(pub ecdsa.PublicKey,private *ecdsa.PrivateKey){
	//ECDSA KEYPAIR 생성
	pubkeyCurve := elliptic.P256()
	privateKey := new(ecdsa.PrivateKey)

	fmt.Printf("뭐냐 %x :",pubkeyCurve)

	privateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var pubkey ecdsa.PublicKey
	pubkey = privateKey.PublicKey

	return pubkey, privateKey
}


