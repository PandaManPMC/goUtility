package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mRand "math/rand"
	"strings"
	"testing"
	"time"
)

func TestRandNumber(t *testing.T) {
	random := GetInstanceByRandomUtil()
	for i := 0; i < 100; i++ {
		t.Log(random.RandNumber(100))
	}
	s := strings.ToUpper(random.RandCharacterString(320))
	t.Log(s)

	var sum uint64 = 36
	for i := 0; i < 16; i++ {
		sum *= 36
	}
	t.Log(sum)

	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = int(random.RandNumber(100))
	}
	t.Log(arr)
}

func TestUUID(t *testing.T) {
	t.Log(GetInstanceByNumberUtil())
	for i := 0; i < 30; i++ {
		ran := GetInstanceByRandomUtil()
		t.Log(ran.RandCharacterString(32))
	}

	for i := 0; i < 48; i++ {
		t.Log(GetInstanceByRandomUtil().RandCharacterString(12))
	}

}

type A struct {
	Name string
}

func Test1(t *testing.T) {
	a := A{Name: "ssa"}
	b := a

	t.Log(a)
	t.Log(b)

	b.Name = "二狗子b"

	t.Log(a)
	t.Log(b)

	t.Log("------------")
	c := 1
	var d interface{} = c

	s, isOk := d.(string)
	t.Log(isOk)
	t.Log(s)

}

func TestRandCharacterString(t *testing.T) {
	start := time.Now().Unix()
	mp := make(map[string]int)
	// 6字符，1千万次测试，重复870次,耗时20
	// 7字符，1千万次测试，重复15次,耗时29
	// 8字符，1千万次测试，重复0次,耗时28
	// 10字符，1亿次测试，重复0次,耗时371
	count := 100000000
	repeat := 0
	for i := 0; i < count; i++ {
		rand := GetInstanceByRandomUtil().RandCharacterString(10)
		_, isOk := mp[rand]
		if isOk {
			repeat++
			//t.Log(fmt.Sprintf("重复%s", rand))
		}
		mp[rand] = 1
	}

	t.Log(fmt.Sprintf("重复%d次,耗时%d", repeat, time.Now().Unix()-start))

}

func TestMathRand(t *testing.T) {
	mr := make(map[int64]int, 0)

	count := 100000
	begin := time.Now().Unix()
	for i := 0; i < count; i++ {
		mRand.Seed(time.Now().UnixNano())
		r := mRand.Int63n(int64(count))
		if _, isOk := mr[r]; isOk {
			mr[r] += 1
		} else {
			mr[r] = 1
		}
	}
	t.Log("math rand ", time.Now().Unix()-begin)
	fmt.Println(mr)
	mr = make(map[int64]int, 0)

	begin = time.Now().Unix()
	for i := 0; i < count; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(count)))
		r := result.Int64()
		if _, isOk := mr[r]; isOk {
			mr[r] += 1
		} else {
			mr[r] = 1
		}
	}
	t.Log("crypto rand ", time.Now().Unix()-begin)
	fmt.Println(mr)
}

func TestRand(t *testing.T) {

	mr := make(map[int64]int, 0)
	count := 1000000
	max := 10000
	for i := 0; i < count; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		r := result.Int64()
		if _, isOk := mr[r]; isOk {
			mr[r] += 1
		} else {
			mr[r] = 1
		}
	}

	cou := 0
	target := 130
	for _, v := range mr {
		if v > target {
			cou++
		}
	}
	t.Log(fmt.Sprintf("最大 10000，%d次重复超过%d次的数字数量%d", count, target, cou))

}
