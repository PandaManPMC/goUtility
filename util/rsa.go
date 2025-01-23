package util

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type RSA struct {
	PKCSModel RSAPemPKCSModel
}

type RSAPemPKCSModel string
type RSABits int

const (
	RSAPemPKCS1 RSAPemPKCSModel = "PKCS1"
	RSAPemPKCS8 RSAPemPKCSModel = "PKCS8"
	RSA1024     RSABits         = 1024
	RSA2048     RSABits         = 2048
)

func NewRSA(pemModel RSAPemPKCSModel) *RSA {
	return &RSA{PKCSModel: pemModel}
}

func NewRSAPemPKCS1() *RSA {
	return &RSA{PKCSModel: RSAPemPKCS1}
}

func NewRSAPemPKCS8() *RSA {
	return &RSA{PKCSModel: RSAPemPKCS8}
}

func (that *RSA) GenRsaKey1024() (prvKey, pubKey []byte, err error) {
	return that.GenRsaKey(RSA1024)
}

func (that *RSA) GenRsaKey1024ToBase64() (prvKey, pubKey string, err error) {
	prvKeyB, pubKeyB, err := that.GenRsaKey(RSA1024)
	if nil != err {
		return "", "", err
	}
	prvKey = base64.StdEncoding.EncodeToString(prvKeyB)
	pubKey = base64.StdEncoding.EncodeToString(pubKeyB)
	return
}

func (that *RSA) GenRsaKey2048() (prvKey, pubKey []byte, err error) {
	return that.GenRsaKey(RSA2048)
}

func (that *RSA) GenRsaKey2048ToBase64() (prvKey, pubKey string, err error) {
	prvKeyB, pubKeyB, err := that.GenRsaKey(RSA2048)
	if nil != err {
		return "", "", err
	}
	prvKey = base64.StdEncoding.EncodeToString(prvKeyB)
	pubKey = base64.StdEncoding.EncodeToString(pubKeyB)
	return
}

// GenRsaKey	RSA公钥私钥产生
func (that *RSA) GenRsaKey(bits RSABits) (prvKey, pubKey []byte, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, int(bits))
	if nil != err {
		return
	}

	var derStream []byte
	if RSAPemPKCS1 == that.PKCSModel {
		derStream = x509.MarshalPKCS1PrivateKey(privateKey)
	} else {
		derStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if nil != err {
			return nil, nil, err
		}
	}

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvKey = pem.EncodeToMemory(block)

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if nil != err {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey = pem.EncodeToMemory(block)
	return
}

// RsaSignWithSha256 签名:用私钥生成签名，用公钥验证签名，确定发送者的是不是目标身份。
//
//	内部会对数据进行 sha256 一次，因此外部可以不做处理。
func (that *RSA) RsaSignWithSha256(data []byte, privateKeyByte []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyByte)
	var keyBytes []byte
	if nil == block {
		//return nil, errors.New("private key error")
		keyBytes = privateKeyByte
	} else {
		keyBytes = block.Bytes
	}

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	var privateKey *rsa.PrivateKey
	var err error
	if RSAPemPKCS1 == that.PKCSModel {
		privateKey, err = x509.ParsePKCS1PrivateKey(keyBytes)
	} else {
		var key any
		key, err = x509.ParsePKCS8PrivateKey(keyBytes)
		privateKey = key.(*rsa.PrivateKey)
	}
	if nil != err {
		return nil, err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if nil != err {
		return nil, err
	}
	return signature, nil
}

// RsaSignWithMD5 私钥对数据签名
// 内部对数据 md5 一次，外面不做操作
func (that *RSA) RsaSignWithMD5(data []byte, privateKeyByte []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyByte)
	var keyBytes []byte
	if nil == block {
		//return nil, errors.New("private key error")
		keyBytes = privateKeyByte
	} else {
		keyBytes = block.Bytes
	}

	h := md5.New()
	h.Write(data)
	hash := h.Sum(nil)

	var privateKey *rsa.PrivateKey
	var err error
	if RSAPemPKCS1 == that.PKCSModel {
		privateKey, err = x509.ParsePKCS1PrivateKey(keyBytes)
	} else {
		var key any
		key, err = x509.ParsePKCS8PrivateKey(keyBytes)
		privateKey = key.(*rsa.PrivateKey)
	}
	if nil != err {
		return nil, err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hash)
	if nil != err {
		return nil, err
	}
	return signature, nil
}

// RsaVerySignWithSha256 验证签名
func (*RSA) RsaVerySignWithSha256(data, signData, pubKeyBytes []byte) (bool, error) {
	block, _ := pem.Decode(pubKeyBytes)
	var keyBytes []byte
	if nil == block {
		// 认为不是 pem 则直接使用 pubKey
		//return false, errors.New("public key error")
		keyBytes = pubKeyBytes
	} else {
		keyBytes = block.Bytes
	}
	pubKey, err := x509.ParsePKIXPublicKey(keyBytes)
	if nil != err {
		return false, err
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if nil != err {
		return false, err
	}
	return true, nil
}

// RsaVerySignWithMD5 验证签名
func (*RSA) RsaVerySignWithMD5(data, signData, pubKeyBytes []byte) (bool, error) {
	block, _ := pem.Decode(pubKeyBytes)
	var keyBytes []byte
	if nil == block {
		//return false, errors.New("public key error")
		keyBytes = pubKeyBytes
	} else {
		keyBytes = block.Bytes
	}
	pubKey, err := x509.ParsePKIXPublicKey(keyBytes)
	if nil != err {
		return false, err
	}

	hashed := md5.Sum(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.MD5, hashed[:], signData)
	if nil != err {
		return false, err
	}
	return true, nil
}

// RsaEncrypt 公钥加密
func (*RSA) RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if nil == block {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if nil != err {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if nil != err {
		return nil, err
	}
	return ciphertext, nil
}

// RsaDecrypt 私钥解密
func (that *RSA) RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	var priKey *rsa.PrivateKey
	var err error
	if RSAPemPKCS1 == that.PKCSModel {
		priKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	} else {
		var key any
		key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		priKey = key.(*rsa.PrivateKey)
	}

	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}
