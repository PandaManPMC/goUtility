package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestMD5Double(t *testing.T) {
	pwd := GetInstanceByMessageDigest().MD5Double("666666")
	t.Log(pwd)
	t.Log(len(pwd[17:]))

	t.Log(GetInstanceByMessageDigest().MD5Double("Coin@123456"))

	t.Log(GetInstanceByMessageDigest().MD5Double("123456"))
	t.Log(GetInstanceByMessageDigest().MD5("123456"))
	t.Log(GetInstanceByMessageDigest().MD5Double("q@123456"))

}

func TestSha256(t *testing.T) {
	message := "hello world!"
	s := GetInstanceByMessageDigest().Sha256(message)
	t.Log(s)

	t.Log(GetInstanceByMessageDigest().Sha256("75305c956e7d8e3280b6519ed1e33f7aceb335e8ac94b5afd6c6894a5084b867"))

	t.Log(GetInstanceByMessageDigest().Sha256("333333"))

	t.Log(GetInstanceByMessageDigest().Sha256("qaz123$fA"))

}

func aTestMap(mp map[string]string) {
	mp["abc"] = "123"
}

func TestMap(t *testing.T) {
	mp := make(map[string]string)
	aTestMap(mp)
	t.Log(mp)
	t.Log(GetInstanceByMessageDigest().Sha256("Qwer1234"))
}

func TestMD5Sig(t *testing.T) {
	now := time.Now().Unix()
	t.Log(now)

	r := make(map[string]any)
	r["id"] = 2
	r["timestamp"] = now

	md5Key := "e10adc3949ba59abbe56e057f20f883e"
	d, _ := json.Marshal(r)
	ds := fmt.Sprintf("%s%s", string(d), md5Key)
	t.Log(ds)
	t.Log(GetInstanceByMessageDigest().MD5(ds))

	t.Log(strings.ToUpper(GetInstanceByMessageDigest().MD5("aaaa1111")))

}
