package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

type randomUtil struct {
	chars            string
	lowercaseLetters string
	uppercaseLetters string
	numbers          string
}

var randomUtilInstance randomUtil

func GetInstanceByRandomUtil() *randomUtil {
	return &randomUtilInstance
}

func init() {
	randomUtilInstance.chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomUtilInstance.lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	randomUtilInstance.uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomUtilInstance.numbers = "0123456789"
}

func (that *randomUtil) randString(chars string, num int) (string, error) {
	if num <= 0 {
		return "", nil
	}

	var str strings.Builder
	str.Grow(num)

	max_ := big.NewInt(int64(len(chars)))

	for i := 0; i < num; i++ {
		result, err := rand.Int(rand.Reader, max_)
		if err != nil {
			return "", err
		}

		str.WriteByte(chars[result.Int64()])
	}

	return str.String(), nil
}

// RandCrypto 生成随机数 0-9 crypto/rand
// num int 随机数长度
func (that *randomUtil) RandCrypto(num int) (string, error) {
	return that.randString(that.numbers, num)
}

// RandCharacterString 生成 0-9 a-z A-Z 随机字符
// num int 指定生成字符数量
func (that *randomUtil) RandCharacterString(num int) (string, error) {
	return that.randString(that.chars, num)
}

func (that *randomUtil) RandLowercaseLetters(num int) (string, error) {
	return that.randString(that.lowercaseLetters, num)
}

func (that *randomUtil) RandUppercaseLetters(num int) (string, error) {
	return that.randString(that.uppercaseLetters, num)
}

// RandNumber 生成 max 范围内的随机整数
func (that *randomUtil) RandNumber(max int) (int64, error) {
	result, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if nil != err {
		return 0, err
	}
	return result.Int64(), nil
}

// RandNumberNotZero 生成 1 ~ max 的随机整数
func (that *randomUtil) RandNumberNotZero(max int) (int64, error) {
	result, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if nil != err {
		return 0, err
	}
	r := result.Int64() + 1
	return r, nil
}
