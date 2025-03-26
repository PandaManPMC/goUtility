package ethWal

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"goUtility/util"
	"golang.org/x/crypto/sha3"
	"regexp"
	"strings"
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

func ImportWallet(mnemonic string, index int) (privateKey *ecdsa.PrivateKey, address string, err error) {
	wallet, err := util.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		return nil, "", fmt.Errorf("failed to import mnemonic: %v", err)
	}

	privateKey, err = util.GetInstanceByHDWalletUtil().WalletPrivateKey(wallet, index)
	if nil != err {
		return nil, "", fmt.Errorf("failed to WalletPrivateKey: %v", err)
	}

	address = PrivateKeyToAddressETH(privateKey)
	return
}

// IValidAddress 校验 ETH 地址格式（0x 开头，40 个十六进制字符）
func IValidAddress(address string) bool {
	address = strings.ToLower(address)
	match, _ := regexp.MatchString("^0x[a-fA-F0-9]{40}$", address)
	return match
}
