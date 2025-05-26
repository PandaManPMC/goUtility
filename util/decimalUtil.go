package util

import "github.com/shopspring/decimal"

// DecimalDivRoundStr a/b = .8 string
func DecimalDivRoundStr(a, b string) (string, error) {
	c, err := decimal.NewFromString(a)
	if nil != err {
		return "", err
	}
	d, err := decimal.NewFromString(b)
	if nil != err {
		return "", err
	}

	result := c.Div(d).Round(8)
	return result.String(), nil
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
