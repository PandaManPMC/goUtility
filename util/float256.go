package util

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strings"
)

//  author: pmc
//  since: 2022/9/1
//  desc: util Float256 相关运算、转换和比较等函数，该工具的目标是能够完成区块链数字货币的极限运算。
//  float256Prec 为 256 时最大表示正负 59位整数，小数精确保存18位，不能完整表示区块链 uint256。
//  调整 float256Prec 可以增大精度，为了防止运算溢出默认设置为 512，可以完整表示区块链 uint256。
//	mysql 数据库的 decimal(65,18) 类型，最长保存 65 个数字，如果保存18个小数则整数只能保存 47 个，
//	最大是 【正负 99999999999999999999999999999999999999999999999.999999999999999999】，离 uint256 有比较大差距。
//	decimal(65,18) 用 decimal.Decimal 表示而不能用 float256，由于 decimal 用科学计数法表示所以要在传输中转为 string（包括 protobuf），避免精度丢失导致前台展示出错误或可读性差的科学计数法字符。

const (
	float256Prec = 256
)

func NewFloat256() *big.Float {
	return new(big.Float).SetPrec(float256Prec)
}

var zeroFloat256 *big.Float

func init() {
	zeroFloat256 = NewFloat256ByStringMust("0")
}

// NewFloat256ByStringSafety 当转换失败时返回 nil 和 异常
// 支持负数
func NewFloat256ByStringSafety(val string) (*big.Float, error) {
	req1 := regexp.MustCompile(`^-?\d+$`)
	req2 := regexp.MustCompile(`^-?\d+\.\d+$`)
	if !req1.MatchString(val) && !req2.MatchString(val) {
		return nil, errors.New("val is not number")
	}
	fl, isOk := new(big.Float).SetPrec(float256Prec).SetString(val)
	if !isOk {
		return nil, errors.New("parsing failed")
	}
	return fl, nil
}

// NewFloat256ByStringPositive 当转换失败时返回 nil 和 异常
// 当字符串为负数时会抛出异常，字符串不是正常数值也会转换失败抛出异常。
func NewFloat256ByStringPositive(val string) (*big.Float, error) {
	req1 := regexp.MustCompile(`^\d+$`)
	req2 := regexp.MustCompile(`^\d+\.\d+$`)
	if !req1.MatchString(val) && !req2.MatchString(val) {
		return nil, errors.New("val is not positive number")
	}
	fl, isOk := new(big.Float).SetPrec(float256Prec).SetString(val)
	if !isOk {
		return nil, errors.New("parsing failed")
	}
	return fl, nil
}

func NewFloat256ByString(fl string) (*big.Float, bool) {
	if "" == fl {
		fl = "0"
	}
	return new(big.Float).SetPrec(float256Prec).SetString(fl)
}

func NewFloat256ByStringMust(fl string) *big.Float {
	if "" == fl {
		fl = "0"
	}
	r, isOk := new(big.Float).SetPrec(float256Prec).SetString(fl)
	if isOk {
		return r
	}

	r, _ = new(big.Float).SetPrec(float256Prec).SetString("0")
	return r
}

func NewFloat256ByStringPanic(fl string) *big.Float {
	if "" == fl {
		panic("NewFloat256ByStringPanic val incorrect")
	}
	r, isOk := new(big.Float).SetPrec(float256Prec).SetString(fl)
	if !isOk {
		panic("NewFloat256ByStringPanic big.Float incorrect")
	}
	return r
}

func NewFloat256ByFloat64(f64 float64) *big.Float {
	sf := fmt.Sprintf("%.14f", f64)
	return NewFloat256ByStringMust(sf)
}

func NewFloat256ByFloat32(f32 float32) *big.Float {
	return new(big.Float).SetPrec(float256Prec).SetFloat64(float64(f32))
}

func NewFloat256ByInt64(num int64) *big.Float {
	return new(big.Float).SetPrec(float256Prec).SetInt64(num)
}

func NewFloat256ByUint64(num uint64) *big.Float {
	return new(big.Float).SetPrec(float256Prec).SetUint64(num)
}

// Float256Add 加法（不会影响原值，避开 big.Float 陷阱），return a+b 不影响 a 和 b
func Float256Add(a, b big.Float) big.Float {
	//c := a.Add(&a, &b) // 会影响原来 a 的值，弃用。
	c := NewFloat256()
	c.Add(c, &a)
	c.Add(c, &b)
	return *c
}

// Float256Sub 减法（不影响原值）,return a - b 不影响 a 和 b
func Float256Sub(a, b big.Float) big.Float {
	a2 := NewFloat256().Copy(&a)
	b2 := NewFloat256().Copy(&b)
	return *(a2.Sub(a2, b2))
}

// Float256Mul 乘法（不影响原值），return a * b 不影响 a 和 b
func Float256Mul(a, b big.Float) big.Float {
	a2 := NewFloat256().Copy(&a)
	b2 := NewFloat256().Copy(&b)
	return *(a2.Mul(a2, b2))
}

// Float256Mul100 乘法（不影响原值），return a * 100
func Float256Mul100(a big.Float) big.Float {
	a2 := NewFloat256().Copy(&a)
	return *(a2.Mul(a2, NewFloat256ByInt64(100)))
}

// Float256MulAccumulative 累乘
func Float256MulAccumulative(args ...big.Float) big.Float {
	o := NewFloat256ByInt64(1)
	for _, v := range args {
		a := NewFloat256().Copy(&v)
		o = o.Mul(a, o)
	}
	return *o
}

// Float256Quo 除法（不影响原值），return a / b 不影响 a 和 b
func Float256Quo(a, b big.Float) big.Float {
	a2 := NewFloat256().Copy(&a)
	b2 := NewFloat256().Copy(&b)
	return *(a2.Quo(a2, b2))
}

// Float256Quo100 除法（不影响原值），return a / 100
func Float256Quo100(a big.Float) big.Float {
	a2 := NewFloat256().Copy(&a)
	return *(a2.Quo(a2, NewFloat256ByInt64(100)))
}

// Float256Equals 相等 【a == b】 true
func Float256Equals(a, b big.Float) bool {
	if 0 == a.Cmp(&b) {
		return true
	}
	return false
}

// Float256Greater 大于 【a > b】 true
func Float256Greater(a, b big.Float) bool {
	if 1 == a.Cmp(&b) {
		return true
	}
	return false
}

// Float256GreaterStr 大于 【a1 > b1】 true
func Float256GreaterStr(a1, b1 string) bool {
	a := NewFloat256ByStringMust(a1)
	b := NewFloat256ByStringMust(b1)
	if 1 == a.Cmp(b) {
		return true
	}
	return false
}

// Float256Less 小于 【a < b】 true
// 入参 big.Float 需要是同一个类型转换，如果是 "0.1" 与 0.1 的 big.Float 进行比较， 0.1 大于 "0.1"
func Float256Less(a, b big.Float) bool {
	if -1 == a.Cmp(&b) {
		return true
	}
	return false
}

// Float256LessByStr 小于 【a1 < b1】 true
func Float256LessByStr(a1, b1 string) bool {
	a := NewFloat256ByStringMust(a1)
	b := NewFloat256ByStringMust(b1)
	if -1 == a.Cmp(b) {
		return true
	}
	return false
}

// Float256GreaterOrEquals 大于等于 【a >= b】 true
func Float256GreaterOrEquals(a, b big.Float) bool {
	if 0 <= a.Cmp(&b) {
		return true
	}
	return false
}

// Float256GreaterOrEqualsByStr 大于等于 【a >= b】 true
func Float256GreaterOrEqualsByStr(a1, b1 string) bool {
	a := NewFloat256ByStringMust(a1)
	b := NewFloat256ByStringMust(b1)
	if 0 <= a.Cmp(b) {
		return true
	}
	return false
}

// Float256LessOrEquals 小于等于 【a <= b】 true
func Float256LessOrEquals(a, b big.Float) bool {
	if 0 >= a.Cmp(&b) {
		return true
	}
	return false
}

// Float256NotZero 非零 a != 0 则 true
func Float256NotZero(a big.Float) bool {
	if 0 != zeroFloat256.Cmp(&a) {
		return true
	}
	return false
}

// Float256BiggerThanZero 大于0 a > 0 则 true
func Float256BiggerThanZero(a big.Float) bool {
	if 0 > zeroFloat256.Cmp(&a) {
		return true
	}
	return false
}

// Float256LessThanZero 小于0 a < 0 则 true
func Float256LessThanZero(a big.Float) bool {
	if 0 < zeroFloat256.Cmp(&a) {
		return true
	}
	return false
}

// Float256LessOrEqualsZero 若 a 小于等于 0 则返回 true
func Float256LessOrEqualsZero(a big.Float) bool {
	if 0 >= a.Cmp(zeroFloat256) {
		return true
	}
	return false
}

// Float256BiggerOrEqualsZero 若 a 大于等于 0 则返回 true
func Float256BiggerOrEqualsZero(a big.Float) bool {
	if a.Cmp(zeroFloat256) >= 0 {
		return true
	}
	return false
}

func Float256ToFloat32(a big.Float) float32 {
	f32, _ := a.Float32()
	return f32
}

func Float256ToFloat64(a big.Float) float64 {
	f32, _ := a.Float64()
	return f32
}

// ToDecimalsVal 精度转换(小数 -> 整数)
func ToDecimalsVal(decimals int, val big.Float) big.Float {
	if 0 == decimals {
		return val
	}
	dec := *NewFloat256ByInt64(int64(math.Pow10(decimals)))
	to := Float256Mul(val, dec)
	return to
}

// FormatDecimalsVal 精度格式化（整数 -> 小数）
func FormatDecimalsVal(decimals, val int) big.Float {
	if 0 == decimals {
		return *NewFloat256ByInt64(int64(val))
	}
	dec := *NewFloat256ByInt64(int64(math.Pow10(decimals)))
	to := Float256Quo(*NewFloat256ByInt64(int64(val)), dec)
	return to
}

// Float256ToString 格式化 big.Float 保留 2 位小数
func Float256ToString(fl big.Float) string {
	s := GetInstanceByNumberUtil().EfficientFloatToPrecisionString(fl, 3)
	if "-0" == s {
		s = "0"
	}
	p := strings.Split(s, ".")
	if 2 == len(p) {
		if 3 == len(p[1]) {
			s = p[1][:2]
			s = fmt.Sprintf("%s.%s", p[0], s)
		}
	}

	if "-0" == s || "-0.0" == s || "-0.00" == s || "0.0" == s || "0.00" == s {
		s = "0"
	}

	return s
}
