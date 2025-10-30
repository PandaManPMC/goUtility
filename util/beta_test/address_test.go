package beta_test

import (
	"fmt"
	"goUtility/util"
	"goUtility/util/ethWal"
	"goUtility/util/tronWal"
	"testing"
)

func TestRandAddress(t *testing.T) {
	fmt.Println("ETH :", util.RandomETH())
	fmt.Println("SOL :", util.RandomSOL())
	fmt.Println("DOGE:", util.RandomDOGE())
	fmt.Println("LTC :", util.RandomLTC())
	fmt.Println("RVN :", util.RandomRVN())
	fmt.Println("TRX :", util.RandomTRX())

	t.Log(ethWal.ValidAddress(util.RandomETH()))
	t.Log(tronWal.ValidAddress(util.RandomTRX()))
}
