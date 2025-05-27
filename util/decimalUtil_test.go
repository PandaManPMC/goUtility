package util

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestInt(t *testing.T) {
	t.Log(DecimalNewFromStringMust("1").IntPart())
	t.Log(DecimalNewFromStringMust("1.22").IntPart())
	t.Log(DecimalNewFromStringMust("0.3333").IntPart())

	a := DecimalNewFromStringMust("1")
	b := DecimalNewFromStringMust("1")
	c := a.Sub(b)

	t.Log(a)
	t.Log(b)
	t.Log(c)

	t.Log("------------------------")
	t.Log(DecimalNewFromStringMust("0.0"))
	t.Log(DecimalIsZero(DecimalNewFromStringMust("0.0")))
	t.Log("------------------------")

	t.Log(DecimalLessOrEquZero(DecimalNewFromStringMust("0.1")))
	t.Log(DecimalLessOrEquZero(DecimalNewFromStringMust("0")))
	t.Log(DecimalLessOrEquZero(DecimalNewFromStringMust("-0.5")))

}

func TestDivRoundStr(t *testing.T) {
	t.Log(DecimalDivRoundDown8Str("10", "3"))

	t.Log("-----------")
	t.Log(DecimalGreaterThanOrEqualStr("1", "1"))
	t.Log(DecimalGreaterThanOrEqualStr("1", "0.3"))
	t.Log(DecimalGreaterThanOrEqualStr("1", "11"))
	t.Log("-----------")
	t.Log(DecimalLessThanStr("1", "1"))
	t.Log(DecimalLessThanStr("1", "0.3"))
	t.Log(DecimalLessThanStr("1", "11"))

	val := "145000000000000000000000000000.9"
	dec, err := decimal.NewFromString(val)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(dec.Add(DecimalNewFromStringMust("1")).Round(8).String())
	t.Log(dec.Add(DecimalNewFromStringMust("1")).Round(0).String())
	t.Log(dec.Add(DecimalNewFromStringMust("1")).RoundBank(0).String())
	t.Log(dec.Add(DecimalNewFromStringMust("1")).RoundDown(0).String())

	t.Log(DecimalAddStr(val, "1"))
	t.Log(DecimalSubStr(val, "1"))

	t.Log("-----------")

	t.Log(DecimalAddRoundDown8Str(val, "1"))
	t.Log(DecimalSubRoundDown8Str(val, "1"))
}
