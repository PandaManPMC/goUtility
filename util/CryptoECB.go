package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

type CryptoECB struct {
	block cipher.Block
	key   []byte
}

const (
	CryptModelPkcs5 = "Pkcs5"
	CryptModelPkcs7 = "Pkcs7"
)

func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	return base64.URLEncoding.DecodeString(data)
}

func Base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain ‘/‘ and ‘+‘ (replaced by ‘_’ and ‘-’) and trailing ‘=‘ are removed.
	byteArr := base64.StdEncoding.EncodeToString(source)
	safeUrl := strings.Replace(string(byteArr), "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return safeUrl
}

func AesDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := NewECBDecrypt(block)
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func AesEncrypt(src string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ecb := NewECBEncrypt(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	encrypted := make([]byte, len(content))
	ecb.CryptBlocks(encrypted, content)
	return encrypted, nil
}

func AesEcbPkcs7Encode(src string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ecb := NewECBEncrypt(block)
	content := []byte(src)
	content = pkcs7Pad(content, block.BlockSize())
	encrypted := make([]byte, len(content))
	ecb.CryptBlocks(encrypted, content)
	return encrypted, nil
}

func AesEcbPkcs7Decode(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := NewECBDecrypt(block)
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)
	origData = pkcs7Unpad(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// remove the last byte unPadding times
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypt ecb

// NewECBEncrypt returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypt)(newECB(b))
}

func (x *ecbEncrypt) BlockSize() int { return x.blockSize }

func (x *ecbEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto / cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto / cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypt ecb

// NewECBDecrypt returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypt)(newECB(b))
}

func (x *ecbDecrypt) BlockSize() int { return x.blockSize }

func (x *ecbDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto / cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto / cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// pkcs7Pad right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func pkcs7Pad(b []byte, blocksize int) []byte {
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7Unpad(b []byte) []byte {
	c := b[len(b)-1]
	n := int(c)
	return b[:len(b)-n]
}
