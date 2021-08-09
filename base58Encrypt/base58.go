package base58Encrypt

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func Base58Encryption(ver []uint8, pubKeyHash []uint8, checkSum []uint8) string {

	data := append(ver[:], pubKeyHash[:]...)
	data = append(data[:], checkSum[:]...)
	fmt.Println("data :", data)

	encoded := base58.Encode(data)
	return encoded
}
