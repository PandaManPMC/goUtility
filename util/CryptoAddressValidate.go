package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
)

type CryptoAddressValidate struct {
}

var cryptoAddressValidateInstance CryptoAddressValidate

func GetInstanceByCryptoAddressValidate() CryptoAddressValidate {
	return cryptoAddressValidateInstance
}

func (that CryptoAddressValidate) ValidateAddress(chain, addr string) bool {
	switch strings.ToUpper(chain) {
	case "BITCOIN":
		return that.ValidateBTC(addr)
	case "LITECOIN":
		return that.ValidateLTC(addr)
	case "DOGECOIN":
		return that.ValidateBase58Check(addr)
	case "RAVENCOIN":
		return that.ValidateBase58Check(addr)
	case "BSC":
		fallthrough
	case "POLYGON":
		fallthrough
	case "ETHEREUM":
		return that.ValidateETH(addr)
	case "SOLANA":
		return that.ValidateSOL(addr)
	case "NANO":
		return that.ValidateNANO(addr)
	case "TRON":
		return that.ValidateTRON(addr)
	default:
		return false
	}
}

/////////////////////////////
// Base58Check
/////////////////////////////

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func (that CryptoAddressValidate) base58Decode(input string) ([]byte, error) {
	result := []byte{}

	for _, r := range input {
		idx := strings.IndexRune(base58Alphabet, r)
		if idx < 0 {
			return nil, errors.New("invalid base58 char")
		}

		carry := idx
		for j := len(result) - 1; j >= 0; j-- {
			carry += int(result[j]) * 58
			result[j] = byte(carry % 256)
			carry /= 256
		}

		for carry > 0 {
			result = append([]byte{byte(carry % 256)}, result...)
			carry /= 256
		}
	}

	for _, r := range input {
		if r != '1' {
			break
		}
		result = append([]byte{0x00}, result...)
	}

	return result, nil
}

func (that CryptoAddressValidate) doubleSHA256(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

func (that CryptoAddressValidate) ValidateBase58Check(addr string) bool {
	decoded, err := that.base58Decode(addr)
	if err != nil || len(decoded) < 4 {
		return false
	}

	payload := decoded[:len(decoded)-4]
	cs := decoded[len(decoded)-4:]

	return bytes.Equal(that.doubleSHA256(payload)[:4], cs)
}

/////////////////////////////
// Bech32 / Bech32m
/////////////////////////////

const (
	bech32Const  = 1
	bech32mConst = 0x2bc830a3
)

var bech32Charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

func (that CryptoAddressValidate) bech32Polymod(values []byte) uint32 {
	generator := []uint32{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
	chk := uint32(1)

	for _, v := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ uint32(v)
		for i := 0; i < 5; i++ {
			if ((top >> i) & 1) == 1 {
				chk ^= generator[i]
			}
		}
	}
	return chk
}

func (that CryptoAddressValidate) bech32Decode(addr string) (string, []byte, error) {
	addr = strings.ToLower(addr)
	pos := strings.LastIndex(addr, "1")
	if pos < 1 {
		return "", nil, errors.New("no separator")
	}

	hrp := addr[:pos]
	data := []byte{}

	for _, c := range addr[pos+1:] {
		idx := strings.IndexRune(bech32Charset, c)
		if idx < 0 {
			return "", nil, errors.New("invalid char")
		}
		data = append(data, byte(idx))
	}

	if len(data) < 6 {
		return "", nil, errors.New("too short")
	}

	return hrp, data, nil
}

func (that CryptoAddressValidate) verifyBech32Checksum(hrp string, data []byte) bool {
	values := append(that.hrpExpand(hrp), data...)
	polymod := that.bech32Polymod(values)

	return polymod == bech32Const || polymod == bech32mConst
}

func (that CryptoAddressValidate) hrpExpand(hrp string) []byte {
	ret := []byte{}
	for _, c := range hrp {
		ret = append(ret, byte(c>>5))
	}
	ret = append(ret, 0)
	for _, c := range hrp {
		ret = append(ret, byte(c&31))
	}
	return ret
}

func (that CryptoAddressValidate) validateBech32(addr, expectHRP string) bool {
	hrp, data, err := that.bech32Decode(addr)
	if err != nil {
		return false
	}
	if hrp != expectHRP {
		return false
	}
	return that.verifyBech32Checksum(hrp, data)
}

/////////////////////////////
// BTC / LTC
/////////////////////////////

func (that CryptoAddressValidate) ValidateBTC(addr string) bool {
	if strings.HasPrefix(addr, "1") || strings.HasPrefix(addr, "3") {
		return that.ValidateBase58Check(addr)
	}
	if strings.HasPrefix(addr, "bc1") {
		return that.validateBech32(addr, "bc")
	}
	return false
}

func (that CryptoAddressValidate) ValidateLTC(addr string) bool {
	if strings.HasPrefix(addr, "L") ||
		strings.HasPrefix(addr, "M") ||
		strings.HasPrefix(addr, "3") {
		return that.ValidateBase58Check(addr)
	}
	if strings.HasPrefix(addr, "ltc1") {
		return that.validateBech32(addr, "ltc")
	}
	return false
}

/////////////////////////////
// ETH (EIP-55)
/////////////////////////////

func (that CryptoAddressValidate) ValidateETH(addr string) bool {
	if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
		return false
	}

	raw := addr[2:]
	_, err := hex.DecodeString(raw)
	if err != nil {
		return false
	}

	if raw == strings.ToLower(raw) || raw == strings.ToUpper(raw) {
		return true
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(strings.ToLower(raw)))
	sum := hash.Sum(nil)

	for i := 0; i < len(raw); i++ {
		c := raw[i]
		if c >= '0' && c <= '9' {
			continue
		}
		v := sum[i/2]
		if i%2 == 0 {
			v >>= 4
		} else {
			v &= 0x0f
		}

		if (v > 7 && c != byte(strings.ToUpper(string(c))[0])) ||
			(v <= 7 && c != byte(strings.ToLower(string(c))[0])) {
			return false
		}
	}

	return true
}

/////////////////////////////
// SOL
/////////////////////////////

const b58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var b58Map = func() map[rune]int {
	m := make(map[rune]int)
	for i, c := range b58Alphabet {
		m[c] = i
	}
	return m
}()

func (that CryptoAddressValidate) base58DecodeStrict(input string) ([]byte, bool) {
	result := big.NewInt(0)
	base := big.NewInt(58)

	for _, c := range input {
		val, ok := b58Map[c]
		if !ok {
			return nil, false
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(val)))
	}

	decoded := result.Bytes()

	// 处理前导 '1' → 0x00
	leadingZeros := 0
	for _, c := range input {
		if c != '1' {
			break
		}
		leadingZeros++
	}

	out := make([]byte, leadingZeros+len(decoded))
	copy(out[leadingZeros:], decoded)

	return out, true
}

func (that CryptoAddressValidate) ValidateSOL(addr string) bool {
	// 合理长度范围（经验过滤）
	if len(addr) < 32 || len(addr) > 44 {
		return false
	}

	b, ok := that.base58DecodeStrict(addr)
	if !ok {
		return false
	}

	return len(b) == 32
}

// ///////////////////////////
// NANO
// ///////////////////////////
const nanoAlphabet = "13456789abcdefghijkmnopqrstuwxyz"

var nanoMap = func() map[rune]byte {
	m := make(map[rune]byte)
	for i, c := range nanoAlphabet {
		m[c] = byte(i)
	}
	return m
}()

func (that CryptoAddressValidate) decodeNanoBase32ToBits(input string) ([]byte, error) {
	// 每个字符 5 bit
	totalBits := len(input) * 5
	out := make([]byte, (totalBits+7)/8)

	bitPos := 0

	for _, c := range input {
		v, ok := nanoMap[c]
		if !ok {
			return nil, errors.New("invalid char")
		}

		for i := 4; i >= 0; i-- {
			bit := (v >> i) & 1

			byteIndex := bitPos / 8
			bitIndex := 7 - (bitPos % 8)

			if bit == 1 {
				out[byteIndex] |= 1 << bitIndex
			}

			bitPos++
		}
	}

	return out, nil
}

func (that CryptoAddressValidate) ValidateNANO(addr string) bool {
	if !strings.HasPrefix(addr, "nano_") && !strings.HasPrefix(addr, "xrb_") {
		return false
	}

	raw := addr[5:]
	if len(raw) != 60 {
		return false
	}

	keyPart := raw[:52]
	checkPart := raw[52:]

	// ✅ 1. 解码 bit 流（260 bit）
	keyBits, err := that.decodeNanoBase32ToBits(keyPart)
	if err != nil {
		return false
	}

	if len(keyBits) < 33 {
		return false
	}

	// ✅ 2. 丢掉前 4 bit
	pubKey := make([]byte, 32)
	for i := 0; i < 32; i++ {
		pubKey[i] = (keyBits[i] << 4) | (keyBits[i+1] >> 4)
	}

	// ✅ 3. checksum 解码（40 bit）
	checkBits, err := that.decodeNanoBase32ToBits(checkPart)
	if err != nil || len(checkBits) < 5 {
		return false
	}

	checksum := checkBits[:5]

	// ✅ 4. Blake2b-5
	h, _ := blake2b.New(5, nil)
	h.Write(pubKey)
	sum := h.Sum(nil)

	// ✅ 5. reverse
	for i := 0; i < len(sum)/2; i++ {
		sum[i], sum[len(sum)-1-i] = sum[len(sum)-1-i], sum[i]
	}

	return bytes.Equal(sum, checksum)
}

/////////////////////////////
// TRON
/////////////////////////////

func (that CryptoAddressValidate) ValidateTRON(addr string) bool {
	// 地址必须以 T 开头（主网）
	if !strings.HasPrefix(addr, "T") {
		return false
	}

	decoded, err := that.base58Decode(addr)
	if err != nil || len(decoded) != 25 {
		return false
	}

	payload := decoded[:21]
	checksum := decoded[21:]

	// Base58Check 校验
	if !bytes.Equal(that.doubleSHA256(payload)[:4], checksum) {
		return false
	}

	// TRON 地址必须是 0x41 开头
	if payload[0] != 0x41 {
		return false
	}

	return true
}
