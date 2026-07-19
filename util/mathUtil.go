package util

type mathUtil struct {
}

var mathUtilInstance mathUtil

func GetMathUtil() *mathUtil {
	return &mathUtilInstance
}

// Combination 组合数计算 C(n,k)
func (that *mathUtil) Combination(n, k int64) int64 {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}

	var res int64 = 1
	for i := int64(1); i <= k; i++ {
		res = res * (n - k + i) / i
	}
	return res
}
