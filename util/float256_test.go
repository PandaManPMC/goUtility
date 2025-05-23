package util

import (
	"fmt"
	"math/big"
	"testing"
)

// TestFloat256 加减乘除极限测试
func TestFloat256(t *testing.T) {
	s := "1001.000000000000000001"
	f, _ := NewFloat256ByString(s)
	s2 := "205.000000000000005"
	f2, _ := NewFloat256ByString(s2)

	// 相加 （测试结果正常）
	f3 := f2.Add(f2, f)
	t.Log(f3) // 1206.000000000000005001

	// 相减 （测试结果正常）
	f4 := f3.Sub(f3, f)
	t.Log(f4) // 205.000000000000005

	// 乘 （正常）
	t.Log(f)
	t.Log(f2)
	f5 := f2.Mul(f2, f)
	t.Log(f5)
	// 205205.000000000005005205000000000000005
	// 205205.000000000005005205
	// 205205.000000000005005205
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f5)) // 保留有效小数位，18位后 205205.000000000005005205

	// 除 （正常）
	f7, _ := NewFloat256ByString(s2)
	f8, _ := NewFloat256ByString(s)
	t.Log(f8) // 1001.000000000000000001
	t.Log(f7) // 205.000000000000005
	f6 := f8.Quo(f8, f7)
	t.Log(f6)                                                    // 4.8829268292682925638359309934562789308309513791151480285133809971915114996736
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f6)) // 4.882926829268292564
}

func TestFloat256_2(t *testing.T) {
	// float256 陷阱
	s := "1001.000000000000000001"
	f, _ := NewFloat256ByString(s)
	s2 := "205.000000000000005"
	f2, _ := NewFloat256ByString(s2)

	t.Log(f) // 1001.000000000000000001

	// big.Float 固有的bug，调用 myTestFloat256Add 进行运算后，即使是取指针值拷贝依然会影响后续的使用。因此要再次使用 f和f2 就要使用 string 再次转成 float 使用
	// big.Float 值拷贝传递后 big.Float 真实存储数据的内存依然是指针，指向同一片内存导致值拷贝后函数修改依然会改变原值
	a := myTestFloat256Add(*f, *f2) // 计算后，f 会发生改变
	t.Log(&a)                       // 1206.000000000000005001

	t.Log(f)  // 603.0000000000000025005
	t.Log(f2) // 205.000000000000005
}

func myTestFloat256Add(a, b big.Float) big.Float {
	c := a.Add(&a, &b) // 依然会影响原来的值
	return *c
}

func TestFloat256Add(t *testing.T) {
	// 避开 float256 陷阱
	s := "1001.000000000000000001"
	f, _ := NewFloat256ByString(s)
	s2 := "205.000000000000005"
	f2, _ := NewFloat256ByString(s2)
	t.Log(f) // 1001.000000000000000001

	a := Float256Add(*f, *f2)
	t.Log(&a) // 1206.000000000000005001
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005
}

func TestFloat256Sub(t *testing.T) {
	s := "1001.000000000000000001"
	f, _ := NewFloat256ByString(s)
	s2 := "205.000000000000005"
	f2, _ := NewFloat256ByString(s2)
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005

	a := Float256Sub(*f, *f2)
	t.Log(&a) // 795.999999999999995001
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005
}

func TestFloat256Mul(t *testing.T) {
	s := "1001.000000000000000001"
	f, _ := NewFloat256ByString(s)
	s2 := "205.000000000000005"
	f2, _ := NewFloat256ByString(s2)
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005

	a := Float256Mul(*f, *f2)
	t.Log(&a) // 205205.000000000005005205000000000000005
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005
}

func TestFloat256Quo(t *testing.T) {
	s := "1001.000000000000000001"
	f, _ := NewFloat256ByString(s)
	s2 := "205.000000000000005"
	f2, _ := NewFloat256ByString(s2)
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005

	a := Float256Quo(*f, *f2)
	t.Log(&a) // 4.8829268292682925638359309934562789308309513791151480285133809971915114996736
	t.Log(f)  // 1001.000000000000000001
	t.Log(f2) // 205.000000000000005
}

func TestFloat256Max(t *testing.T) {
	var a uint64 = 0
	t.Log(a - 1)

	s := "18446744073709551615.18446744073709551615"
	f, b := NewFloat256ByString(s)
	t.Log(b)
	t.Log(f)
	t.Log(fmt.Sprintf("%.18f", f)) // 18446744073709551615.184467440737095516

	// 比 babyDoge 总量还大的数
	s2 := "4200000000000000000000000005.18446744073709551615"
	f2, b2 := NewFloat256ByString(s2)
	t.Log(b2)
	t.Log(f2)
	t.Log(fmt.Sprintf("%.18f", f2)) // 4200000000000000000000000005.184467440737095516

	// 极限数转换函数测试
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f2)) // 4200000000000000000000000005.184467440737095516

	// fmt.Sprintf 的四舍五入问题，第19位是5,第18位+1
	s3 := "4200000000000000000000000005.18446744073709551655"
	f3, _ := NewFloat256ByString(s3)
	t.Log(fmt.Sprintf("%.18f", f3)) // 4200000000000000000000000005.184467440737095517

	// 30个整数30个小数
	s4 := "420000000000000000000000000589.420000000000000002100000000589"
	f4, _ := NewFloat256ByString(s4)
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f4)) // 420000000000000000000000000589.420000000000000002

	s5 := "420000000000000000000000000589.420000000000000002100000000589"
	f5, _ := NewFloat256ByString(s5)
	t.Log(GetInstanceByNumberUtil().EfficientFloatToPrecisionString(*f5, 6))

	s6 := "420000000000000000000000000589.0"
	s7 := "420000000000000000000000000589"
	f6, _ := NewFloat256ByString(s6)
	f7, _ := NewFloat256ByString(s7)

	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f6))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f7))

	// 40 位整数 40 位小数
	s8 := "4200000000000000000000000005890000067850.4200000000002200000000000005890000067851"
	f8, _ := NewFloat256ByString(s8)
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f8)) // 4200000000000000000000000005890000067850.42000000000022

	// 极限 MAX 测试【整数 59 位，40位小数】
	s9 := "92000000000000000000000000058900000678502891230000123456789.4200000000002212310000000001890000067851"
	f9, _ := NewFloat256ByString(s9)
	t.Log(fmt.Sprintf("%.40f", f9))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f9)) // 92000000000000000000000000058900000678502891230000123456789.42000000000022

	t.Log(f9.MinPrec())

	// 负数
	s10 := "-92000000000000000000000000058900000678502891230000123456789.4200000000002212310000000001890000067851"
	f10, _ := NewFloat256ByString(s10)
	t.Log(fmt.Sprintf("%.40f", f10))
}

func TestCompare(t *testing.T) {
	f1, _ := NewFloat256ByString("100.88")
	f2, _ := NewFloat256ByString("3500.1212")

	t.Log(Float256Equals(*f1, *f2))          // 相等 false
	t.Log(Float256Greater(*f1, *f2))         // 大于 false
	t.Log(Float256Less(*f1, *f2))            // 小于 true
	t.Log(Float256GreaterOrEquals(*f1, *f2)) // 大于等于 false
	t.Log(Float256LessOrEquals(*f1, *f2))    // 小于等于 true

	f3, _ := NewFloat256ByString("100.88")
	t.Log(Float256Equals(*f1, *f3)) // 等于 true
}

func TestFloat256_(t *testing.T) {
	f1, b1 := NewFloat256ByString("-12.4343")
	f2, b2 := NewFloat256ByString("-0.051")
	t.Log(f1)
	t.Log(b1)
	t.Log(f2)
	t.Log(b2)

	f3, _ := NewFloat256ByString("12")
	f4 := Float256Add(*f3, *f1)
	t.Log(&f4)

	f5 := NewFloat256ByFloat64(19791.9609375)
	f6, _ := NewFloat256ByString("0.052")

	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f1))

	f7 := Float256Quo(*f6, *f5)
	t.Log(&f7)
	t.Log(fmt.Sprintf("%.18f", &f7))
	t.Log(fmt.Sprintf("%.19f", &f7))

}

func TestFloat256NotZero(t *testing.T) {

	t.Log(Float256NotZero(*big.NewFloat(0)))
	t.Log(Float256NotZero(*big.NewFloat(0.0)))
	t.Log(Float256NotZero(*big.NewFloat(0.1)))
	t.Log(Float256NotZero(*big.NewFloat(-10.2)))

}

func TestFloat256BiggerThanZero(t *testing.T) {

	t.Log(Float256BiggerThanZero(*big.NewFloat(0)))
	t.Log(Float256BiggerThanZero(*big.NewFloat(0.0)))
	t.Log(Float256BiggerThanZero(*big.NewFloat(0.1)))
	t.Log(Float256BiggerThanZero(*big.NewFloat(-10.2)))

}

func TestFloat256LessThanZero(t *testing.T) {

	t.Log(Float256LessThanZero(*big.NewFloat(0)))
	t.Log(Float256LessThanZero(*big.NewFloat(0.0)))
	t.Log(Float256LessThanZero(*big.NewFloat(0.1)))
	t.Log(Float256LessThanZero(*big.NewFloat(-10.2)))

}

func TestNewFloat256ByStringSafety(t *testing.T) {
	_, err := NewFloat256ByStringSafety("1.0")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("1")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("0")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("-10")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("-1.0")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("-1.2")
	t.Log(err)
	t.Log("---------------------------------------")

	_, err = NewFloat256ByStringSafety("1.a")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("1d.033")
	t.Log(err)
	_, err = NewFloat256ByStringSafety("1.0d02")
	t.Log(err)
}

func TestNewFloat256ByStringPositive(t *testing.T) {
	_, err := NewFloat256ByStringPositive("1.0")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("1")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("0")
	t.Log(err)

	t.Log("---------------------------------------")
	_, err = NewFloat256ByStringPositive("0.")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("-10")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("-1.0")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("-1.2")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("1.a")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("1d.033")
	t.Log(err)
	_, err = NewFloat256ByStringPositive("1.0d02")
	t.Log(err)
}

func TestFloat256ToStr(t *testing.T) {
	f := NewFloat256ByStringMust("1.2334")
	t.Log(fmt.Sprintf("%s", f))
}

func TestFloat256LessOrEqualsZero(t *testing.T) {
	f := NewFloat256ByStringMust("1")
	t.Log(Float256LessOrEqualsZero(*f))
	t.Log(Float256LessOrEqualsZero(*NewFloat256ByStringMust("0")))
	t.Log(Float256LessOrEqualsZero(*NewFloat256ByStringMust("-20")))

}

func TestFloat256MulAccumulative(t *testing.T) {
	o := Float256MulAccumulative(*NewFloat256ByInt64(4), *NewFloat256ByInt64(5), *NewFloat256ByInt64(2))
	t.Log(o.Float64())
	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256MulAccumulative(*NewFloat256ByStringMust("0.3"), *NewFloat256ByInt64(-5), *NewFloat256ByInt64(2))))

	a := 1.001
	b := "1"

	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Mul(*NewFloat256ByFloat64(a), *NewFloat256ByStringMust(b))))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringBy8(Float256Quo(*NewFloat256ByStringMust(b), *NewFloat256ByFloat64(a))))

}

func TestLess(t *testing.T) {
	a := "0.1"
	b := 0.1
	af := *NewFloat256ByStringMust(a)
	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringBy8(af))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringBy8(*NewFloat256ByFloat64(b)))

	t.Log(Float256Less(*NewFloat256ByStringMust(a), *NewFloat256ByFloat64(b)))

	c := "0.1"
	t.Log(Float256Less(*NewFloat256ByStringMust(a), *NewFloat256ByStringMust(c)))

	// true 严重 BUG 坑！
	t.Log(Float256Greater(*NewFloat256ByFloat64(b), *NewFloat256ByStringMust(a)))

	// false 请使用这个
	t.Log(Float256LessByStr(a, fmt.Sprintf("%f", b)))
}

func TestFloat256Less1E9(t *testing.T) {
	a, isOk := big.NewInt(0).SetString("1999", 10)
	t.Log(isOk)
	t.Log(a)

	t.Log(a.Int64())
	t.Log(a.String())

	b, isOk := big.NewInt(0).SetString("1999", 10)
	t.Log(isOk)

	a.Sub(b, a)
	t.Log(a.String())

}
