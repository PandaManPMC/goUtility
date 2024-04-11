package util

import (
	"encoding/base64"
	"testing"
)

func TestDesEncrypt(t *testing.T) {
	encryptKey := "12341234"
	originalText := `cagent=81288128/\\\\/method=tc`
	if encryptedMessage, err := GetInstanceByDesUtil().DesPkcs5Encryption([]byte(encryptKey), []byte(originalText)); nil != err {
		t.Logf("error: %+v", err)
	} else {
		t.Logf("crypted string: %s", base64.StdEncoding.EncodeToString(encryptedMessage))
	}
}

func TestDesDecrypt(t *testing.T) {
	encryptKey := "12341234"
	originalTextBase64 := `IGcOAYEQN88F1NFLtBOK29IcQSW2a8b/G8UgSaeEyaA=`
	if originalText, err := base64.StdEncoding.DecodeString(originalTextBase64); nil != err {
		t.Logf("base64 decode error: %+v", err)
	} else if encryptedMessage, err := GetInstanceByDesUtil().DesPkcs5Decryption([]byte(encryptKey), originalText); nil != err {
		t.Logf("error: %+v", err)
	} else {
		t.Logf("crypted string: %s", string(encryptedMessage))
	}
}
