package util

import (
	"crypto/des"
)

type desUtil struct {
}

var desUtilInstance desUtil

func GetInstanceByDesUtil() *desUtil {
	return &desUtilInstance
}

func (that *desUtil) DesPkcs5Encryption(key, plainText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ecbEncryptIns := NewECBEncrypt(block)
	content := PKCS5Padding(plainText, block.BlockSize())
	encrypted := make([]byte, len(content))
	ecbEncryptIns.CryptBlocks(encrypted, content)
	return encrypted, nil
}

func (that *desUtil) DesPkcs5Decryption(key, cipherText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := NewECBDecrypt(block)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
