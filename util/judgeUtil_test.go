package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestIsErrors(t *testing.T) {
	e := errors.New("1111")
	e2 := e
	t.Log(GetInstanceByJudgeUtil().IsErrors(e))
	e = nil
	t.Log(GetInstanceByJudgeUtil().IsErrors(e))
	t.Log(GetInstanceByJudgeUtil().IsErrors(e, e2))

	s := "  aa  "
	s2 := strings.TrimSpace(s)
	t.Log(len(s))
	t.Log(len(s2))

}

func TestJson(t *testing.T) {
	str := `{"ETHUSDT": {"days_7": {"profit": "4.44%", "profit_annual": "231.02%", "max_drawdown": "-13.33%", "Volatility": "92.53%", "range": "1422.08 - 1711.53", "grid_pnl": "0.68% - 0.81%", "grid_num": "25"}, "days_30": {"profit": "17.77%", "profit_annual": "213.25%", "max_drawdown": "-53.31%", "Volatility": "87.57%", "range": "1422.08 - 1711.53", "grid_pnl": "0.68% - 0.81%", "grid_num": "25"}, "days_180": {"profit": "133.28%", "profit_annual": "266.56%", "max_drawdown": "-399.84%", "Volatility": "94.44%", "range": "1422.08 - 1711.53", "grid_pnl": "0.68% - 0.81%", "grid_num": "25"}}, "BTCUSDT": {"days_7": {"profit": "0.09%", "profit_annual": "4.76%", "max_drawdown": "-0.27%", "Volatility": "76.31%", "range": "19520.0 - 21886.77", "grid_pnl": "0.43% - 0.48%", "grid_num": "25"}, "days_30": {"profit": "0.37%", "profit_annual": "4.39%", "max_drawdown": "-1.10%", "Volatility": "77.96%", "range": "19520.0 - 21886.77", "grid_pnl": "0.43% - 0.48%", "grid_num": "25"}, "days_180": {"profit": "2.74%", "profit_annual": "5.49%", "max_drawdown": "-8.23%", "Volatility": "91.76%", "range": "19520.0 - 21886.77", "grid_pnl": "0.43% - 0.48%", "grid_num": "25"}}}`

	jsonObject := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &jsonObject)
	if nil != err {
		t.Fatal(err)
	}

	eu := jsonObject["ETHUSDT"]
	euMap, isOk := eu.(map[string]interface{})
	if !isOk {
		t.Fatal("类型断言失败")
	}

	t.Log(euMap)
	t.Log(euMap["days_7"])

	d7 := euMap["days_7"]
	d7Map, isOk := d7.(map[string]interface{})
	if !isOk {
		t.Fatal("d7 类型断言失败")
	}

	profit, isOk := d7Map["profit"].(string)
	if !isOk {
		t.Fatal("profit 类型断言失败")
	}
	t.Log(profit)
}

func TestIsPhoneInternationality(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsPhoneInternationality("15301991892"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneInternationality("+8615301991892"))
}

func TestIsPhoneChinese(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsPhoneChinese("15301991892"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneChinese("+8615301991892"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneChinese("153019918921"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneChinese("15301991892a"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneChinese("15r019918921"))
}

func TestIsEmail(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsEmail("AAAA@qq.com"))
	t.Log(GetInstanceByJudgeUtil().IsEmail("12222@qq.com"))
	t.Log(GetInstanceByJudgeUtil().IsEmail("12222.qq.com"))
	t.Log(GetInstanceByJudgeUtil().IsEmail("1222#@+qq.com"))
}

func TestIsUserName(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("哈哈"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("哈哈——"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("哈哈_"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("abc_123"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("abc"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("123"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("abc123"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("abcA123"))
	t.Log(GetInstanceByJudgeUtil().IsLetterNumber("abc@123"))
	s := "哈哈哈"
	t.Log(len(s))
	r := []rune(s)
	t.Log(len(r))

}

func TestIsPhoneArea(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+86"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+05"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+9996"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+995"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+0"))
	t.Log("-----------------")
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+9a"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("+00000d"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneArea("++12"))
}

func TestIsPhoneNum(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("1"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("153"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("177881"))
	t.Log("----------------------------")
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("1dsa"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("dsd1dsd"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("as1ds"))
	t.Log(GetInstanceByJudgeUtil().IsPhoneNum("ds1sas"))
}

func TestIsContainsSpecialSymbols(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsContainsSpecialSymbols("sd"))
	t.Log(GetInstanceByJudgeUtil().IsContainsSpecialSymbols("12d"))
	t.Log(GetInstanceByJudgeUtil().IsContainsSpecialSymbols("sd12"))
	t.Log("---------------------------------")
	t.Log(GetInstanceByJudgeUtil().IsContainsSpecialSymbols("s009c s"))
	t.Log(GetInstanceByJudgeUtil().IsContainsSpecialSymbols("s009c s!"))
	t.Log(GetInstanceByJudgeUtil().IsContainsSpecialSymbols("s009c s#$“"))

}

func TestSubRune(t *testing.T) {
	s := "留得青山在"
	t.Log(len(s))
	t.Log(s[1:3])

	t.Log(GetInstanceByJudgeUtil().SubRune(s, 1, 3))
	t.Log(GetInstanceByJudgeUtil().SubRune(s, 0, 5))

}

func TestChartSort(t *testing.T) {
	type telegramUser struct {
		Id        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		UserName  string `json:"username"`
		PhotoUrl  string `json:"photo_url"`
		AuthDate  int    `json:"auth_date"`
		Hash      string `json:"hash"`
	}

	s := `{"id":5684901207,"first_name":"Lao","last_name":"Gou","username":"laogou2211","photo_url":"https://t.me/i/userpic/320/lH8XYnSITZAAVgJRx_akSMsAh93mxcbrl5ZGn5AP_zGf2JSwptAvY01d75pY1pIb.jpg","auth_date":1672363858,"hash":"67b150fba14ca220e1dbeb70fe5cd5230806e3da34df97e81f8beae78fe8fc60"}`
	t.Log(s)

	user := new(telegramUser)
	json.Unmarshal([]byte(s), user)
	t.Log(user)

	mp := make(map[string]interface{})
	json.Unmarshal([]byte(s), &mp)

	t.Log(mp)
	hash := mp["hash"]

	delete(mp, "hash")
	//delete(mp, "username")

	t.Log(mp)

	keys := make([]string, 0)

	for k, _ := range mp {
		//t.Log(fmt.Sprintf("%s = %v", k, v))
		keys = append(keys, k)
	}

	sort.Strings(keys)
	t.Log(keys)

	// 签名
	userData := strings.Builder{}
	for inx, k := range keys {
		userRef := reflect.ValueOf(user)
		userRef = userRef.Elem()
		userTyp := reflect.TypeOf(user).Elem()

		for i := 0; i < userRef.NumField(); i++ {
			val := userRef.Field(i)
			ft := userTyp.Field(i)
			json := ft.Tag.Get("json")
			if k == json {
				kind := val.Kind()
				if kind == reflect.String {
					f := fmt.Sprintf("%s=%s", k, val.String())
					userData.WriteString(f)
				}
				if kind == reflect.Int {
					f := fmt.Sprintf("%s=%d", k, val.Int())
					userData.WriteString(f)
				}
			}
		}

		if inx != len(keys)-1 {
			userData.WriteString("\n")
		}
	}
	ud := userData.String()
	t.Log(ud)

	token := "5898534009:AAFV-aClPcyRAxFZi33F6xfVoFuSPldZFXY"
	tokenMD := GetInstanceByMessageDigest().Sha256Buf(token)
	hmac1 := GetInstanceByMessageDigest().HmacSha256(ud, tokenMD)
	t.Log(hmac1)
	t.Log(hash)
	t.Log(hmac1 == hash)

	dc := `auth_date=1672363858\nfirst_name=Lao\nid=5684901207\nlast_name=Gou\nphoto_url=https://t.me/i/userpic/320/lH8XYnSITZAAVgJRx_akSMsAh93mxcbrl5ZGn5AP_zGf2JSwptAvY01d75pY1pIb.jpg\nusername=laogou2211`
	t.Log(dc)
	hmac2 := GetInstanceByMessageDigest().HmacSha256(dc, tokenMD)
	t.Log(hmac2)

	//hd, _ := hex.DecodeString(hmac1)
	//hs, _ := hex.DecodeString(user.Hash)
	//t.Log(hmac.Equal(hd, hs))

	t.Log(ud == dc)

}

func TestIsIp(t *testing.T) {
	isIp := GetInstanceByJudgeUtil().IsIp("145.145.145.145/202")
	t.Log(isIp)
}

func TestVer(t *testing.T) {
	pattern := `^v\d+$`
	reg := regexp.MustCompile(pattern)

	t.Log(reg.MatchString("v1"))
	t.Log(reg.MatchString("v121"))
	t.Log(reg.MatchString("v111"))
	t.Log(reg.MatchString("v0"))
	t.Log(reg.MatchString("v9"))

	t.Log("--------------------------")

	t.Log(reg.MatchString("v1ds"))
	t.Log(reg.MatchString("vds"))
	t.Log(reg.MatchString("v123ds"))
	t.Log(reg.MatchString("v"))
	t.Log(reg.MatchString("v1.2"))
	t.Log(reg.MatchString("v1.2f"))
	t.Log(reg.MatchString("1"))
	t.Log(reg.MatchString("d1"))
	t.Log(reg.MatchString("v1v1v2"))
	t.Log(reg.MatchString("vv9"))

}

func TestIsSQLInject(t *testing.T) {
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("select * from a"))
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("select 哈哈"))
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("select哈哈"))
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("selec哈哈"))
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("devcat"))
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("where 1==1"))
	t.Log(GetInstanceByJudgeUtil().IsSQLInject("WHERE 1==1"))
}

func TestGetMaxLenStr(t *testing.T) {
	s := "12345"
	t.Log(GetInstanceByJudgeUtil().GetMaxLenStr(s, 5))
	t.Log(GetInstanceByJudgeUtil().GetMaxLenStr(s, 4))
	t.Log(GetInstanceByJudgeUtil().GetMaxLenStr(s, 1))
}
