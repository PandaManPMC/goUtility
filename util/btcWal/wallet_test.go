package btcWal

import (
	"goUtility/util"
	"testing"
)

func TestHDWallet(t *testing.T) {
	mne := "gown super smile wing hunt keep carpet stereo nurse umbrella case gun list fun valve stock job debate drip angry dumb tree finish lend"
	hd, err := util.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mne)
	if nil != err {
		t.Fatal(err)
	}

	wallet, err := util.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(hd, 3, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(wallet)

	addr, err := GenerateAddress(wallet, 0x1E)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(addr) // DJo1dayr1qmRkkmVMHomY1uf3a2neU9fTc
}

func TestImportWallet(t *testing.T) {
	mne, err := util.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mne)
	//mne := "gown super smile wing hunt keep carpet stereo nurse umbrella case gun list fun valve stock job debate drip angry dumb tree finish lend"
	pk, addr, err := ImportWallet(mne, 3, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(pk)
	t.Log(addr)
	t.Log(IsValidDOGEAddress(addr))
}

func TestHDWallet2(t *testing.T) {
	coinType := util.DOGEHDCoinType
	addressType := DogeAddress

	mnemonic, err := util.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	wallet, err := util.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	privateKey, err := util.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(wallet, coinType, 0)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyStr := util.GetInstanceByHDWalletUtil().PriKeyToHexString(privateKey)

	t.Log(privateKeyStr)

	addr, err := GenerateAddress(privateKey, addressType)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))

	//wallet_test.go:115: flat globe fuel spend brother tornado pistol remember survey bless spread kitchen hill current deny rail crisp witness siren elder office leg aware seed
	//wallet_test.go:128: 1ea107cf1e8cbca5a1e9ee2661505b1836db495c574eace64ecdbc20b29b83fd
	//wallet_test.go:135: DNyqb4rfhb3niqVGJm3tUjcZzZoybNNoxf
	//wallet_test.go:136: true

}

func TestGenerateAddressCompressed(t *testing.T) {
	coinType := util.DOGEHDCoinType
	addressType := DogeAddress

	mnemonic, err := util.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	wallet, err := util.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	privateKey, err := util.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(wallet, coinType, 0)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyStr := util.GetInstanceByHDWalletUtil().PriKeyToHexString(privateKey)

	t.Log(privateKeyStr)

	addr, err := GenerateAddressCompressed(privateKey, addressType)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))
}

func TestPrivateKetToAddr(t *testing.T) {
	privateKey := "1ea107cf1e8cbca5a1e9ee2661505b1836db495c574eace64ecdbc20b29b83fd"

	pk, err := util.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(privateKey)
	if nil != err {
		t.Fatal(err)
	}

	addr, err := GenerateAddress(pk, LTCAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log(IsValidDOGEAddress(addr))

}
