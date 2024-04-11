package util

import "testing"

func TestGetSecret(t *testing.T) {
	secret := GetInstanceByGoogleAuth().GetSecret()
	t.Log(secret)
}
