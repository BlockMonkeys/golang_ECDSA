package main

import (
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"ecdsa/base58Encrypt"
	"ecdsa/genKey"
	"ecdsa/ripemdEncrypt"
	"ecdsa/shaEncrypt"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

type Node struct {
	pubKey					[]byte
	privateKey				*ecdsa.PrivateKey
	bitcoinAddress 			string
	signHash				[]uint8
	signature				[]uint8
}

func main(){
	var N Node
	Init(N)
}

func (N *Node)CreateBitcoinAddress() string{
	//Generate Key Pair
	pubKey, privateKey := genKey.GenKey()
	N.pubKey = pubKey
	N.privateKey = privateKey

	//Make Version
	version := []byte{0}

	//Generate CheckSum
	checkSum := shaEncrypt.GenCheckSum(N.pubKey)

	//Generate PubKeyHash
	pubKeyHash := ripemdEncrypt.GenPubKeyHash(N.pubKey)

	//BASE58 Encrypt To make Bitcoin Address
	encode := base58Encrypt.Base58Encryption(version, pubKeyHash, checkSum)
	N.bitcoinAddress = encode

	return encode
}

func (N *Node)Signing(privateKey *ecdsa.PrivateKey, data string) ([]uint8, []uint8){
	var h hash.Hash

	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, data)

	signHash := h.Sum(nil)

	r, s, seer := ecdsa.Sign(rand.Reader, privateKey, signHash)
	if seer != nil {
		fmt.Println(seer)
		os.Exit(1)
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	N.signHash = signHash
	N.signature = signature

	return signHash, signature
}

func (N *Node)Verification(privateKey *ecdsa.PrivateKey, signHash []uint8, signature []uint8) bool{
	r := MakeByteToBigint(signature[:32])
	s := MakeByteToBigint(signature[32:])

	//Public Key가 []byte Type이어서, ecdsa.PublicKey 타입으로 타입변환
	var pubK ecdsa.PublicKey
	pubK = privateKey.PublicKey

	verifyStatus := ecdsa.Verify(&pubK, signHash, r, s)
	return verifyStatus
}

func MakeByteToBigint(data []byte) *big.Int{
	result := new(big.Int)
	result.SetBytes(data)
	return result
}

func Init(N Node){
	var user int
	var data string
	var flag bool
	flag = true

	fmt.Println("👋🏼 환영합니다. 사용법 안내입니다. 👋🏼")
	fmt.Println("📌 도움말 다시보기 : 0")
	fmt.Println("📌 비트코인 지갑생성 : 1")
	fmt.Println("📌 데이터 사인하기 : 2")
	fmt.Println("📌 데이터 검증하기 : 3")
	fmt.Println("📌 종료하기 : 4")

	for flag{
		fmt.Scanln(&user)

		switch user {
		case 0:
			fmt.Println("👋🏼 환영합니다. 사용법 안내입니다. 👋🏼")
			fmt.Println("📌 도움말 다시보기 : 0")
			fmt.Println("📌 비트코인 지갑생성 : 1")
			fmt.Println("📌 데이터 사인하기 : 2")
			fmt.Println("📌 데이터 검증하기 : 3")
			fmt.Println("📌 종료하기 : 4")
		case 1:
			//비트코인 지갑생성.
			bitcoinAddress := N.CreateBitcoinAddress()
			fmt.Println("👜 당신의 비트코인 지갑이 생성되었습니다. 지갑 주소는 :", bitcoinAddress)
		case 2:
			//데이터 사인하기.
			fmt.Println("서명을 원하는 데이터를 입력해주세요")
			fmt.Scan(&data)
			if N.privateKey != nil {
				signHash, signature := N.Signing(N.privateKey, data)
				reSignHash := fmt.Sprintf("%x", signHash)
				reSig := fmt.Sprintf("%x", signature)
				fmt.Println("✅ 서명이 완료된 데이터 해싱값은 다음과 같습니다. :", reSignHash)
				fmt.Println("✅ 서명이 완료되었습니다. 서명은 다음과 같습니다. :", reSig)
			} else {
				fmt.Println("❌ 지갑을 먼저 생성해야 서명이 가능합니다. ❌")
			}

		case 3:
			//데이터 검증하기
			fmt.Println("데이터에 대해 검증을 시작합니다. 🤔")
			result := N.Verification(N.privateKey, N.signHash, N.signature)
			if result == true {
				fmt.Println("✅ 데이터의 검증이 일치합니다.")
				flag = false
			} else {
				fmt.Println("❌ 데이터 검증결과 일치하지 않습니다. ❌")
			}
		default:
			fmt.Println("항목에 원하시는 선택지가 없습니다. 다시 입력해주세요.")
		}
	}
}