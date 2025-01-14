package util

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// pathDrive 路径
const pathDrive = "m/44'/60'/0'/0/%d"

type hDWalletUtil struct {
}

var hDWalletUtilInstance hDWalletUtil

func GetInstanceByHDWalletUtil() *hDWalletUtil {
	return &hDWalletUtilInstance
}

// GenerateMnemonicBy12 生成助记词
func (that *hDWalletUtil) GenerateMnemonicBy12() (string, error) {
	// 128 位熵（生成 12 个单词的助记词）
	return that.GenerateMnemonicBits(128)
}

// GenerateMnemonicBy15 生成助记词
func (that *hDWalletUtil) GenerateMnemonicBy15() (string, error) {
	return that.GenerateMnemonicBits(160)
}

// GenerateMnemonicBy18 生成助记词
func (that *hDWalletUtil) GenerateMnemonicBy18() (string, error) {
	return that.GenerateMnemonicBits(192)
}

// GenerateMnemonicBy21 生成助记词
func (that *hDWalletUtil) GenerateMnemonicBy21() (string, error) {
	return that.GenerateMnemonicBits(224)
}

// GenerateMnemonicBy24 生成助记词
func (that *hDWalletUtil) GenerateMnemonicBy24() (string, error) {
	return that.GenerateMnemonicBits(256)
}

// GenerateMnemonicBits
//
//	熵位数 (bits)	校验位数 (bits)	总位数 (bits)	助记词单词数
//	128					4				132				12
//	160					5				165				15
//	192					6				198				18
//	224					7				231				21
//	256					8				264				24
func (that *hDWalletUtil) GenerateMnemonicBits(bits int) (string, error) {
	entropy, err := hdwallet.NewEntropy(bits)
	if nil != err {
		return "", fmt.Errorf("failed to generate entropy: %v", err)
	}

	mnemonic, err := hdwallet.NewMnemonicFromEntropy(entropy)
	if nil != err {
		return "", fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return mnemonic, nil
}

func (that *hDWalletUtil) ImportWalletFromMnemonic(mnemonic string) (*hdwallet.Wallet, error) {
	return hdwallet.NewFromMnemonic(mnemonic)
}

func (that *hDWalletUtil) WalletPrivateKey(wallet *hdwallet.Wallet, index int) (privateKey *ecdsa.PrivateKey, err error) {
	// 动态生成路径
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf(pathDrive, index))
	var account accounts.Account
	account, err = wallet.Derive(path, true)
	if nil != err {
		return nil, err
	}
	return wallet.PrivateKey(account)
}

func (that *hDWalletUtil) PriKeyToHexString(key *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(key))
}

// LoadWalletByPrivateKey 读取钱包
func (that *hDWalletUtil) LoadWalletByPrivateKey(privateKeyHexString string) (*ecdsa.PrivateKey, error) {
	privateKey, e := crypto.HexToECDSA(privateKeyHexString)
	if nil != e {
		return nil, e
	}
	return privateKey, e
}
