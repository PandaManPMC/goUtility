package btcWal

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/crypto"
	"goUtility/util"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"strings"
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
	RVNAddress  byte = 0x3C
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

// GenerateAddress 生成地址(Legacy（P2PKH）地址)
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

func GenerateAddressCompressed(privateKey *ecdsa.PrivateKey, coinType byte) (string, error) {
	// 使用压缩公钥（33 字节，02 或 03 开头）
	var compressedPubKey []byte
	if privateKey.PublicKey.Y.Bit(0) == 0 {
		compressedPubKey = append([]byte{0x02}, privateKey.PublicKey.X.Bytes()...)
	} else {
		compressedPubKey = append([]byte{0x03}, privateKey.PublicKey.X.Bytes()...)
	}

	// 计算公钥哈希（RIPEMD160(SHA256(pubKey)))
	pubKHash := pubKeyHash(compressedPubKey)

	// 添加版本字节（coinType: 比如 0x1E for DOGE, 0x00 for BTC, etc.）
	addressBytes := append([]byte{coinType}, pubKHash...)

	// 计算校验和：双SHA256
	checksum := sha256.Sum256(addressBytes)
	checksum = sha256.Sum256(checksum[:])
	checksum4 := checksum[:4]

	// 拼接地址并 base58 编码
	addressBytes = append(addressBytes, checksum4...)
	address := base58.Encode(addressBytes)

	return address, nil
}

//func GenerateAddressByRVN(privateKey *ecdsa.PrivateKey) (string, error) {
//	return GenerateAddress(privateKey, RVNAddress)
//}
//
//func GenerateAddressCompressedByRVN(privateKey *ecdsa.PrivateKey) (string, error) {
//	return GenerateAddressCompressed(privateKey, RVNAddress)
//}

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
	case util.BTCHDCoinType:
		coinByte = 0x00
	case util.LTCHDCoinType:
		coinByte = 0x30
	case util.DOGEHDCoinType:
		coinByte = 0x1E
	case util.RVNHDCoinType: // RVN
		coinByte = 0x3C
	default:
		return nil, "", errors.New("invalid coinType")
	}
	address, err = GenerateAddressCompressed(privateKey, coinByte)
	return privateKey, address, err
}

func GenerateAddressCompressedByPK(privateKey *ecdsa.PrivateKey, coinType util.HDCoinType) (string, error) {
	var coinByte byte = 0x00
	switch coinType {
	case util.BTCHDCoinType:
		coinByte = 0x00
	case util.LTCHDCoinType:
		coinByte = 0x30
	case util.DOGEHDCoinType:
		coinByte = 0x1E
	case util.RVNHDCoinType: // RVN
		coinByte = 0x3C
	default:
		return "", errors.New("invalid coinType")
	}
	address, err := GenerateAddressCompressed(privateKey, coinByte)
	return address, err
}

func GenerateAddressCompressedByPKStr(privateKeyStr string, coinType util.HDCoinType) (*ecdsa.PrivateKey, string, error) {
	privateKey, e := crypto.HexToECDSA(privateKeyStr)
	if nil != e {
		return nil, "", e
	}
	address, err := GenerateAddressCompressedByPK(privateKey, coinType)
	return privateKey, address, err
}

// IsValidBTCAddress 验证比特币地址是否合法
func IsValidBTCAddress(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	return err == nil
}

// IsValidLTCAddress 验证莱特币地址是否合法
func IsValidLTCAddress(address string) bool {
	if ValidateLitecoinAddress(address) {
		return true
	}
	return validateLitecoinBech32(address)
	// 莱特币的主网参数
	//ltcParams := &chaincfg.Params{
	//	PubKeyHashAddrID: 0x30, // 莱特币地址前缀为 L（0x30）
	//	ScriptHashAddrID: 0x32, // 莱特币脚本地址前缀为 M（0x32）
	//	Bech32HRPSegwit:  "ltc",
	//}
	//_, err := btcutil.DecodeAddress(address, ltcParams)
	//return err == nil
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

var base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base58 解码
func base58Decode(input string) ([]byte, error) {
	result := big.NewInt(0)
	base := big.NewInt(58)
	for _, r := range input {
		charIndex := int64(strings.IndexRune(base58Alphabet, r))
		if charIndex < 0 {
			return nil, fmt.Errorf("invalid character: %c", r)
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(charIndex))
	}

	decoded := result.Bytes()

	// 添加前缀 0（对应 '1'）填补长度
	leadingZeros := 0
	for _, r := range input {
		if r != '1' {
			break
		}
		leadingZeros++
	}

	return append(make([]byte, leadingZeros), decoded...), nil
}

// 获取双重 SHA256 校验和前4字节
func checksum(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:4]
}

// ValidateLitecoinAddress 验证 Litecoin 地址合法性
func ValidateLitecoinAddress(address string) bool {
	decoded, err := base58Decode(address)
	if err != nil || len(decoded) < 5 {
		return false
	}

	// 拆分 payload 和校验和
	payload := decoded[:len(decoded)-4]
	checksumBytes := decoded[len(decoded)-4:]

	// 校验校验和
	expectedChecksum := checksum(payload)
	for i := 0; i < 4; i++ {
		if checksumBytes[i] != expectedChecksum[i] {
			return false
		}
	}

	// 校验前缀字节（Version Byte）
	// 主网地址版本号：
	// P2PKH：0x30 (48, 即 L/M 开头)
	// P2SH：0x32 (50, 即 3 开头)
	version := payload[0]
	if version != 0x30 && version != 0x32 {
		return false
	}

	return true
}

var bech32Charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

// Bech32 polymod constant
var bech32Gen = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}

// Expand HRP
func hrpExpand(hrp string) []int {
	var result []int
	for _, c := range hrp {
		result = append(result, int(c>>5))
	}
	result = append(result, 0)
	for _, c := range hrp {
		result = append(result, int(c&31))
	}
	return result
}

// Polymod checksum
func polymod(values []int) int {
	chk := 1
	for _, v := range values {
		top := chk >> 25
		chk = ((chk & 0x1ffffff) << 5) ^ v
		for i := 0; i < 5; i++ {
			if (top>>i)&1 == 1 {
				chk ^= bech32Gen[i]
			}
		}
	}
	return chk
}

// Verify checksum
func verifyBech32Checksum(hrp string, data []int) bool {
	return polymod(append(hrpExpand(hrp), data...)) == 1
}

// Decode Bech32
func decodeBech32(addr string) (string, []int, error) {
	addr = strings.ToLower(addr)
	if len(addr) < 8 || len(addr) > 90 {
		return "", nil, errors.New("invalid length")
	}

	pos := strings.LastIndexByte(addr, '1')
	if pos < 1 || pos+7 > len(addr) {
		return "", nil, errors.New("invalid position for '1'")
	}

	hrp := addr[:pos]
	dataPart := addr[pos+1:]

	var data []int
	for _, c := range dataPart {
		idx := strings.IndexRune(bech32Charset, c)
		if idx < 0 {
			return "", nil, fmt.Errorf("invalid character: %c", c)
		}
		data = append(data, idx)
	}

	if !verifyBech32Checksum(hrp, data) {
		return "", nil, errors.New("invalid checksum")
	}

	return hrp, data[:len(data)-6], nil // data (excluding checksum)
}

// Litecoin Bech32 地址验证器
func validateLitecoinBech32(addr string) bool {
	hrp, _, err := decodeBech32(addr)
	return err == nil && hrp == "ltc"
}
