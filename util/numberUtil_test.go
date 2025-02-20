package util

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestEfficientFloatToString(t *testing.T) {
	s1 := "0.000000000000000001"
	s2 := "2.0"
	s3 := "3.0000067890000"
	s4 := "2"
	s5 := "2.50"
	f1, _ := new(big.Float).SetString(s1)
	f2, _ := new(big.Float).SetString(s2)
	f3, _ := new(big.Float).SetString(s3)
	f4, _ := new(big.Float).SetString(s4)
	f5, _ := new(big.Float).SetString(s5)

	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f1))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f2))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f3))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f4))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f5))

	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f1))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f2))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f3))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f4))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f5))
}

func TestEfficientFloatByString(t *testing.T) {
	s1 := "0.000000000000000001"
	s2 := "2.0"
	s3 := "3.0000067890000"
	s4 := "2"
	s5 := "2.50"

	t.Log(GetInstanceByNumberUtil().EfficientFloatByString(s1))
	t.Log(GetInstanceByNumberUtil().EfficientFloatByString(s2))
	t.Log(GetInstanceByNumberUtil().EfficientFloatByString(s3))
	t.Log(GetInstanceByNumberUtil().EfficientFloatByString(s4))
	t.Log(GetInstanceByNumberUtil().EfficientFloatByString(s5))

	s6 := ""
	f6, b := new(big.Float).SetString(s6)
	t.Log(f6)
	t.Log(b) // 失败
}

func TestFString(t *testing.T) {
	s := "150868.557"
	bal, rea21 := new(big.Float).SetPrec(256).SetString(s)
	t.Log(rea21)
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*bal))

	s2 := "150685.557000"
	f2, _ := NewFloat256().SetString(s2)
	t.Log(GetInstanceByNumberUtil().EfficientFloatToString(*f2))
}

func TestFloatToInteger(t *testing.T) {
	f1, _ := NewFloat256().SetString("0.1000")
	f2, _ := NewFloat256().SetString("5550.1000")
	f3, _ := NewFloat256().SetString("0")
	f4, _ := NewFloat256ByString("")
	f5, _ := NewFloat256().SetString("1.1000")

	t.Log(f4)
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f1))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f2))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f3))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f4))
	t.Log(GetInstanceByNumberUtil().FloatToInteger(*f5))

}

func TestEfficientFloatToPrecisionString(t *testing.T) {
	decimals, _ := strconv.ParseInt("6", 10, 64)
	amount, _ := NewFloat256ByString("1.123456789")
	ret := GetInstanceByNumberUtil().EfficientFloatToPrecisionString(*amount, uint8(decimals))
	r, _ := NewFloat256ByString(ret)
	t.Log(ret)
	t.Log(r)

	t.Log(fmt.Sprintf("%.7g", amount))
}

func TestFmtRounding(t *testing.T) {
	// fmt 四舍五入问题

	f1, _ := NewFloat256ByString("1.2352")
	f2, _ := NewFloat256ByString("1.2328")
	f3, _ := NewFloat256ByString("1.2349")

	t.Log(fmt.Sprintf("%.2f", f1))
	t.Log(fmt.Sprintf("%.2f", f2))
	t.Log(fmt.Sprintf("%.2f", f3))

}

func TestStringFloatToPrecision(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().StringFloatToPrecision("1.222121", 3))
	t.Log(GetInstanceByNumberUtil().StringFloatToPrecision("1.2", 3))
	t.Log(GetInstanceByNumberUtil().StringFloatToPrecision("1", 3))
	t.Log(GetInstanceByNumberUtil().StringFloatToPrecision("1.230", 3))    // 1.23
	t.Log(GetInstanceByNumberUtil().StringFloatToPrecision("1.220121", 3)) // 1.22
	t.Log(GetInstanceByNumberUtil().StringFloatToPrecision("1.20", 3))     // 1.2

}

func TestIsInteger(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("0"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("2"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12012359"))
	t.Log(GetInstanceByNumberUtil().IsInteger("-12"))
	t.Log("----------------------")
	t.Log(GetInstanceByNumberUtil().IsInteger("12.2"))
	t.Log(GetInstanceByNumberUtil().IsInteger("12.d"))
	t.Log(GetInstanceByNumberUtil().IsInteger("12.d"))
}

func TestIsPositiveInteger(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("0"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("2"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12012359"))
	t.Log("----------------------")
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("120.12359"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger(""))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("-12"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12.2"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12.d"))
	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12.d"))
}

func TestIsNegativeInteger(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("-12"))

	t.Log("----------------------")
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("0"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("2"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("12"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("12012359"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("120.12359"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger(""))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("-12.2"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("12.2"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("12.d"))
	t.Log(GetInstanceByNumberUtil().IsNegativeInteger("12.d"))
}

func TestIsFloat(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsFloat("120.12359"))
	t.Log(GetInstanceByNumberUtil().IsFloat("-12.2"))
	t.Log(GetInstanceByNumberUtil().IsFloat("-12.0"))
	t.Log("----------------------")
	t.Log(GetInstanceByNumberUtil().IsFloat("-12"))
	t.Log(GetInstanceByNumberUtil().IsFloat("0"))
	t.Log(GetInstanceByNumberUtil().IsFloat("2"))
	t.Log(GetInstanceByNumberUtil().IsFloat("12"))
	t.Log(GetInstanceByNumberUtil().IsFloat("12012359"))
	t.Log(GetInstanceByNumberUtil().IsFloat(""))
	t.Log(GetInstanceByNumberUtil().IsFloat("12.d"))
	t.Log(GetInstanceByNumberUtil().IsFloat("12.d"))
}

func TestIsPositiveFloat(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("120.0"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("120.12359"))
	t.Log("----------------------")
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("120."))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("-12.2"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("-12.0"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("-12"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("0"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("2"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("12"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("12012359"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat(""))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("12.d"))
	t.Log(GetInstanceByNumberUtil().IsPositiveFloat("12.d"))
}

func TestIsNegativeFloat(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("-12.2"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("-12.0"))
	t.Log("----------------------")
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("120."))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("120.0"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("120.12359"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("-12"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("0"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("2"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("12"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("12012359"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat(""))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("12.d"))
	t.Log(GetInstanceByNumberUtil().IsNegativeFloat("12.d"))
}

func TestIsNumber(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsNumber("1"))
	t.Log(GetInstanceByNumberUtil().IsNumber("12"))
	t.Log(GetInstanceByNumberUtil().IsNumber("-1"))
	t.Log(GetInstanceByNumberUtil().IsNumber("0"))
	t.Log(GetInstanceByNumberUtil().IsNumber("-12"))
	t.Log(GetInstanceByNumberUtil().IsNumber("1.0"))
	t.Log(GetInstanceByNumberUtil().IsNumber("12.3"))
	t.Log(GetInstanceByNumberUtil().IsNumber("-1.0"))
	t.Log(GetInstanceByNumberUtil().IsNumber("-12.3"))

	t.Log("------------------------------")

	t.Log(GetInstanceByNumberUtil().IsNumber("12.DSA"))
	t.Log(GetInstanceByNumberUtil().IsNumber("12.d"))

	t.Log("------------------------------")

	t.Log(GetInstanceByNumberUtil().IsPositiveInteger("12.d"))

	t.Log(GetInstanceByNumberUtil().IsFloat("12.d"))
	t.Log(GetInstanceByNumberUtil().IsNumber("as.ds"))
	t.Log(GetInstanceByNumberUtil().IsNumber("as"))
	t.Log(GetInstanceByNumberUtil().IsNumber("a"))

}

func TestIsPositiveNumber(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("1"))
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("1.23"))
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("1.0"))

	t.Log("-------------------------")
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber(""))
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("1."))
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("-1"))
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("-1.23"))
	t.Log(GetInstanceByNumberUtil().IsPositiveNumber("-1.0"))

}

func TestIsPN(t *testing.T) {
	req := regexp.MustCompile(`^\d+\.?\d?$`)
	t.Log(req.MatchString("0"))
	t.Log(req.MatchString("12"))
	t.Log(req.MatchString("12.2"))
	t.Log(req.MatchString("12.0"))
	t.Log("----------------------------------")
	t.Log(req.MatchString("12."))
	t.Log(req.MatchString("12.d"))

}

func TestEfficientFloatToStringByNegative8(t *testing.T) {
	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringByNegative8(*NewFloat256ByStringMust("12.8")))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToStringByNegative8(*NewFloat256ByStringMust("-12.8")))
}

func TestEF(t *testing.T) {
	s := "1.9999999999"
	ff := NewFloat256ByStringMust(s)

	t.Log(ff.String())
	f64, _ := ff.Float64()
	t.Log(f64)
	t.Log(ff.Text('f', 25))
	t.Log(GetInstanceByNumberUtil().ParseFloatMust(ff.Text('f', 25), 8))

	bf := new(big.Float)
	bf.SetPrec(128)
	bf.SetString(s)
	t.Log(bf.Text('f', 8))

	t.Log("1.99保留8 =", GetInstanceByNumberUtil().EfficientFloatToPrecisionString(*NewFloat256ByStringMust("1.99"), 8))
	t.Log("10.0909保留8 =", GetInstanceByNumberUtil().EfficientFloatToPrecisionString(*NewFloat256ByStringMust("10.0909"), 8))

	t.Log(GetInstanceByNumberUtil().EfficientFloatToFloatBy8Must(*ff))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToFloatBy8Must(*ff))
	t.Log(GetInstanceByNumberUtil().EfficientFloatToPrecisionString(*ff, 8))

	pointIndex := strings.Index(s, ".")
	t.Log(fmt.Sprintf("pointIndex=%d", pointIndex))

	p := GetInstanceByNumberUtil().ParseFloatMust(s, 8)
	t.Log(s)
	t.Log(p) // 1.99999999
	t.Log(GetInstanceByNumberUtil().ParseFloatMust("1", 8))
	t.Log(GetInstanceByNumberUtil().ParseFloatMust("800", 8))
	t.Log(GetInstanceByNumberUtil().ParseFloatMust("1.9999", 8))
	t.Log(GetInstanceByNumberUtil().ParseFloatMust("0.5", 8))
	t.Log(GetInstanceByNumberUtil().ParseFloatMust("1.1111111111111111111111", 8))
	t.Log(GetInstanceByNumberUtil().ParseFloatMust("9.9999", 0))

	return
}
