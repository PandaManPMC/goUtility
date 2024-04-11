package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
)

type messageDigest struct{}

var messageDigestInstance messageDigest

func GetInstanceByMessageDigest() *messageDigest {
	return &messageDigestInstance
}

// MD5 加密
// plaintext string	明文
// string	密文，32个字符16进制string
func (*messageDigest) MD5(plaintext string) string {
	m := md5.New()
	_, err := io.WriteString(m, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func (*messageDigest) MD5ToByte(plaintext string) []byte {
	m := md5.New()
	_, err := io.WriteString(m, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return arr
}

// MD5ByteArray 数据摘要 []byte
func (*messageDigest) MD5ByteArray(buf []byte) string {
	m := md5.New()
	m.Write(buf)
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

// MD5Double 两次MD5加密
// plaintext string	明文
// string	密文，32个字符16进制string
func (that *messageDigest) MD5Double(plaintext string) string {
	return that.MD5(that.MD5(plaintext))
}

// Sha1 加密
// plaintext string	明文
func (*messageDigest) Sha1(plaintext string) string {
	m := sha1.New()
	m.Write([]byte(plaintext))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// Sha1InBase64 将Sha1以Base64输出
// plaintext string	明文
func (*messageDigest) Sha1InBase64(plaintext string) string {
	m := sha1.New()
	m.Write([]byte(plaintext))
	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}

// Sha256 加密
// plaintext string	明文
// string	密文，64个字符16进制string
func (*messageDigest) Sha256(plaintext string) string {
	m := sha256.New()
	m.Write([]byte(plaintext))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// Sha256Bytes 加密
func (*messageDigest) Sha256Bytes(data []byte) string {
	m := sha256.New()
	m.Write(data)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func (*messageDigest) Sha256Buf(plaintext string) []byte {
	m := sha256.New()
	m.Write([]byte(plaintext))
	return m.Sum(nil)
}

func (that *messageDigest) HmacSha256(data string, secret []byte) string {
	return that.HmacSha256Buf([]byte(data), secret)
}

func (*messageDigest) HmacSha256Buf(data []byte, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func (that *messageDigest) HmacSha1InBytes(key []byte, input []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(input)
	return h.Sum(nil)
}

func (that *messageDigest) HmacSha1InBase64(key []byte, input []byte) string {
	return base64.StdEncoding.EncodeToString(that.HmacSha1InBytes(key, input))
}

func (that *messageDigest) AesEcbEncodeInBytes(data string, key []byte, pkcsModel string) ([]byte, error) {
	switch pkcsModel {
	case CryptModelPkcs5:
		return AesEncrypt(data, key)
	case CryptModelPkcs7:
		return AesEcbPkcs7Encode(data, key)
	}

	return nil, errors.New("cannot find aes model")
}

func (that *messageDigest) AesEcbEncodeInBase64(data string, key []byte, pkcsModel string) (string, error) {
	cryptBytes, err := that.AesEcbEncodeInBytes(data, key, pkcsModel)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cryptBytes), nil
}

func (that *messageDigest) DesEcbEncodeInBytes(data, key string, pkcsModel string) ([]byte, error) {
	switch pkcsModel {
	case CryptModelPkcs5:
		return GetInstanceByDesUtil().DesPkcs5Encryption([]byte(key), []byte(data))
	}
	return nil, errors.New(fmt.Sprintf("cannot find des model: %s", pkcsModel))
}

func (that *messageDigest) DesEcbEncodeInBase64(data, key string, pkcsModel string) (string, error) {
	cryptBytes, err := that.DesEcbEncodeInBytes(data, key, pkcsModel)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cryptBytes), nil
}

func (that *messageDigest) AesEcbDecodeFromBytes(data, key []byte, pkcsModel string) ([]byte, error) {
	switch pkcsModel {
	case CryptModelPkcs5:
		return AesDecrypt(data, key)
	case CryptModelPkcs7:
		return AesEcbPkcs7Decode(data, key)
	}

	return nil, errors.New("cannot find aes model")
}

func (that *messageDigest) AesEcbDecodeFromBase64(base64String string, key []byte, pkcsModel string) ([]byte, error) {
	base64DecodeBytes, err := base64.StdEncoding.DecodeString(base64String)
	if nil != err {
		return nil, err
	}
	decryptBytes, err := that.AesEcbDecodeFromBytes(base64DecodeBytes, key, pkcsModel)
	if err != nil {
		return nil, err
	}
	return decryptBytes, nil
}

func (that *messageDigest) DesEcbDecodeFromBytes(data []byte, key string, pkcsModel string) ([]byte, error) {
	switch pkcsModel {
	case CryptModelPkcs5:
		return GetInstanceByDesUtil().DesPkcs5Decryption([]byte(key), data)
	}
	return nil, errors.New(fmt.Sprintf("cannot find des model: %s", pkcsModel))
}

func (that *messageDigest) DesEcbDecodeFromBase64(base64String, key string, pkcsModel string) ([]byte, error) {
	base64DecodeBytes, err := base64.StdEncoding.DecodeString(base64String)
	if nil != err {
		return nil, err
	}
	decryptBytes, err := that.DesEcbDecodeFromBytes(base64DecodeBytes, key, pkcsModel)
	if err != nil {
		return nil, err
	}
	return decryptBytes, nil
}
