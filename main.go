package main

import (
	"ecdsa/base58Encrypt"
	"ecdsa/genKey"
	"ecdsa/ripemdEncrypt"
	"ecdsa/shaEncrypt"
	"fmt"
)

func main(){
	//** GenKey를 통해 키페어를 생성하는데,
	//main함수가 실행될 때마다, 계속 재실행되어 keyPair가 바뀌는 문제가 존재.
	pubKey, _ := genKey.GenKey()
	version := []uint8{1}

	//해시 암호화 & PubKey를 암호화 한 뒤, 다시 SHA-256 암호화를 함으로써 CheckSum 생성함.
	hashed := shaEncrypt.ShaEnCryption(pubKey)
	checkSum := shaEncrypt.ShaEnCryptionAgain(hashed)

	//RIPEMD
	pubKeyHash := ripemdEncrypt.RipemdEncryption(hashed)

	bitAddress := base58Encrypt.Base58Encryption(version, pubKeyHash, checkSum)
	fmt.Println("BitCoinAddress :", bitAddress)
}