package tronWal

import (
	"goUtility/util"
	"testing"
)

func TestHexToTronAddress(t *testing.T) {
	//tronWeb.address.fromHex("418840E6C55B9ADA326D211D818C34A994AECED808")
	//> "TNPeeaaFB7K9cmo4uQpcU32zGK8G1NYqeL"
	s := "418840E6C55B9ADA326D211D818C34A994AECED808"

	a, e := HexToTronAddress(s)
	if nil != e {
		t.Fatal(e)
	}
	t.Log(a)
}

func TestValid(t *testing.T) {
	t.Log(ValidAddress("TKjPqUwwT9Q4VPKH9BMKvgafxEyDWT9t8Z"))
}

func TestMnemonic(t *testing.T) {

	mnemonic, err := util.GetInstanceByHDWalletUtil().GenerateMnemonicBy12()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	mnemonic, err = util.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	hd, err := util.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	pri, err := util.GetInstanceByHDWalletUtil().WalletPrivateKey(hd, 0)
	if nil != err {
		t.Fatal(err)
	}

	pri_s := util.GetInstanceByHDWalletUtil().PriKeyToHexString(pri)
	t.Log(pri_s)

	addr := TronAddressByPrivateKey(pri)
	t.Log(addr)

	l_pri, err := util.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(pri_s)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(l_pri)

	addr = PriKeyToAddressTron(l_pri)
	t.Log(addr)
}

func TestMnemonic2(t *testing.T) {

	mnemonic, err := util.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	pri, addr, err := ImportWallet(mnemonic, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(addr)

	pri_s := util.GetInstanceByHDWalletUtil().PriKeyToHexString(pri)
	t.Log(pri_s)

	addr = TronAddressByPrivateKey(pri)
	t.Log(addr)

	l_pri, err := util.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(pri_s)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(l_pri)

	addr = PriKeyToAddressTron(l_pri)
	t.Log(addr)
}
