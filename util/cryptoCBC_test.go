package util

import (
	"encoding/base64"
	"net/url"
	"testing"
)

func TestEncryptWithIV(t *testing.T) {
	j := `{"account":"kkkqqq3","nickname":"123"}`
	secretKey := "09936CD67213466DB4EF67DC0E92C4B5"

	//"HykjzxldgwfSIYc6nyvlUxkqmCR8TEzeyYVxQmAXvk6ZzHtu38ReQml6F9oMI2cc"

	cbc, err := NewAESCryptoCBC([]byte(secretKey))
	if nil != err {
		t.Fatal(err)
	}

	pwd := secretKey[:16]

	// 测试加密
	ct := cbc.EncryptWithIV([]byte(j), []byte(pwd))
	t.Log(string(ct))
	urlCt := base64.StdEncoding.EncodeToString(ct)
	t.Log(urlCt)

	//	测试解密
	dUrlCt, err2 := base64.StdEncoding.DecodeString(urlCt)
	if nil != err2 {
		t.Fatal(err2)
	}
	t.Log(string(dUrlCt))
	txt := cbc.DecryptWithIV(dUrlCt, []byte(pwd))
	t.Log(string(txt))
}

func TestEncryptWithIV7(t *testing.T) {
	plainText := `{"account":"kkkqqq3","nickname":"123"}`
	//secretKey := "09936CD67213466DB4EF67DC0E92C4B5"
	secretKey := "D2434fgdgfdgfg121"

	cbc, err1 := NewAESCryptoCBC([]byte(secretKey[:16]))
	if nil != err1 {
		t.Fatal(err1)
	}

	cipherText := cbc.EncryptWithIV([]byte(plainText), []byte(secretKey[:16]))
	t.Log(string(cipherText))

	cipherTextB64 := cbc.EncryptWithIVToBase64([]byte(plainText), []byte(secretKey[:16]))
	t.Log(cipherTextB64)

	cipherTextURL := base64.URLEncoding.EncodeToString([]byte(cipherTextB64))
	t.Log(cipherTextURL)

	t.Log("------------------------------------------")
	cipherTextURL = cbc.EncryptWithIVToBase64URLEncode([]byte(plainText), []byte(secretKey[:16]))
	t.Log(cipherTextURL)

}

func TestDecryptWithIV(t *testing.T) {
	s := `qCNc%2Bhh1Dj9P79cVyHX70shCUknXw7tFrKWu3H0kv0UhVTjo8CiOns5N8MJzQCD2S%2FUhqEE9ycS1GHw3arNBz1bVTkKVil50aem1w5271i3fdUBo7LthAtyiM3pTw1anEeLNz0guhjQQi02E9KzKTyADm4Ljjc9sUXxwT7eM%2Fi0UpRmJei0t6guTlBogU
6vHkbLGJ67wExXAFZaEnPyMUtEQ%2B%2FxhKA70fqvNjeWBf0dbnAjU%2ByK5UKVKDOPsIkYQyK3kfVhES6IvKXUMV359bW2oYSLv7iINWe8p03YKZ0mMFOTld5PQLKirsMFCKAbNDJTcJsr%2Fp%2BktHgVE8PRA9ZvaBVvwkIh7mdSlNTPapwJN2JNdjv1akyH0Vd250vmxlN6k182ltFzjRkM2fUCTSVbw
KY90LcuWnBm9aWPVaEyfVPhw1jvj5P6fYTdDM6z3F31BYno2yBWs%2BWK89q8yfv%2FE4QpS8VAqnvyqI1O%2BQfF7QodfwD3e8OwwTgPsochyxnw9INDVSQrmfP%2Foe4D8BcxQjCCh3VLrNLInSDu095bS0raD9PzVbFPlq%2FdNiIqFu7lT8ZKl0njyb2d4UxfavDDf0aI8GDxPCchITqxGW1xivsL7xDH
571cZDf5TSEK7B5ksVd5Jbgt4I4qd1XL1G%2FPBpff5WYix8uEY8Fy50xNM4Hk8LP7qv8Mlau%2BsHTPgaVDtNt5Jli6XMLp%2BJSr6aVkaiyxVSWWoqaR34HnMwjmX7KFRaDKmOGt6LrNC8rnlkRZXB%2Ba9xQ5MbFWb5mJGfSLSgRYjm4BPkpk46cUBe7SoIZTZkM00NLwtrcMnUvjizlsAUvHbxgamggbC
6BgjMUsS8jeeJZhWlrvypXzj%2FlkIMk%2FxT15z8DqYHkZF0c0ZynvhhJAPxeN8AlY4z0pOk5jutP9S5GlKoLZdzSKztg3pzmaXUjF7ScNz0YhyFK03KyHCT0ehH7S%2Fb8Z%2B1yL%2BKcRo6dVcBI9SRiOc7vAcUqX40BrbjjbiB%2BJpSm23frV557%2FXAu6buRNTtB7VDUC9rqnwWJCQU9dmLAJIlQm
cF47Bu%2B%2FTLLo9Dy2gGCVWAcIUCCcZwDuikWQC0jItiiU6AQ%3D%3D`
	//buf, err1 := base64.StdEncoding.DecodeString(s)
	buf, err1 := url.QueryUnescape(s)
	//buf, err1 := base64.URLEncoding.DecodeString(s)
	if nil != err1 {
		t.Fatal(err1)
	}
	t.Log(buf)

	dBuf, err2 := base64.StdEncoding.DecodeString(buf)
	if nil != err2 {
		t.Fatal(err2)
	}

	t.Log(string(dBuf))

	secretKey := "09936CD67213466DB4EF67DC0E92C4B5"
	pwdIv := []byte(secretKey[:16])

	cbc, err01 := NewAESCryptoCBC([]byte(secretKey[:16]))
	if nil != err01 {
		t.Fatal(err01)
		return
	}

	plainBuf := cbc.DecryptWithIV(dBuf, pwdIv)
	t.Log(plainBuf)
	t.Log(string(plainBuf))

}
