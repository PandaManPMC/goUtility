package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

type googleAuth struct {
}

var googleAuthInstance googleAuth

func GetInstanceByGoogleAuth() *googleAuth {
	return &googleAuthInstance
}

func (that *googleAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

func (that *googleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (that *googleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (that *googleAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (that *googleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (that *googleAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (that *googleAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := that.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := that.toUint32(hashParts)
	return number % 1000000
}

//	GetSecret 获取秘钥
//	密钥生成二维码的规则  otpauth://totp/用户的用户名?secret=秘钥&issuer=UUWIN
func (that *googleAuth) GetSecret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, that.un())
	return strings.ToUpper(that.base32encode(that.hmacSha1(buf.Bytes(), nil)))
}

func (that *googleAuth) GetCode(secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := that.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := that.oneTimePassword(secretKey, that.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

func (that *googleAuth) GetQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

func (that *googleAuth) GetQrcodeUrl(user, secret string) string {
	qrcode := that.GetQrcode(user, secret)
	return fmt.Sprintf("http://www.google.com/chart?chs=200x200&chld=M%%7C0&cht=qr&chl=%s", qrcode)
}

func (that *googleAuth) VerifyCode(secret, code string) (bool, error) {
	if "" == code {
		return false, nil
	}
	_code, err := that.GetCode(secret)
	if err != nil {
		return false, err
	}
	return _code == code, nil
}
