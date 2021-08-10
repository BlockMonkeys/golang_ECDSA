package base58Encrypt

import "github.com/btcsuite/btcutil/base58"

func Base58Encryption(ver, pubKeyHash, checkSum []byte) string {
	verAddPubKeyHash := append(ver, pubKeyHash...)
	data := append(verAddPubKeyHash, checkSum...)
	encoded := base58.Encode(data)
	return encoded
}
