package util

import "testing"

func TestUnixUUID(t *testing.T) {
	uid, err := GetInstanceByUUIDUtil().UnixUUID()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(uid)
}
