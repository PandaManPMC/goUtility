package util

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
)

type CryptoWallet struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

type CryptoAddressGenerate struct{}

var cryptoAddressGenerateInstance CryptoAddressGenerate

func GetInstanceByCryptoAddressGenerate() CryptoAddressGenerate {
	return cryptoAddressGenerateInstance
}

func (g CryptoAddressGenerate) GenerateWallet(chain string) (*CryptoWallet, error) {
	switch strings.ToUpper(chain) {
	case "ETH", "ETHEREUM", "BSC", "BNB", "POL", "POLYGON":
		return g.generateETH()

	case "TRX", "TRON":
		return g.generateTRON()

	case "SOL", "SOLANA":
		return g.generateSOL()

	case "DOGE", "DOGECOIN":
		return g.generateBTCFamily(0x1e)

	case "LTC", "LITECOIN":
		return g.generateBTCFamily(0x30)

	case "RVN", "RAVENCOIN":
		return g.generateBTCFamily(0x3c)

	default:
		return nil, errors.New("unsupported chain")
	}
}

//////////////////////////////////////////////////////
// ETH / BSC / POLYGON
//////////////////////////////////////////////////////

func (g CryptoAddressGenerate) generateETH() (*CryptoWallet, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	pubBytes := crypto.FromECDSAPub(&priv.PublicKey)

	return &CryptoWallet{
		PrivateKey: hex.EncodeToString(
			crypto.FromECDSA(priv),
		),
		PublicKey: hex.EncodeToString(pubBytes),
		Address:   crypto.PubkeyToAddress(priv.PublicKey).Hex(),
	}, nil
}

//////////////////////////////////////////////////////
// TRON
//////////////////////////////////////////////////////

func (g CryptoAddressGenerate) generateTRON() (*CryptoWallet, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	pubBytes := crypto.FromECDSAPub(&priv.PublicKey)

	hash := crypto.Keccak256(pubBytes[1:])

	payload := append(
		[]byte{0x41},
		hash[12:]...,
	)

	checksum := g.doubleSHA256(payload)[:4]

	address := g.base58Encode(
		append(payload, checksum...),
	)

	return &CryptoWallet{
		PrivateKey: hex.EncodeToString(
			crypto.FromECDSA(priv),
		),
		PublicKey: hex.EncodeToString(pubBytes),
		Address:   address,
	}, nil
}

//////////////////////////////////////////////////////
// SOLANA
//////////////////////////////////////////////////////

func (g CryptoAddressGenerate) generateSOL() (*CryptoWallet, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &CryptoWallet{
		PrivateKey: hex.EncodeToString(priv.Seed()),
		PublicKey:  g.base58Encode(pub),
		Address:    g.base58Encode(pub),
	}, nil
}

//////////////////////////////////////////////////////
// DOGE / LTC / RVN
//////////////////////////////////////////////////////

func (g CryptoAddressGenerate) generateBTCFamily(version byte) (*CryptoWallet, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	pubBytes := crypto.FromECDSAPub(&priv.PublicKey)

	sha := sha256.Sum256(pubBytes)

	h := ripemd160.New()
	h.Write(sha[:])

	pubHash := h.Sum(nil)

	payload := append(
		[]byte{version},
		pubHash...,
	)

	checksum := g.doubleSHA256(payload)[:4]

	address := g.base58Encode(
		append(payload, checksum...),
	)

	return &CryptoWallet{
		PrivateKey: hex.EncodeToString(
			crypto.FromECDSA(priv),
		),
		PublicKey: hex.EncodeToString(pubBytes),
		Address:   address,
	}, nil
}

//////////////////////////////////////////////////////
// Base58
//////////////////////////////////////////////////////

func (g CryptoAddressGenerate) base58Encode(input []byte) string {
	if len(input) == 0 {
		return ""
	}

	x := make([]byte, len(input))
	copy(x, input)

	var answer []byte

	for len(x) > 0 {
		var remainder int
		var quotient []byte

		for _, b := range x {
			acc := int(b) + remainder*256

			digit := acc / 58
			remainder = acc % 58

			if len(quotient) > 0 || digit != 0 {
				quotient = append(
					quotient,
					byte(digit),
				)
			}
		}

		answer = append(
			answer,
			base58Alphabet[remainder],
		)

		x = quotient
	}

	for _, b := range input {
		if b != 0 {
			break
		}
		answer = append(answer, '1')
	}

	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}

	return string(answer)
}

//////////////////////////////////////////////////////
// Double SHA256
//////////////////////////////////////////////////////

func (g CryptoAddressGenerate) doubleSHA256(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}
