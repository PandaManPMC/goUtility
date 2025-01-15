package ethWal

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func PubKeyToAddressETH(publicKey ecdsa.PublicKey) string {

	publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes)
	hashed := hash.Sum(nil)

	address := hashed[len(hashed)-20:]

	return fmt.Sprintf("0x%s", hex.EncodeToString(address))
}

func PrivateKeyToAddressETH(privateKey *ecdsa.PrivateKey) string {
	return PubKeyToAddressETH(privateKey.PublicKey)
}
