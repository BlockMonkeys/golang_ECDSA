package shaEncrypt

import(
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
)

func ShaEnCryption(key ecdsa.PublicKey) []uint8 {
	data := elliptic.Marshal(key, key.X, key.Y) //pubKey 형변환 []byte로.
	hash := sha256.New()
	hash.Write([]byte(data))
	hashedPubkey := hash.Sum(data)
	return hashedPubkey
}