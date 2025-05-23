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

func TestBigInt2(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(5)

	c := BigIntQuo(a, b)
	t.Log(c.String())

	t.Log(BigIntEquals(a, b))
	t.Log(BigIntEquals(a, big.NewInt(100)))

}
