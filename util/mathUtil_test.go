package util

import "testing"

func TestCombination(t *testing.T) {
	t.Log(GetMathUtil().Combination(6, 4))
	t.Log(GetMathUtil().Combination(6, 0))
	t.Log(GetMathUtil().Combination(27, 6))
}
