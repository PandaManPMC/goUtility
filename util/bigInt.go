package util

import (
	"errors"
	"fmt"
	"math/big"
)

func NewBigInt(val string) (*big.Int, error) {
	v, isOk := big.NewInt(0).SetString(val, 10)
	if !isOk {
		return nil, errors.New(fmt.Sprintf("%s format int is error", val))
	}
	return v, nil
}

// NewBigIntByStringPositive val 必须是整数并且大于等于 0（不可为负数）
func NewBigIntByStringPositive(val string) (*big.Int, bool) {
	v, isOk := big.NewInt(0).SetString(val, 10)
	if !isOk {
		return nil, false
	}
	isLessZero := BigIntLessThanZero(v)
	if isLessZero {
		return nil, false
	}
	return v, true
}

// BigIntAdd 加法（不会影响原值），return a+b 不影响 a 和 b
func BigIntAdd(a, b *big.Int) big.Int {
	//c := a.Add(&a, &b) // 会影响原来 a 的值，弃用。
	c := big.NewInt(0)
	c.Add(c, a)
	c.Add(c, b)
	return *c
}

// BigIntSub 减法（不影响原值）,return a - b 不影响 a 和 b
func BigIntSub(a, b *big.Int) big.Int {
	c := big.NewInt(0)
	c.SetString(a.String(), 10)
	return *c.Sub(c, b)
}

// BigIntAddMul 乘法（不影响原值），return a * b 不影响 a 和 b
func BigIntAddMul(a, b *big.Int) big.Int {
	c := big.NewInt(0)
	c.SetString(a.String(), 10)
	c = c.Mul(c, b)
	return *c
}

// BigIntAddMul100 乘法（不影响原值），return a * 100
func BigIntAddMul100(a *big.Int) big.Int {
	c := big.NewInt(0)
	c.SetString(a.String(), 10)

	return *c.Mul(c, big.NewInt(100))
}

// BigIntQuo 除法（不影响原值），return a / b 不影响 a 和 b
func BigIntQuo(a, b *big.Int) big.Int {
	c := big.NewInt(0)
	c.SetString(a.String(), 10)

	c = c.Quo(c, b)
	return *c
}

// BigIntEquals 相等 【a == b】 true
func BigIntEquals(a, b *big.Int) bool {
	c := a.String()
	d := b.String()
	return c == d
}

// BigIntGreater 大于 【a > b】 true
func BigIntGreater(a, b *big.Int) bool {
	if 1 == a.Cmp(b) {
		return true
	}
	return false
}

// BigIntGreaterStr 大于 【a1 > b1】 true
func BigIntGreaterStr(a1, b1 string) bool {
	a, _ := big.NewInt(0).SetString(a1, 10)
	b, _ := big.NewInt(0).SetString(b1, 10)

	if 1 == a.Cmp(b) {
		return true
	}
	return false
}

// BigIntLess 小于 【a < b】 true
// 入参 big.Float 需要是同一个类型转换，如果是 "0.1" 与 0.1 的 big.Float 进行比较， 0.1 大于 "0.1"
func BigIntLess(a, b *big.Int) bool {
	if -1 == a.Cmp(b) {
		return true
	}
	return false
}

// BigIntLessByStr 小于 【a1 < b1】 true
func BigIntLessByStr(a1, b1 string) bool {
	a, _ := big.NewInt(0).SetString(a1, 10)
	b, _ := big.NewInt(0).SetString(b1, 10)
	if -1 == a.Cmp(b) {
		return true
	}
	return false
}

// BigIntGreaterOrEquals 大于等于 【a >= b】 true
func BigIntGreaterOrEquals(a, b big.Float) bool {
	if 0 <= a.Cmp(&b) {
		return true
	}
	return false
}

// BigIntGreaterOrEqualsByStr 大于等于 【a >= b】 true
func BigIntGreaterOrEqualsByStr(a1, b1 string) bool {
	a, _ := big.NewInt(0).SetString(a1, 10)
	b, _ := big.NewInt(0).SetString(b1, 10)
	if 0 <= a.Cmp(b) {
		return true
	}
	return false
}

// BigIntLessOrEquals 小于等于 【a <= b】 true
func BigIntLessOrEquals(a, b *big.Int) bool {
	if 0 >= a.Cmp(b) {
		return true
	}
	return false
}

// BigIntNotZero 非零 a != 0 则 true
func BigIntNotZero(a *big.Int) bool {
	if 0 != big.NewInt(0).Cmp(a) {
		return true
	}
	return false
}

// BigIntBiggerThanZero 大于0 a > 0 则 true
func BigIntBiggerThanZero(a *big.Int) bool {
	if 1 == a.Cmp(big.NewInt(0)) {
		return true
	}
	return false
}

// BigIntLessThanZero 小于0 a < 0 则 true
func BigIntLessThanZero(a *big.Int) bool {
	if -1 == a.Cmp(big.NewInt(0)) {
		return true
	}
	return false
}

// BigIntLessOrEqualsZero 若 a 小于等于 0 则返回 true
func BigIntLessOrEqualsZero(a *big.Int) bool {
	if 0 >= a.Cmp(big.NewInt(0)) {
		return true
	}
	return false
}

// BigIntBiggerOrEqualsZero 若 a 大于等于 0 则返回 true
func BigIntBiggerOrEqualsZero(a *big.Int) bool {
	if a.Cmp(big.NewInt(0)) >= 0 {
		return true
	}
	return false
}
