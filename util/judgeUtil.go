package util

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type judgeUtil struct {
	exprSQLInject []string
}

var judgeUtilInstance judgeUtil

func GetInstanceByJudgeUtil() *judgeUtil {
	return &judgeUtilInstance
}

func init() {
	judgeUtilInstance.exprSQLInject = make([]string, 0)
	judgeUtilInstance.exprSQLInject = append(judgeUtilInstance.exprSQLInject, `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(where|select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`)
	judgeUtilInstance.exprSQLInject = append(judgeUtilInstance.exprSQLInject, `((\x27|')?\s+(o|O)(r|R)|^(o|O)(r|R))\s+?[[:alnum:]]+\s*(=|>|<|<>|>=|<=)\s*[[:alnum:]]+\s*`)
}

// Length 字符串的长度，而不是 len(s) 的字节长度
func (that *judgeUtil) Length(s string) int {
	r := []rune(s)
	return len(r)
}

// SubRune 字符串按字符长度切割
func (that *judgeUtil) SubRune(s string, begin, end int) string {
	r := []rune(s)
	if begin > len(r) {
		return ""
	}
	if end > len(r) {
		return ""
	}
	c := ""
	for i := begin; i < end; i++ {
		c = fmt.Sprintf("%s%c", c, r[i])
	}
	return c
}

// IsEmpty 字符串是否为空
func (that *judgeUtil) IsEmpty(text string) bool {
	if "" == text {
		return true
	}
	text = strings.TrimSpace(text)
	if "" == text {
		return true
	}
	return false
}

// IsEmptys 多字符串是否为空
// texts 中有一个为空则所有都为空
func (that *judgeUtil) IsEmptys(texts ...string) bool {
	for _, s := range texts {
		if that.IsEmpty(s) {
			return true
		}
	}
	return false
}

func (that *judgeUtil) IsNotEmpty(text string) bool {
	return !that.IsEmpty(text)
}

// EqualsIgnoreCase 不区分大小写的比较
func (that *judgeUtil) EqualsIgnoreCase(s1, s2 string) bool {
	if strings.ToLower(s1) == strings.ToLower(s2) {
		return true
	}
	return false
}

// IsErrors 有一个 err 不为 nil 则存在错误返回 true，没有错误则返回false
func (that *judgeUtil) IsErrors(errs ...error) bool {
	for _, e := range errs {
		if nil != e {
			return true
		}
	}
	return false
}

// IsEmail 如果 val 是邮箱则返回 true，否则 false
func (that *judgeUtil) IsEmail(val string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(val)
}

// IsPhoneInternationality 如果 val 是手机号（国际，必须有+区号）则返回 true，否则返回 false
func (that *judgeUtil) IsPhoneInternationality(val string) bool {
	pattern := `^\+((?:9[679]|8[035789]|6[789]|5[90]|42|3[578]|2[1-689])|9[0-58]|8[1246]|6[0-6]|5[1-8]|4[013-9]|3[0-469]|2[70]|7|1)(?:\W*\d){0,13}\d$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(val)
}

// IsPhoneChinese 手机号判断，大陆版本
func (that *judgeUtil) IsPhoneChinese(val string) bool {
	reg := regexp.MustCompile(`\d+`)
	if strings.HasPrefix(val, "+86") && 14 == len(val) {
		v := val[1:]
		return reg.MatchString(v)
	}
	if strings.HasPrefix(val, "1") && 11 == len(val) {
		return reg.MatchString(val)
	}
	return false
}

// IsLetterNumber 只有数字字母组成返回 true
func (that *judgeUtil) IsLetterNumber(val string) bool {
	ln := `^[\d-a-zA-Z]+$`
	reg := regexp.MustCompile(ln)
	n := `\d+`
	l := `[a-zA-Z]+`
	reg2 := regexp.MustCompile(n)
	reg3 := regexp.MustCompile(l)
	return reg.MatchString(val) && reg2.MatchString(val) && reg3.MatchString(val)
}

// IsLetterOrNumber 只有数字或字母组成返回 true
func (that *judgeUtil) IsLetterOrNumber(val string) bool {
	reg := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	return reg.MatchString(val)
}

// IsLetterNumber_ 只有数字或字母或下划线组成的返回 true【防不住 unicode】
func (that *judgeUtil) IsLetterNumber_(val string) bool {
	reg := regexp.MustCompile(`/^\w+$/`)
	return reg.MatchString(val)
}

// IsPhoneArea 手机区号必须是 + 开头，后面跟随数字
func (that *judgeUtil) IsPhoneArea(val string) bool {
	reg := regexp.MustCompile(`^\+\d+$`)
	return reg.MatchString(val)
}

// IsPhoneNum 手机号【国际通用】
func (that *judgeUtil) IsPhoneNum(val string) bool {
	reg := regexp.MustCompile(`^\d+$`)
	return reg.MatchString(val)
}

// IsDomainUrl 判断域名， 只能是 a.com，或者 b.com
func (that *judgeUtil) IsDomainUrl(val string) bool {
	reg := regexp.MustCompile("^[a-zA-Z0-9][-a-zA-Z0-9]{0,62}\\.([a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$")
	return reg.MatchString(val)
}

// IsContainsSpecialSymbols 是否包含特殊符号【常见的】
func (that *judgeUtil) IsContainsSpecialSymbols(val string) bool {
	regEx := "[\n`~!@#$%^&*()+=|{}':;',\\[\\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。， 、？✅]"
	reg := regexp.MustCompile(regEx)
	return reg.MatchString(val)
}

// IsIp 如果 val 是ip或者ip段则返回 true，否则 false   【192.168.1.1 或 192.168.1.1/24】
func (that *judgeUtil) IsIp(val string) bool {
	regEx := "^(((\\d{1,2})|(1\\d{1,2})|(2[0-4]\\d)|(25[0-5]))\\.){3}((\\d{1,2})|(1\\d{1,2})|(2[0-4]\\d)|(25[0-5]))(\\/((\\d{1,2})|(1\\d{1,2})|(2[0-4]\\d)|(25[0-5])))?$"
	reg := regexp.MustCompile(regEx)
	return reg.MatchString(val)
}

// IsSQLInject 是 sql 注入语句返回  True
func (that *judgeUtil) IsSQLInject(val string) bool {
	//****************************************
	//‘ or 1=1  and 1=1  'a'='a' 类型：
	//*****************************************
	//((\x27|')?\s+(o|O)(r|R)|^(o|O)(r|R))\s+?[[:alnum:]]+\s*(=|>|<|<>|>=|<=)\s*[[:alnum:]]+\s*
	//((\x27|')?\s+(o|O)(r|R)|^(o|O)(r|R))\s+?(\x27|')[[:alnum:]]+(\x27|')\s*(=|>|<|<>|>=|<=)\s*(\x27|')[[:alnum:]]+(\x27|')\s*
	//((\x27|')?\s+(a|A)(n|N)(d|D)|^(a|A)(n|N)(d|D))\s+?[[:alnum:]]+\s*(=|>|<|<>|>=|<=)\s*[[:alnum:]]+\s*
	//((\x27|')?\s+(a|A)(n|N)(d|D)|^(a|A)(n|N)(d|D))\s+?(\x27|')[[:alnum:]]+(\x27|')\s*(=|>|<|<>|>=|<=)\s*(\x27|')[[:alnum:]]+(\x27|')\s*
	//*****************************************
	//and 0<(select count(*) from admin) 类型：
	//*****************************************
	//((\x27|')?\s+(a|A)(n|N)(d|D)|^(a|A)(n|N)(d|D))\s+?\w+(=|>|<|<>|>=|<=)\s*?\(.*\)?
	//((\x27|')?\s+(a|A)(n|N)(d|D)|^(a|A)(n|N)(d|D))\s+?\(.*\)\s*?(=|>|<|<>|>=|<=)\s*?\w+
	//*****************************************
	//select ..from..
	//*****************************************
	//(select|Select|SELECT)\s+.*\s+(from|From|FROM)\s+
	//*****************************************
	//关键字
	//*****************************************
	//((;|\(| )(insert|INSERT|Insert)|^(insert|INSERT|Insert))\s+(into|Into|INTO)\s+
	//((;|\(| )(update|Update|UPDATE)|^(update|Update|UPDATE))\s+\w+\s+(set|Set|SET)\s+\w+\s?=.*

	for _, v := range that.exprSQLInject {
		re, err := regexp.Compile(v)
		if nil != err {
			fmt.Println(err)
			return false
		}
		if re.MatchString(strings.ToLower(val)) {
			return true
		}
	}

	return false
}

// JoinStringsInASCII 按照规则，参数名ASCII码从小到大排序后拼接
// data 待拼接的数据
// sep 连接符
// onlyValues 是否只包含参数值，true则不包含参数名，否则参数名和参数值均有
// includeEmpty 是否包含空值，true则包含空值，否则不包含，注意此参数不影响参数名的存在
// exceptKeys 被排除的参数名，不参与排序及拼接
func (that *judgeUtil) JoinStringsInASCII(data map[string]string, sep string, onlyValues, includeEmpty bool, exceptKeys ...string) string {
	var list []string
	var keyList []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if !includeEmpty && value == "" {
			continue
		}
		if onlyValues {
			keyList = append(keyList, k)
		} else {
			list = append(list, fmt.Sprintf("%s=%s", k, value))
		}
	}
	if onlyValues {
		sort.Strings(keyList)
		for _, v := range keyList {
			list = append(list, data[v])
		}
	} else {
		sort.Strings(list)
	}
	return strings.Join(list, sep)
}
