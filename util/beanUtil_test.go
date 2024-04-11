package util

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestModelCopyToPointer(t *testing.T) {
}

type TestBean struct {
	Name string
}

func TestBeanStringTrimSpace(t *testing.T) {
	tb := TestBean{Name: "  ss "}
	GetInstanceByBeanUtil().BeanStringTrimSpace(&tb)
	t.Log(tb)
	tb2 := TestBean{Name: "  s  s "}
	GetInstanceByBeanUtil().BeanStringTrimSpace(&tb2)
	t.Log(tb2)
	a := -4
	t.Log(math.Abs(float64(a)))
}

type TestMouse struct {
	Name string
	Age  int
}

type TestCat struct {
	Name    string `json:"name" required:"true" min:"2" max:"5"`
	Age     int    `json:"age" required:"true"`
	Balance *float64
	Mouse   TestMouse
}

func TestRequiredParams(t *testing.T) {
	mp := make(map[string]interface{})
	//mp["name"] = "小1"
	//mp["age"] = 19
	//
	//c := new(TestCat)
	//s, b := GetInstanceByBeanUtil().RequiredParams(c, mp)
	//t.Log(b)
	//t.Log(s)
	//t.Log(c)

	t.Log("----------------------------")
	s := `{"account":"Hcvnrqik8cJ4PDH","action":0,"amount":"10","betArea":2,"betAreaEx":0,"betOrderId":69,"currency":"TRX/CNY","gameInstanceId":99522,"gameType":5,"tableId":2,"timestamp":1670984586000,"tradeId":"2_1_69"}`
	err := json.Unmarshal([]byte(s), &mp)
	if nil != err {
		t.Fatal(err)
	}

	sw := new(swBetIn)
	err = json.Unmarshal([]byte(s), sw)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(sw)

	m, b := GetInstanceByBeanUtil().RequiredParams(sw, mp)
	t.Log(b)
	t.Log(m)

}

type swBetIn struct {
	Account        string `json:"account" required:"true" min:"6" max:"16"`  // 账号
	TradeId        string `json:"tradeId" required:"true"`                   // 钱罐子变化的订单 id，唯一，同一笔注单下注和结算的tradeId 也不一样
	Currency       string `json:"currency" required:"true" min:"1" max:"10"` // 币种交易对(币种 i/币种 j)； 形如`BTC/CNY`
	Amount         string `json:"amount" required:"true" min:"1" max:"79"`   // 下注时的实际扣款数量, 大于等于下注筹码金额;(amount >
	GameType       int    `json:"gameType" required:"true"`                  // 游戏类型;见 gameType 枚举;
	TableId        int    `json:"tableId" required:"true"`                   // 桌台
	GameInstanceId int64  `json:"gameInstanceId" required:"true"`            // 游戏局号
	BetOrderId     int64  `json:"betOrderId" required:"true"`                // betOrderId: long ; 注单号，同一笔注单号 betOrderId 相同， 但可能会对应两个 tradeId (加点 tradeId 和扣点 tradeId)
	Timestamp      int    `json:"timestamp" required:"true"`                 // 时间戳， 单位:毫秒
	Action         int64  `json:"action"`                                    // int;默认填 0;
	BetArea        int64  `json:"betArea"`                                   // int; 投注区域(详见 下注区参数说明)
	BetAreaEx      int64  `json:"betAreaEx"`                                 // int; 投注区域补充字段
}

func TestBeanInstanceToListAddr(t *testing.T) {

	var c1 TestCat
	b := 11.22222223
	c1.Balance = &b

	lst := GetInstanceByBeanUtil().BeanInstanceToListAddr(&c1)
	//lst := make([]any, 3)
	//GetInstanceByBeanUtil().BeanInstanceToListAddr(&c1, lst)
	t.Log(&c1.Name)
	t.Log(&c1.Mouse.Name)
	t.Log(lst)
	//t.Log(c1)
	var name any
	name = "小黑"

	switch s := name.(type) {
	case string:
		switch d := lst[0].(type) {
		case *string:
			t.Log("switch -----------------------------------")
			t.Log(s)
			*d = "小黑"
		}
	}

	var bal1 float64 = 11.22
	var bal any = &bal1
	switch s := bal.(type) {
	case *float64:
		switch d := lst[2].(type) {
		case *float64:
			t.Log("switch 2 -----------------------------------")
			t.Log(s)
			*d = 12.223
		}
	}

	switch s := name.(type) {
	case string:
		switch d := lst[3].(type) {
		case *string:
			t.Log("switch333 -----------------------------------")
			t.Log(s)
			*d = "杰瑞"
		}
	}

	fmt.Println(c1)
	//t.Log(&lst[0])
	//t.Log(lst)
	//t.Log(c1)
	t.Log(*c1.Balance)

}
