package util

import "testing"

func TestGetSecret(t *testing.T) {
	secret := GetInstanceByGoogleAuth().GetSecret()
	t.Log(secret)

	c, e := GetInstanceByGoogleAuth().GetCode("WID4UJUANTPCLZXIA7JC7TOUKIJ2OSGY")
	if nil != e {
		t.Fatal(e)
	}

	t.Log(c)
}
