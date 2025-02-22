package ethWal

import (
	"encoding/hex"
	"fmt"
	"goUtility/util"
	"golang.org/x/crypto/sha3"
	"testing"
)

func Test1(t *testing.T) {
	privateKey, err := util.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey("")
	if nil != err {
		t.Fatal(err)
	}

	// 2. 获取公钥
	publicKey := privateKey.PublicKey

	// 3. 将公钥序列化为字节数组 (只需要 X 和 Y 坐标，去掉前缀 0x04)
	publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)

	// 4. 对公钥进行 Keccak-256 哈希运算
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes)
	hashed := hash.Sum(nil)

	// 5. 获取哈希的后 20 字节
	address := hashed[len(hashed)-20:]

	// 6. 转换为十六进制字符串
	fmt.Printf("Ethereum Address: 0x%s\n", hex.EncodeToString(address))
}

func TestPrivateKeyToAddressETH(t *testing.T) {
	privateKey, err := util.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey("")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(PrivateKeyToAddressETH(privateKey))
}
