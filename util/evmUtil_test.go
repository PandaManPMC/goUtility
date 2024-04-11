package util

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestSign(t *testing.T) {
	//地址0=0x97Fa4A2d3C3df5b3b05F4B9265F438c3a8B064eb
	//私钥=ec8da26e3e59fa40ac0410fb09df8ff05bd7ebad5ef5fa89ffeea7408ac9739f
	//助记词=sister pyramid polar oyster describe empty unknown night ill youth arrow awkward
	//公钥=8ea0bc55ea26af995f760edde7a07c86faa6404d8f9dc9915b8b9ea30c058f8fbe2172315265a21388185f7ce34db0ddc765401e0a310cd9a609458aa139091c

	address := "0x97Fa4A2d3C3df5b3b05F4B9265F438c3a8B064eb"
	//pubKey := "8ea0bc55ea26af995f760edde7a07c86faa6404d8f9dc9915b8b9ea30c058f8fbe2172315265a21388185f7ce34db0ddc765401e0a310cd9a609458aa139091c"
	privateKey := "ec8da26e3e59fa40ac0410fb09df8ff05bd7ebad5ef5fa89ffeea7408ac9739f"

	// sign
	data := "hello test sign"
	//dataHash := encrypt.Keccak256Hash([]byte(data))
	dataHash := crypto.Keccak256([]byte(data))
	fmt.Println("daraHash=", dataHash)

	pkey, err := crypto.HexToECDSA(privateKey)
	if nil != err {
		t.Fatal(err)
	}
	// 签名
	signature, err := crypto.Sign(dataHash, pkey)
	if nil != err {
		t.Fatal(err)
	}
	sigHex := hexutil.Encode(signature)
	fmt.Println("签名长度:", len(signature))
	fmt.Println("签名:", sigHex)

	// 从签名推出公钥
	sigByte, _ := hexutil.Decode(sigHex)
	recoveredPub, err := crypto.Ecrecover(dataHash, sigByte)
	if nil != err {
		fmt.Println(err)
		return
	}
	recoveredPubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*recoveredPubKey)
	t.Log("钱包地址：", recoveredAddr)
	t.Log("公钥", hex.EncodeToString(recoveredPub))
	if address == recoveredAddr.String() {
		t.Log("地址比较一致")
	}

	// VerifySignature 验证签名
	publicKeyBytes := crypto.FromECDSAPub(recoveredPubKey)
	ok := crypto.VerifySignature(publicKeyBytes, dataHash[:], sigByte[:len(sigByte)-1])
	fmt.Println("签名结果：", ok)

	pubK := "048ea0bc55ea26af995f760edde7a07c86faa6404d8f9dc9915b8b9ea30c058f8fbe2172315265a21388185f7ce34db0ddc765401e0a310cd9a609458aa139091c"
	k, _ := hex.DecodeString(pubK)
	ecdsaKey, err := crypto.UnmarshalPubkey(k)
	t.Log(err)
	t.Log(ecdsaKey)
	ecdsaKeyBytes := crypto.FromECDSAPub(ecdsaKey)
	ok = crypto.VerifySignature(ecdsaKeyBytes, dataHash[:], sigByte[:len(sigByte)-1])
	fmt.Println("签名结果：", ok)
}

func TestVerifySignature(t *testing.T) {
	address := "0x97Fa4A2d3C3df5b3b05F4B9265F438c3a8B064eb"
	//pubKey := "8ea0bc55ea26af995f760edde7a07c86faa6404d8f9dc9915b8b9ea30c058f8fbe2172315265a21388185f7ce34db0ddc765401e0a310cd9a609458aa139091c"
	privateKey := "ec8da26e3e59fa40ac0410fb09df8ff05bd7ebad5ef5fa89ffeea7408ac9739f"

	// sign
	data := "hello test sign"
	sigHex, err := GetEvmUtil().Sign(data, privateKey)
	if nil != err {
		t.Fatal(err)
	}

	// VerifySignature 验证签名
	address0, ok, err := GetEvmUtil().VerifySignature(data, sigHex)
	if nil != err {
		t.Fatal(err)
	}
	if ok {
		t.Log("签名验证成功")
	}
	t.Logf("地址比较：%s==%s=%v", address0, address, address0 == address)
}

func TestVerifySignature2(t *testing.T) {
	// sign
	data := "abcdef"
	dataHash := crypto.Keccak256([]byte(data))
	t.Log(dataHash)

	address := "0x6a2423a19d2e2f939f908fa5ff325337194953cf"
	sigHex := "0x02ad6df89771cfa7f08c1b8c82734b850e249443be468b803a5d0a09573702625158e174ca8252aaab50655d769f203bc25c14bc2d90a8481851b590313c80421c"
	// VerifySignature 验证签名
	address0, ok, err := GetEvmUtil().VerifySignature(data, sigHex)
	if nil != err {
		t.Fatal(err)
	}
	if ok {
		t.Log("签名验证成功")
	}
	t.Logf("地址比较：%s==%s=%v", address0, address, address0 == address)

}

func TestVerifyTypedDataHexSignatureEx(t *testing.T) {
	data := "sign data 201020192"
	address := "0x6a2423a19d2e2f939f908fa5ff325337194953cf"
	sigHex := "0x764d625e5de77b4be71ef65faf77e51804e81628e1bdd8ef81efa5e62aca36082d7cfff561671ffa6251516a219fe745df36ba42ef0557a01d4e91c2c2591ef11b"

	valid, err := GetEvmUtil().VerifyEllipticCurveHexSignatureEx(address, []byte(data), sigHex)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(valid)
}
