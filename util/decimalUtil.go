package util

import "github.com/shopspring/decimal"

func DecimalToP0Str(a decimal.Decimal) string {
	return a.RoundDown(0).String()
}

func DecimalNewFromFloat(a float64) decimal.Decimal {
	return decimal.NewFromFloat(a)
}

func DecimalNewFromStringMust(a string) decimal.Decimal {
	d, _ := decimal.NewFromString(a)
	return d
}

func DecimalNewFromStringPanic(a string) decimal.Decimal {
	d, err := decimal.NewFromString(a)
	if nil != err {
		panic(err)
	}
	return d
}

func DecimalIsZero(a decimal.Decimal) bool {
	if "0" == a.String() {
		return true
	}
	return false
}

func DecimalLessOrEquZero(a decimal.Decimal) bool {
	v := DecimalNewFromStringMust("0")
	return a.LessThanOrEqual(v)
}

func DecimalSubRoundDown8Str(a, b string) (string, error) {
	r, err := DecimalSubStr(a, b)
	if nil != err {
		return "", err
	}
	return r.RoundDown(8).String(), nil
}

func DecimalSubStr(a, b string) (decimal.Decimal, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return c, err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return d, err
	}
	res := c.Sub(d)
	return res, nil
}

func DecimalAddRoundDown8Str(a, b string) (string, error) {
	r, err := DecimalAddStr(a, b)
	if nil != err {
		return "", err
	}
	return r.RoundDown(8).String(), nil
}

func DecimalAddStr(a, b string) (decimal.Decimal, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return c, err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return d, err
	}
	res := c.Add(d)
	return res, nil
}

// DecimalDivRoundDown8Str a/b = .8_0 string
func DecimalDivRoundDown8Str(a, b string) (string, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return "", err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return "", err
	}

	result := c.Div(d).RoundDown(8)
	return result.String(), nil
}

func DecimalGreaterThanOrEqualStrPanic(a, b string) bool {
	is, err := DecimalGreaterThanOrEqualStr(a, b)
	if nil != err {
		panic(err)
	}
	return is
}

func DecimalGreaterThanOrEqualStr(a, b string) (bool, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return false, err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return false, err
	}

	if c.GreaterThanOrEqual(d) {
		return true, nil
	}

	return false, nil
}

func DecimalLessThanStr(a, b string) (bool, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return false, err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return false, err
	}

	if c.LessThan(d) {
		return true, nil
	}

	return false, nil
}

func DecimalLessThanOrEqualStrPanic(a, b string) bool {
	is, err := DecimalLessThanOrEqualStr(a, b)
	if nil != err {
		panic(err)
	}
	return is
}

func DecimalLessThanOrEqualStr(a, b string) (bool, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return false, err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return false, err
	}

	if c.LessThanOrEqual(d) {
		return true, nil
	}

	return false, nil
}
