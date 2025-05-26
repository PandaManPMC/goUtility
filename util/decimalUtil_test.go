package util

import "testing"

func TestDivRoundStr(t *testing.T) {
	t.Log(DecimalDivRoundStr("10", "3"))

	t.Log("-----------")
	t.Log(DecimalGreaterThanOrEqualStr("1", "1"))
	t.Log(DecimalGreaterThanOrEqualStr("1", "0.3"))
	t.Log(DecimalGreaterThanOrEqualStr("1", "11"))
	t.Log("-----------")
	t.Log(DecimalLessThanStr("1", "1"))
	t.Log(DecimalLessThanStr("1", "0.3"))
	t.Log(DecimalLessThanStr("1", "11"))
}
