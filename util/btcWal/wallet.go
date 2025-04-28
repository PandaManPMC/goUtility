package btcWal

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"goUtility/util"
	"golang.org/x/crypto/ripemd160"
)

// 通过路径派生钱包地址
// path 最大 m/0/0/4294967295 uint32
// 生成首个地址的路径，BIP44 规定了多币种支持的路径格式为 m / purpose' / coin_type' / account' / change / address_index。
// 首个地址的路径构成：
// purpose：根据 BIP44，purpose 固定为 44'，表示这是 BIP44 标准定义的路径。
// coin_type：每种加密货币对应一个 coin_type，比如：
// 比特币（BTC）：0'  莱特币（LTC）：2'  以太坊（ETH）：60'
// account：通常是 0'，代表第一个账户。如果你希望派生多个账户，可以调整这个值。
// change：用于区分接收地址和找零地址，0 表示接收地址，1 表示找零地址。
// address_index：用于区分同一账户下的不同地址，通常从 0 开始递增。
// 生成第一个地址的路径： 路径格式：m / 44' / coin_type' / 0' / 0 / 0
// m：表示主私钥。
// 44'：BIP44 标准的 purpose。
// coin_type'：特定币种的 coin_type，例如比特币是 0'。
// 0'：账户索引（通常为 0 表示第一个账户）。
// 0：表示接收地址（找零地址为 1）。
// 0：表示第一个地址。

// 生成公钥哈希
func pubKeyHash(pubKey []byte) []byte {
	sha := sha256.New()
	sha.Write(pubKey)
	shaSum := sha.Sum(nil)
	ripe := ripemd160.New()
	ripe.Write(shaSum)
	return ripe.Sum(nil)
}

const (
	BTCAddress  byte = 0x00
	LTCAddress  byte = 0x30
	DogeAddress byte = 0x1E
)

func GenerateAddressByBTC(privateKey *ecdsa.PrivateKey) (string, error) {
	return GenerateAddress(privateKey, BTCAddress)
}

func GenerateAddressByLTC(privateKey *ecdsa.PrivateKey) (string, error) {
	return GenerateAddress(privateKey, LTCAddress)
}

func GenerateAddressByDoge(privateKey *ecdsa.PrivateKey) (string, error) {
	return GenerateAddress(privateKey, DogeAddress)
}

// GenerateAddress 生成地址
// 要根据 ecdsa.PrivateKey 生成比特币（BTC）、莱特币（LTC）和狗狗币（DOGE）的地址，您需要执行以下步骤：
// 获取公钥：从私钥生成公钥。
// 计算公钥哈希：对公钥进行 SHA-256 哈希，然后对结果进行 RIPEMD-160 哈希。
// 生成地址：
// BTC：使用版本字节 0x00，并进行双重 SHA-256 校验和。
// LTC：使用版本字节 0x30，并进行双重 SHA-256 校验和。
// DOGE：使用版本字节 0x1E，并进行双重 SHA-256 校验和。
// Base58Check 编码：将上述结果进行 Base58Check 编码，得到最终地址。
func GenerateAddress(privateKey *ecdsa.PrivateKey, coinType byte) (string, error) {
	// 获取公钥
	pubKey := append([]byte{0x04}, privateKey.PublicKey.X.Bytes()...)
	pubKey = append(pubKey, privateKey.PublicKey.Y.Bytes()...)

	// 计算公钥哈希
	pubKHash := pubKeyHash(pubKey)

	// 添加版本字节
	addressBytes := append([]byte{coinType}, pubKHash...)

	// 计算校验和
	checksum := sha256.Sum256(addressBytes)
	checksum = sha256.Sum256(checksum[:])

	// 获取前四个字节作为校验和
	var checksum4 [4]byte
	copy(checksum4[:], checksum[:4])

	// 拼接地址
	addressBytes = append(addressBytes, checksum4[:]...)
	address := base58.Encode(addressBytes)

	return address, nil
}

// ImportWallet
// coinType 比特币（BTC）：0 莱特币（LTC）：2  狗狗币（DOGE）：3
func ImportWallet(mnemonic string, coinType util.HDCoinType, index int) (privateKey *ecdsa.PrivateKey, address string, err error) {
	hdWallet, err := util.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		return nil, "", err
	}

	privateKey, err = util.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(hdWallet, coinType, index)
	if nil != err {
		return nil, "", err
	}

	var coinByte byte = 0x00
	switch coinType {
	case 0:
		coinByte = 0x00
	case 2:
		coinByte = 0x30
	case 3:
		coinByte = 0x1E
	default:
		return nil, "", errors.New("invalid coinType")
	}
	address, err = GenerateAddress(privateKey, coinByte)
	return privateKey, address, err
}

// IsValidBTCAddress 验证比特币地址是否合法
func IsValidBTCAddress(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	return err == nil
}

// IsValidLTCAddress 验证莱特币地址是否合法
func IsValidLTCAddress(address string) bool {
	// 莱特币的主网参数需要自定义，以下是示例参数
	ltcParams := &chaincfg.Params{
		PubKeyHashAddrID: 0x30, // 莱特币地址前缀为 L（0x30）
		ScriptHashAddrID: 0x32, // 莱特币脚本地址前缀为 M（0x32）
		Bech32HRPSegwit:  "ltc",
	}
	_, err := btcutil.DecodeAddress(address, ltcParams)
	return err == nil
}

// IsValidDOGEAddress 验证狗狗币地址是否合法
func IsValidDOGEAddress(address string) bool {
	// 狗狗币的主网参数需要自定义，以下是示例参数
	dogeParams := &chaincfg.Params{
		PubKeyHashAddrID: 0x1e, // 狗狗币地址前缀为 D（0x1e）
		ScriptHashAddrID: 0x16, // 狗狗币脚本地址前缀为 9（0x16）
		Bech32HRPSegwit:  "doge",
	}
	_, err := btcutil.DecodeAddress(address, dogeParams)
	return err == nil
}
