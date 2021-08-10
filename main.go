package main

import (
	"crypto/ecdsa"
	"ecdsa/base58Encrypt"
	"ecdsa/genKey"
	"ecdsa/ripemdEncrypt"
	"ecdsa/shaEncrypt"
	"fmt"
)

type BitcoinAddress struct {
	version			[]byte
	pubKeyHash		[]byte
	checkSum		[]byte
}

type Node struct {
	bitcoinAddress 			*BitcoinAddress
	signature				string
	pubKey					[]byte
	privateKey				ecdsa.PrivateKey
}

func main(){
	var B BitcoinAddress
	B.version = []byte{0}

	//Generate Key Pair
	pubKey, _ := genKey.GenKey()

	//Generate CheckSum
	checkSum := shaEncrypt.GenCheckSum(pubKey)
	B.checkSum = checkSum

	//Generate PubKeyHash
	pubKeyHash := ripemdEncrypt.GenPubKeyHash(pubKey)
	B.pubKeyHash = pubKeyHash

	encode := base58Encrypt.Base58Encryption(B.version, B.pubKeyHash, B.checkSum)
	fmt.Println("RESULT BASE58 :", encode)



	////해시 암호화 & PubKey를 암호화 한 뒤, 다시 SHA-256 암호화를 함으로써 CheckSum 생성함.
	//hashedPublicKey := shaEncrypt.ShaEncryption(pubKey) // [32]byte -> []byte
	//checkSum := shaEncrypt.ShaEncryptionAgain(hashedPublicKey)
	//b.checkSum = fmt.Sprintf("%x", checkSum)
	//
	//rPubKey := ripemdEncrypt.RipemdEncryption(hashedPublicKey)
	//b.pubKeyHash = rPubKey
	//
	//address := base58Encrypt.Base58Encryption(b.version, b.pubKeyHash, b.checkSum)
	//fmt.Println()
	//fmt.Println("문자열 어드레스: ", address)
}