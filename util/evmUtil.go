package util

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type evmUtil struct {
}

var evmUtilInstance evmUtil

//	GetEvmUtil evmUtil evm 系列工具，包括签名等
func GetEvmUtil() *evmUtil {
	return &evmUtilInstance
}

//	Sign 签名
//	return *string sigHex 签名，如果 error 不为 nil 则失败
func (that *evmUtil) Sign(data, privateKey string) (string, error) {
	dataByte := crypto.Keccak256([]byte(data))
	pkey, err := crypto.HexToECDSA(privateKey)
	if nil != err {
		return "", err
	}
	signature, err := crypto.Sign(dataByte, pkey)
	if nil != err {
		return "", err
	}
	sigHex := hexutil.Encode(signature)
	return sigHex, err
}

//	VerifySignature 验证签名
//	return *string address 钱包地址, *bool 验签成功 true,error 不为 nil 则出现异常
func (that *evmUtil) VerifySignature(data, signature string) (string, bool, error) {
	dataHash := crypto.Keccak256([]byte(data))
	sigByte, err := hexutil.Decode(signature)
	if nil != err {
		return "", false, err
	}

	recoveredPub, err := crypto.Ecrecover(dataHash, sigByte)
	if nil != err {
		return "", false, err
	}
	recoveredPubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	address := crypto.PubkeyToAddress(*recoveredPubKey).String()

	publicKeyBytes := crypto.FromECDSAPub(recoveredPubKey)
	ok := crypto.VerifySignature(publicKeyBytes, dataHash[:], sigByte[:len(sigByte)-1])
	return address, ok, nil
}

// VerifyEllipticCurveHexSignatureEx is used to verify elliptic curve signatures
// It calls the EcRecoverEx function to verify the signature.
func (that *evmUtil) VerifyEllipticCurveHexSignatureEx(address string, data []byte, signature string) (bool, error) {
	if nil == data {
		return false, nil
	}
	sig, err := that.HexDecode(signature)
	if nil != err {
		return false, err
	}
	return that.VerifyEllipticCurveSignatureEx(ethCommon.HexToAddress(address), data, sig)
}

// VerifyEllipticCurveSignatureEx is used to verify elliptic curve signatures
// It calls the EcRecoverEx function to verify the signature.
func (that *evmUtil) VerifyEllipticCurveSignatureEx(address ethCommon.Address, data []byte, signature []byte) (bool, error) {
	recovered, err := that.EcRecoverEx(data, signature)
	if err != nil {
		return false, err
	}
	return recovered == address, nil
}

func (that *evmUtil) EcRecoverEx(data []byte, sig []byte) (ethCommon.Address, error) {
	return that.RecoveryAddressEx(accounts.TextHash(data), sig)
}

// RecoveryAddressEx is an extension to RecoveryAddress that supports more signature formats, such as ledger signatures.
func (that *evmUtil) RecoveryAddressEx(data []byte, sig []byte) (ethCommon.Address, error) {
	sig = that.CopyBytes(sig)
	if len(sig) != crypto.SignatureLength {
		return ethCommon.Address{}, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}
	// comment(storyicon): fix ledger wallet
	// https://ethereum.stackexchange.com/questions/103307/cannot-verifiy-a-signature-produced-by-ledger-in-solidity-using-ecrecover
	if sig[crypto.RecoveryIDOffset] == 0 || sig[crypto.RecoveryIDOffset] == 1 {
		sig[crypto.RecoveryIDOffset] += 27
	}
	return that.RecoveryAddress(data, sig)
}

func (that *evmUtil) RecoveryAddress(data []byte, sig []byte) (ethCommon.Address, error) {
	sig = that.CopyBytes(sig)
	if len(sig) != crypto.SignatureLength {
		return ethCommon.Address{}, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return ethCommon.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.SigToPub(data, sig)
	if err != nil {
		return ethCommon.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}

// HexDecode returns the bytes represented by the hexadecimal string s.
// s may be prefixed with "0x".
func (that *evmUtil) HexDecode(s string) ([]byte, error) {
	if that.Has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}

// Has0xPrefix validates str begins with '0x' or '0X'.
func (that *evmUtil) Has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// CopyBytes is used to copy slice
func (that *evmUtil) CopyBytes(data []byte) []byte {
	ret := make([]byte, len(data))
	copy(ret, data)
	return ret
}
