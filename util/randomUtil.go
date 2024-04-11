package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type randomUtil struct {
	chaArr []string
}

var randomUtilInstance randomUtil

func GetInstanceByRandomUtil() *randomUtil {
	return &randomUtilInstance
}

func init() {
	randomUtilInstance.chaArr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
}

// RandCrypto 生成随机数 0-9 crypto/rand
// num int 随机数长度
func (that *randomUtil) RandCrypto(num int) string {
	str := strings.Builder{}
	for i := 0; i < num; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(10))
		str.WriteString(fmt.Sprintf("%d", result))
	}
	return str.String()
}

// RandCharacterString 生成 0-9 a-z A-Z 随机字符
// num int 指定生成字符数量
func (that *randomUtil) RandCharacterString(num int) string {
	str := strings.Builder{}
	max := len(that.chaArr)
	for i := 0; i < num; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		str.WriteString(that.chaArr[result.Int64()])
	}
	return str.String()
}

// RandNumber 生成随机数数字 max 以内
// max int 随机数最大
func (that *randomUtil) RandNumber(max int) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return result.Int64()
}

// RandNumberNotZero 生成随机数数字最大 max， 必定大于0的
// max int 随机数最大( result == max || result > 0 )
func (that *randomUtil) RandNumberNotZero(max int) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	r := result.Int64() + 1
	return r
}
