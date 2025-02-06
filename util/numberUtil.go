package util

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

type numberUtil struct {
}

var numberUtilInstance numberUtil

func GetInstanceByNumberUtil() *numberUtil {
	return &numberUtilInstance
}

// EfficientFloatLength 浮点数有效小数位长度，如 3.14000 > 小数有效长度 2
func (instance *numberUtil) EfficientFloatLength(fl big.Float) uint8 {
	return instance.EfficientFloatLengthByPrecision(fl, 18)
}

func (*numberUtil) EfficientFloatLengthByPrecision(fl big.Float, precisionLength uint8) uint8 {
	f := NewFloat256().Copy(&fl)
	sFloat := fmt.Sprintf("%."+fmt.Sprintf("%d", precisionLength)+"f", f)
	spl := strings.Split(sFloat, ".")
	flo := spl[1]
	var end uint8 = 0
	for i := len(flo) - 1; i >= 0; i-- {
		if "0" != flo[i:i+1] {
			end = uint8(i + 1)
			break
		}
	}
	if 0 == end {
		return 0
	} else {
		return end
	}
}

// EfficientFloatToString 保留浮点数最大18位有效精度的字符串（原生 float64 有效浮点位为15，注意精度），如果浮点数是 1.0 则返回 1，浮点数 1.230 则返回 1.23。
// 第 18 位会因为第 19 为被四舍五入，如果第 19 位小数大于 4 则第 18 位 +1。
func (instance *numberUtil) EfficientFloatToString(fl big.Float) string {
	return instance.EfficientFloatToPrecisionString(fl, 18)
}

// EfficientFloatToStringBy8 最多保留 8 位有效小数点位
func (instance *numberUtil) EfficientFloatToStringBy8(fl big.Float) string {
	return instance.EfficientFloatToPrecisionString(fl, 8)
}

// EfficientFloatToStringBy10 最多保留 10 位有效小数点位
func (instance *numberUtil) EfficientFloatToStringBy10(fl big.Float) string {
	return instance.EfficientFloatToPrecisionString(fl, 10)
}

// EfficientFloatToStringBy2 最多保留 2 位有效小数点位
func (instance *numberUtil) EfficientFloatToStringBy2(fl big.Float) string {
	return instance.EfficientFloatToPrecisionString(fl, 2)
}

// EfficientFloatToStringBy2FormatZero 最多保留 2 位有效小数点位，但是会进行格式化，如果是 0 则格式化成 0.00，是 1 则是 1.00
func (instance *numberUtil) EfficientFloatToStringBy2FormatZero(fl big.Float) string {
	s := instance.EfficientFloatToPrecisionString(fl, 2)
	if "0" == s {
		return "0.00"
	}
	inx := strings.Index(s, ".")
	if -1 == inx {
		return fmt.Sprintf("%s.00", s)
	}
	if inx == len(s)-2 {
		return fmt.Sprintf("%s0", s)
	}
	return s
}

//	EfficientFloatToStringBy4 最多保留 4 位有效小数点位
//func (instance *numberUtil) EfficientFloatToStringBy4(fl big.Float) string {
//	return instance.EfficientFloatToPrecisionString(fl, 4)
//}

// EfficientFloatToFloatBy4 最多保留 4 位有效小数 float64
func (instance *numberUtil) EfficientFloatToFloatBy4Must(fl big.Float) float64 {
	f, _ := strconv.ParseFloat(instance.EfficientFloatToPrecisionString(fl, 4), 10)
	return f
}

// EfficientFloatToFloatBy8Must 最多保留 8 位有效小数 float64
func (instance *numberUtil) EfficientFloatToFloatBy8Must(fl big.Float) float64 {
	f, _ := strconv.ParseFloat(instance.EfficientFloatToPrecisionString(fl, 8), 10)
	return f
}

// EfficientFloatToStringBy8 最多保留 8 位有效小数点位，并增加负数符号
func (instance *numberUtil) EfficientFloatToStringByNegative8(fl big.Float) string {
	fs := instance.EfficientFloatToPrecisionString(fl, 8)
	if !strings.HasPrefix(fs, "-") {
		fs = fmt.Sprintf("-%s", fs)
	}
	return fs
}

// EfficientFloatToStringByPositiveNumber2 最多保留 2 位有效小数点位，并去除负符号
func (instance *numberUtil) EfficientFloatToStringByPositiveNumber2(fl big.Float) string {
	fs := instance.EfficientFloatToPrecisionString(fl, 2)
	if strings.HasPrefix(fs, "-") {
		fs = fmt.Sprintf("%s", fs[1:])
	}
	return fs
}

// EfficientFloatToPrecisionString 保留指定 precisionLength 长度的浮点数有效精度的字符串
func (instance *numberUtil) EfficientFloatToPrecisionString(fl big.Float, precisionLength uint8) string {
	f := NewFloat256().Copy(&fl)
	sFloat := fmt.Sprintf("%."+fmt.Sprintf("%d", precisionLength)+"f", f)
	end := instance.EfficientFloatLengthByPrecision(*f, precisionLength)
	if 0 == end {
		return strings.Split(sFloat, ".")[0]
	}
	arr := strings.Split(sFloat, ".")
	return sFloat[:len(arr[0])+int(end)+1]
}

// EfficientFloatByString 保留字符串小数点后的有效位，比如 1.20000 保留位 1.2
func (instance *numberUtil) EfficientFloatByString(num string) string {
	if 0 == len(num) || !strings.Contains(num, ".") {
		return num
	}
	spl := strings.Split(num, ".")
	flo := spl[1]
	var end uint8 = 0
	for i := len(flo) - 1; i >= 0; i-- {
		if "0" != flo[i:i+1] {
			end = uint8(i + 1)
			break
		}
	}
	if 0 == end {
		return spl[0]
	} else {
		return num[:len(spl[0])+int(end)+1]
	}
}

// FloatToInteger 浮点数舍弃浮点转整数
func (instance *numberUtil) FloatToInteger(fl big.Float) int {
	f := NewFloat256().Copy(&fl)
	sFloat := fmt.Sprintf("%.18f", f)
	spl := strings.Split(sFloat, ".")
	if 0 < len(spl) && 0 != len(spl[0]) {
		n, err := strconv.Atoi(spl[0])
		if nil != err {
			return 0
		}
		return n
	}
	return 0
}

// StringFloatToPrecision 保留字符串小数指定位数【字符串切割没有四舍五入】
// f 在 1.222 时 precisionLength 为 1 则输出 1.2，1.0 则输出 1
func (instance *numberUtil) StringFloatToPrecision(f string, precisionLength uint) string {
	if 0 == len(f) || !strings.Contains(f, ".") {
		return f
	}
	spl := strings.Split(f, ".")
	if 2 != len(spl) || int(precisionLength) >= len(spl[1]) {
		return instance.EfficientFloatByString(f)
	}

	if 0 == precisionLength {
		return spl[0]
	}
	f2 := spl[1][:precisionLength]
	return instance.EfficientFloatByString(fmt.Sprintf("%s.%s", spl[0], f2))
}

// IsNumber 数字类型 true，否则 false，正负都可包括浮点数。
func (instance *numberUtil) IsNumber(val string) bool {
	if instance.IsInteger(val) || instance.IsFloat(val) {
		return true
	}
	return false
}

// IsPositiveNumber 正数值 true，否则 false
func (instance *numberUtil) IsPositiveNumber(val string) bool {
	if instance.IsPositiveInteger(val) || instance.IsPositiveFloat(val) {
		return true
	}
	return false
}

// IsInteger 整数
func (*numberUtil) IsInteger(val string) bool {
	req := regexp.MustCompile(`^-?\d+$`)
	return req.MatchString(val)
}

// IsPositiveInteger 正整数
func (*numberUtil) IsPositiveInteger(val string) bool {
	req := regexp.MustCompile(`^\d+$`)
	return req.MatchString(val)
}

// IsNegativeInteger 负整数
func (*numberUtil) IsNegativeInteger(val string) bool {
	req := regexp.MustCompile(`^-\d+$`)
	return req.MatchString(val)
}

// IsFloat 浮点数
func (*numberUtil) IsFloat(val string) bool {
	req := regexp.MustCompile(`^-?\d+\.\d+$`)
	return req.MatchString(val)
}

// IsPositiveFloat 正浮点数
func (*numberUtil) IsPositiveFloat(val string) bool {
	req := regexp.MustCompile(`^\d+\.\d+$`)
	return req.MatchString(val)
}

// IsNegativeFloat 负浮点数
func (*numberUtil) IsNegativeFloat(val string) bool {
	req := regexp.MustCompile(`^-\d+\.\d+$`)
	return req.MatchString(val)
}

// HexToInt 十六进制转十进制
func (*numberUtil) HexToInt(hex string) string {
	// 去掉前缀 "0x"
	if len(hex) > 2 && hex[:2] == "0x" {
		hex = hex[2:]
	}

	// 转换为十进制
	num := new(big.Int)
	num.SetString(hex, 16)
	return num.String()
}

// IntToHex0x 十进制转十六进制字符串0x
func (*numberUtil) IntToHex0x(val int64) string {
	hexStr := strconv.FormatInt(val, 16)
	return "0x" + hexStr
}

// ParseIntMust 屏蔽错误
func (*numberUtil) ParseIntMust(val string) int {
	v, _ := strconv.Atoi(val)
	return v
}
