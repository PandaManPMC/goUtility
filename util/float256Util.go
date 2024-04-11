package util

import "math/big"

type float256 struct {
}

var float256Instance float256

func GetInstanceByFloat256() *float256 {
	return &float256Instance
}

// Float256AddToStr8 a+b 返回保留 8 位小数的字符串
func (*float256) Float256AddToStr8(a, b big.Float) string {
	return GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Add(a, b))
}

// Float256SubToStr8 a-b 返回保留 8 位小数的字符串
func (*float256) Float256SubToStr8(a, b big.Float) string {
	return GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Sub(a, b))
}

// Float256AddByStr8 加法 a+b ，字符串
// 成功时 error 为 nil，保留8位有效位的小数
// 当失败时会返回 "",error
func (*float256) Float256AddByStr8(a, b string) (string, error) {
	af, err1 := NewFloat256ByStringSafety(a)
	if nil != err1 {
		return "", err1
	}
	bf, err2 := NewFloat256ByStringSafety(b)
	if nil != err2 {
		return "", err2
	}
	return GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Add(*af, *bf)), nil
}

// Float256SubByStr8 减法 a - b，字符串
// 成功时 error 为 nil，保留8位有效位的小数
// 当失败时会返回 "",error
func (*float256) Float256SubByStr8(a, b string) (string, error) {
	af, err1 := NewFloat256ByStringSafety(a)
	if nil != err1 {
		return "", err1
	}
	bf, err2 := NewFloat256ByStringSafety(b)
	if nil != err2 {
		return "", err2
	}
	return GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Sub(*af, *bf)), nil
}

// Float256AddByStr8Positive 加法 a+b ，字符串，加数必须是正数
// 成功时 error 为 nil，保留8位有效位的小数
// 当失败时会返回 "",error
func (*float256) Float256AddByStr8Positive(a, b string) (string, error) {
	af, err1 := NewFloat256ByStringPositive(a)
	if nil != err1 {
		return "", err1
	}
	bf, err2 := NewFloat256ByStringPositive(b)
	if nil != err2 {
		return "", err2
	}
	return GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Add(*af, *bf)), nil
}

// Float256SubByStr8Positive 减法 a - b，字符串，减数与被减数都必须是正数
// 成功时 error 为 nil，保留8位有效位的小数
// 当失败时会返回 "",error
func (*float256) Float256SubByStr8Positive(a, b string) (string, error) {
	af, err1 := NewFloat256ByStringPositive(a)
	if nil != err1 {
		return "", err1
	}
	bf, err2 := NewFloat256ByStringPositive(b)
	if nil != err2 {
		return "", err2
	}
	return GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Sub(*af, *bf)), nil
}
