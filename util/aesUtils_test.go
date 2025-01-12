package util

import "testing"

func TestAesEncrypt(t *testing.T) {
	salt := "faAKHysMSSPyZz4qb3RJB1HMxMgptMhj"
	password := "JBSWY3DPEHPK3PXP"
	encryptMessage := GetInstanceByAesUtil().AesEncryptWithSalt(salt, password)
	t.Logf("encrypt message: %v", encryptMessage)
	t.Logf("decrypt message: %v", GetInstanceByAesUtil().AesDecryptWithSalt(salt, encryptMessage))
}
