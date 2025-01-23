package util

import "encoding/base64"

func DecodeBase64(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
