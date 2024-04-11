package util

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"testing"
)

func TestCryptoECB(t *testing.T) {
	src := "0.56"
	key := "0123456789abcdef"

	encrypted, _ := AesEncrypt(src, []byte(key))
	_, err := AesDecrypt(encrypted, []byte(key))
	if err != nil {
		return
	}
	_, err = Base64URLDecode("39W7dWTd_SBOCM8UbnG6qA")
	if err != nil {
		return
	}
}

func TestPkcs7(t *testing.T) {
	src := "playerName=test&loginTime=2016-10-12T10:03:13Z&playerOdds=1960&channelId=1"
	key := "4B600JVJDB82B4H6ZHZ0J62T2B4004NL"

	encode, err := AesEcbPkcs7Encode(src, []byte(key))
	if err != nil {
		t.Logf(fmt.Sprintf("%+v", err))
		return
	}

	base64String := base64.StdEncoding.EncodeToString(encode)
	urlEncode := url.QueryEscape(base64String)
	fmt.Fprintf(os.Stdout, "encode: %s \r\n", urlEncode)

	queryUn, _ := url.QueryUnescape(urlEncode)
	decodeString, err := base64.StdEncoding.DecodeString(queryUn)
	if err != nil {
		return
	}

	decode, err := AesEcbPkcs7Decode(decodeString, []byte(key))
	if err != nil {
		return
	}

	t.Logf("decode: %v", string(decode))

}
