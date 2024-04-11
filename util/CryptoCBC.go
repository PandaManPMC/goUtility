package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"net/url"
)

func __pkcsUnPadding(text []byte) []byte {
	n := len(text)
	if n == 0 {
		return text
	}
	paddingSize := int(text[n-1])
	return text[:n-paddingSize]
}

func __pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

type CryptoCBC struct {
	block cipher.Block
	key   []byte
}

//	NewAESCryptoCBC 密钥 16位则 Pkcs5Padding 32位则 Pkcs7Padding
func NewAESCryptoCBC(key []byte) (*CryptoCBC, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	r := &CryptoCBC{
		block: b,
		key:   key,
	}
	return r, nil
}

func (that *CryptoCBC) EncryptWithIV(plainText []byte, iv []byte) []byte {
	plainText = __pkcs7Padding(plainText, that.block.BlockSize())
	cipherText := make([]byte, len(plainText))
	crypto := cipher.NewCBCEncrypter(that.block, iv)
	crypto.CryptBlocks(cipherText, plainText)
	return cipherText
}

//	EncryptWithIVToBase64 加密后结果转为 base64 字符串
func (that *CryptoCBC) EncryptWithIVToBase64(plainText []byte, iv []byte) string {
	buf := that.EncryptWithIV(plainText, iv)
	cipherText := base64.StdEncoding.EncodeToString(buf)
	return cipherText
}

//	EncryptWithIVToBase64URLEncode 加密后结果转 base64 并 URLEncode
func (that *CryptoCBC) EncryptWithIVToBase64URLEncode(plainText []byte, iv []byte) string {
	cipherText := that.EncryptWithIVToBase64(plainText, iv)
	return url.QueryEscape(cipherText)
}

//	DecryptWithIVByBase64Must base64 字符串解密，当 base64 解析错误时返回 err 不为 nil
func (that *CryptoCBC) DecryptWithIVByBase64Must(cipherText string, iv []byte) ([]byte, error) {
	cipherText, err1 := url.PathUnescape(cipherText)
	if nil != err1 {
		return nil, err1
	}
	s, err2 := base64.StdEncoding.DecodeString(cipherText)
	if nil != err2 {
		return nil, err2
	}
	plainBuf := that.DecryptWithIV(s, iv)
	return plainBuf, nil
}

//	DecryptWithIVByBase64URLEncodeMust base64转码和 URLEncode转码后的字符串解密，当 base64 解析错误时返回 err 不为 nil
func (that *CryptoCBC) DecryptWithIVByBase64URLEncodeMust(cipherText string, iv []byte) ([]byte, error) {
	t, err := url.QueryUnescape(cipherText)
	if nil != err {
		return nil, err
	}
	return that.DecryptWithIVByBase64Must(t, iv)
}

func (that *CryptoCBC) DecryptWithIV(cipherText []byte, iv []byte) []byte {
	plainText := make([]byte, len(cipherText))
	crypto := cipher.NewCBCDecrypter(that.block, iv)
	crypto.CryptBlocks(plainText, cipherText)
	plainText = __pkcsUnPadding(plainText)

	return plainText
}
