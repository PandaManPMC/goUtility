package util

import (
	"math/big"
	"testing"
)

func TestBigIntAdd(t *testing.T) {
	a := big.NewInt(1)
	b := big.NewInt(5)

	c := BigIntAdd(a, b)
	t.Log(a.String())
	t.Log(b.String())
	t.Log(c.String())

	t.Log("-------------------------")

	d := BigIntSub(a, b)
	t.Log(a.String())
	t.Log(b.String())
	t.Log(d.String())

	t.Log("-------------------------")

	e := BigIntAddMul(a, b)
	t.Log(a.String())
	t.Log(b.String())
	t.Log(e.String())

	t.Log("-------------------------")

	f := BigIntAddMul100(a)
	t.Log(a.String())
	t.Log(f.String())

	t.Log("-------------------------")

}

func TestBigIntStr(t *testing.T) {

	t.Log(big.NewInt(0).SetString("1", 10))
	t.Log(big.NewInt(0).SetString("1.2", 10))
	t.Log(big.NewInt(0).SetString("0.5", 10))

}

func TestBigInt2(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(5)

	c := BigIntQuo(a, b)
	t.Log(c.String())

	t.Log(BigIntEquals(a, b))
	t.Log(BigIntEquals(a, big.NewInt(100)))

	t.Log("----------------- NewBigIntByStringPositive")
	t.Log(NewBigIntByStringPositive("1"))
	t.Log(NewBigIntByStringPositive("0"))
	t.Log(NewBigIntByStringPositive("1.4"))
	t.Log(NewBigIntByStringPositive("-1"))

	t.Log("----------------- BigIntLessOrEqualsZero")

	t.Log(BigIntLessOrEqualsZero(big.NewInt(10)))
	t.Log(BigIntLessOrEqualsZero(big.NewInt(0)))
	t.Log(BigIntLessOrEqualsZero(big.NewInt(-1)))
	t.Log("----------------- BigIntBiggerThanZero")
	t.Log(BigIntBiggerThanZero(big.NewInt(10)))
	t.Log(BigIntBiggerThanZero(big.NewInt(0)))
	t.Log(BigIntBiggerThanZero(big.NewInt(-1)))

}
