package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type beanUtil struct {
}

var beanUtilInstance beanUtil

func GetInstanceByBeanUtil() *beanUtil {
	return &beanUtilInstance
}

// ModelCopyToPointer	拷贝单个实例
// resource interface{} 拷贝源
// toPointer interface{}	拷贝至，它应该是个指针
func (*beanUtil) ModelCopyToPointer(resource interface{}, toPointer interface{}) {
	rToPointer := reflect.ValueOf(toPointer)
	if reflect.Ptr != rToPointer.Kind() {
		return
	}
	rType := reflect.TypeOf(resource)
	rVal := reflect.ValueOf(resource)
	if reflect.Ptr == rVal.Kind() {
		rVal = rVal.Elem()
		rType = rType.Elem()
	}
	to := rToPointer.Elem()
	num := rType.NumField()
	for i := 0; i < num; i++ {
		rtf := rType.Field(i)
		jsonField := to.FieldByName(rtf.Name)
		if !jsonField.IsValid() {
			continue
		}
		if rtf.Type.String() != jsonField.Type().String() {
			continue
		}
		rff := rVal.Field(i)
		jsonField.Set(rff)
	}
}

// BeanStringTrimSpace 去除结构体实例所有字符串类型参数的首尾空格，只接受实例指针，改变直接作用于实例。
func (*beanUtil) BeanStringTrimSpace(pointer interface{}) {
	bv := reflect.ValueOf(pointer).Elem()
	num := bv.NumField()
	for i := 0; i < num; i++ {
		field := bv.Field(i)
		ft := field.Type()
		if "string" == ft.Name() {
			v := strings.TrimSpace(field.String())
			field.Set(reflect.ValueOf(v))
		}
	}
}

// BeanInstanceToListAddr 结构体实例所有值指针装入切片
// 支持指针、嵌套结构体。
func (that *beanUtil) BeanInstanceToListAddr(toPointer interface{}) []any {
	refInstance := reflect.ValueOf(toPointer)
	kind := refInstance.Kind()
	if reflect.Ptr != kind {
		return nil
	}
	elem := refInstance.Elem()
	lst := make([]interface{}, 0)

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fKind := field.Kind()
		fName := field.Type().Name()
		if "BaseModel" == fName {
			//	忽略 BaseModel
			continue
		}

		if reflect.Ptr == fKind {
			//	指针，如果为 nil 创建实例赋值
			if field.IsZero() {
				v := reflect.New(field.Type().Elem())
				lst = append(lst, v.Interface())
				field.Set(v)
				continue
			}
			lst = append(lst, field.Interface())
			continue
		}
		if reflect.Struct == fKind {
			// 持嵌套结构体，创建实例并读取指针
			lst2 := that.BeanInstanceToListAddr(field.Addr().Interface())
			lst = append(lst, lst2...)
			continue
		}
		lst = append(lst, field.Addr().Interface())
	}
	return lst
}

// RequiredParams 检查结构体的必传参数等
// 支持 required ，字符串类型支持 min、max 字符长度，会去除首位空格。
func (that *beanUtil) RequiredParams(dataPointer interface{}, params map[string]interface{}) (string, bool) {
	dr := reflect.ValueOf(dataPointer)

	dt := reflect.TypeOf(dataPointer)
	return that.RequiredParamsReflect(dr, dt, params)
}

// RequiredParamsReflect 检查结构体的必传参数等
// 支持 required ，字符串类型支持 min、max 字符长度，会去除首位空格。
func (that *beanUtil) RequiredParamsReflect(data reflect.Value, typeof reflect.Type, params map[string]interface{}) (string, bool) {
	dr := data
	drE := reflect.Indirect(dr)

	dt := typeof
	dtE := typeof

	if !data.IsValid() {
		dtE = dt.Elem()
	}

	for i := 0; i < dtE.NumField(); i++ {
		field := drE.Field(i)
		typ := dtE.Field(i)

		json := typ.Tag.Get("json")
		if "" == json {
			continue
		}
		val, isOk := params[json]
		//if isOk {
		//	field.Set(reflect.ValueOf(val))
		//}

		required := typ.Tag.Get("required")
		if "" == required {
			continue
		}

		// 必传参数，字符串参数核验长度，其它类型只看是否有值
		if !isOk {
			return fmt.Sprintf("%s is a required parameter", json), false
		}

		// 核实字符串长度
		if "string" != field.Type().String() {
			continue
		}

		max := typ.Tag.Get("max")
		if "" != max {
			// 最长
			maxLen, _ := strconv.Atoi(max)
			s := strings.TrimSpace(val.(string))
			r := []rune(s)
			if maxLen < len(r) {
				return fmt.Sprintf("%s up to %s, your length is %d", json, max, len(r)), false
			}
		}

		min := typ.Tag.Get("min")
		if "" != min {
			// 最短
			minLen, _ := strconv.Atoi(min)
			s := strings.TrimSpace(val.(string))
			r := []rune(s)
			if minLen > len(r) {
				return fmt.Sprintf("%s minimum %s, your length is %d", json, max, len(r)), false
			}
		}
	}
	return "", true
}

// Required 检查结构体的必传参数等
// 支持 required ，字符串类型支持 min、max 字符长度，会去除首位空格。
func (that *beanUtil) Required(data any) (string, bool) {
	dr := reflect.ValueOf(data)
	drE := dr.Elem()

	dt := reflect.TypeOf(data)
	dtE := dt.Elem()

	for i := 0; i < dtE.NumField(); i++ {
		field := drE.Field(i)
		typ := dtE.Field(i)

		json := typ.Tag.Get("json")
		if "" == json {
			continue
		}

		required := typ.Tag.Get("required")
		if "" == required {
			continue
		}

		// 核实字符串长度
		if "string" != field.Type().String() {
			continue
		}

		val := field.String()
		if "" == val {
			return fmt.Sprintf("%s is a required parameter", json), false
		}

		max := typ.Tag.Get("max")
		if "" != max {
			// 最长
			maxLen, _ := strconv.Atoi(max)
			s := strings.TrimSpace(val)
			r := []rune(s)
			if maxLen < len(r) {
				return fmt.Sprintf("%s up to %s, your length is %d", json, max, len(r)), false
			}
		}

		min := typ.Tag.Get("min")
		if "" != min {
			// 最短
			minLen, _ := strconv.Atoi(min)
			s := strings.TrimSpace(val)
			r := []rune(s)
			if minLen > len(r) {
				return fmt.Sprintf("%s minimum %s, your length is %d", json, max, len(r)), false
			}
		}
	}
	return "", true
}

// JSONCopy JSON 对象拷贝
// resource 源数据
// toPointer 拷贝目标
func (that *beanUtil) JSONCopy(resource, toPointer any) error {
	buf, e := json.Marshal(resource)
	if nil != e {
		return e
	}
	return json.Unmarshal(buf, toPointer)
}
