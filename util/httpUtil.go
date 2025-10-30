package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mssola/useragent"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type httpUtil struct {
}

var httpUtilInstance httpUtil

func GetInstanceByHttpUtil() *httpUtil {
	return &httpUtilInstance
}

// GetRequestIp 获取客户端 ip，绕过代理，会去除端口号
//func (*httpUtil) GetRequestIp(req *http.Request) string {
//	ip := func(req *http.Request) string {
//		// 优先使用 X-Forwarded-For
//		fIp := req.Header.Get("X-Forwarded-For")
//		if "" != fIp && !strings.Contains(fIp, "[") {
//			// x.x.x.x,xx.xx.x.x,x.x.x.xx ...
//			if strings.Contains(fIp, ",") {
//				ips := strings.Split(fIp, ",")
//				return ips[0]
//			}
//			return fIp
//		}
//
//		rIp := req.RemoteAddr
//		// RemoteAddr=[::1] or 127.0.0.1
//		if "" != rIp && !strings.Contains(rIp, "[") && !strings.HasPrefix(rIp, "127.") {
//			return rIp
//		}
//
//		xIp := req.Header.Get("X-Real-IP")
//		if "" != xIp && !strings.Contains(xIp, "[") {
//			return xIp
//		}
//
//		remoteAddr := req.Header.Get("Remote_addr")
//		if "" != remoteAddr && !strings.Contains(remoteAddr, "[") {
//			return remoteAddr
//		}
//
//		return req.RemoteAddr
//	}(req)
//	if strings.Contains(ip, ":") {
//		return strings.Split(ip, ":")[0]
//	}
//	return strings.Trim(ip, " ")
//}

// GetRequestIp 获取客户端 ip，绕过代理，会去除端口号
func (*httpUtil) GetRequestIp(req *http.Request) string {
	ip := func(req http.Header) string {
		// 优先使用 X-Forwarded-For
		fIp := req.Get("X-Forwarded-For")
		if "" != fIp {
			if strings.Contains(fIp, "[") {
				fIp = strings.ReplaceAll(fIp, "[", "")
				fIp = strings.ReplaceAll(fIp, "]", "")
			}

			// x.x.x.x,xx.xx.x.x,x.x.x.xx ...
			if strings.Contains(fIp, ",") {
				ips := strings.Split(fIp, ",")
				return strings.TrimSpace(ips[0])
			}
			return strings.TrimSpace(fIp)
		}

		rIp := req.Get("RemoteAddr")
		if strings.Contains(rIp, "[") {
			rIp = strings.ReplaceAll(rIp, "[", "")
			rIp = strings.ReplaceAll(rIp, "]", "")
		}

		// RemoteAddr=[::1] or 127.0.0.1
		if "" != rIp && !strings.HasPrefix(rIp, "127.") {
			return rIp
		}

		xIp := req.Get("X-Real-IP")
		if strings.Contains(xIp, "[") {
			xIp = strings.ReplaceAll(xIp, "[", "")
			xIp = strings.ReplaceAll(xIp, "]", "")
		}
		if "" != xIp {
			return xIp
		}

		remoteAddr := req.Get("Remote_addr")
		if strings.Contains(remoteAddr, "[") {
			remoteAddr = strings.ReplaceAll(remoteAddr, "[", "")
			remoteAddr = strings.ReplaceAll(remoteAddr, "]", "")
		}
		if "" != remoteAddr {
			return remoteAddr
		}

		return req.Get("RemoteAddr")
	}(req.Header)

	if "" == ip {
		ip = req.RemoteAddr
		if strings.HasPrefix(ip, "[::") {
			ip = "127.0.0.1"
		}
		if strings.Contains(ip, ":") {
			ip = strings.Split(ip, ":")[0]
		}
	}

	return strings.ReplaceAll(ip, " ", "")
}

// UserAgent 解析 userAgent
func (*httpUtil) UserAgent(userAgent string) useragent.UserAgent {
	ua := useragent.New(userAgent)
	return *ua
}

type UserAgentInfo struct {
	DeviceSystem   string // 设备系统
	DeviceEngine   string // 设备内核
	DevicePlatform string // 设备平台
	DeviceBrowser  string // 浏览器
	DeviceType     uint8  // thing设备类型:1@未知;2@Mobile;3@PC;4@Android;5@IOS
}

// UserAgentInfo App 解析参考如下
// 所有参数中间要去除空格，每个参数都不要超过 30 个字符
// 1.固定 MobileApp 空格 接下一个参数
// 2.DeviceSystem   string 设备系统 填入 Android 或 IOS
// 3.DeviceEngine   string 设备内核 填入 设备唯一标识 如设备 MAC 地址或设备序列号
// 4.DevicePlatform string 设备平台 填入 设备型号 如 SM-S9010
// "MobileApp DeviceSystem DevicePlatform DeviceEngine DevicePlatform"
func (*httpUtil) UserAgentInfo(userAgent string) UserAgentInfo {

	info := new(UserAgentInfo)
	if strings.HasPrefix(userAgent, "MobileApp") {
		// APP
		info.DeviceType = 2
		lst := strings.Split(userAgent, " ")
		if 1 < len(lst) {
			info.DeviceSystem = lst[1]
		}
		if 2 < len(lst) {
			info.DeviceEngine = lst[2]
		}
		if 3 < len(lst) {
			info.DevicePlatform = lst[3]
		}

		if "Android" == info.DeviceSystem {
			info.DeviceType = 4
		}
		if "IOS" == info.DeviceSystem {
			info.DeviceType = 5
		}

		return *info
	}
	// 浏览器

	ua := useragent.New(userAgent)
	if ua.Mobile() {
		info.DeviceType = 2
	} else {
		info.DeviceType = 3
	}

	info.DeviceSystem = ua.OS()
	name, version := ua.Engine()
	info.DeviceEngine = fmt.Sprintf("%s_%s", name, version)
	info.DevicePlatform = ua.Platform()

	name, version = ua.Browser()
	info.DeviceBrowser = fmt.Sprintf("%s_%s", name, version)
	return *info
}

func (*httpUtil) Post(url string, data []byte, header map[string]string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if nil != err {
		return nil, err
	}

	if nil == header || 0 == len(header) {
		req.Header.Set("Content-Type", "application/json")
	} else {
		if _, isOk := header["Content-Type"]; !isOk {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return body, errors.New(fmt.Sprintf("StatusCode=%d", resp.StatusCode))
	}

	return body, nil
}

func (*httpUtil) CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		h.ServeHTTP(w, r)
	})
}

func (*httpUtil) Get(u string, header map[string]string,
	timeOutSecond uint, transport *http.Transport) ([]byte, error) {
	if 0 == timeOutSecond {
		timeOutSecond = 30
	}

	client := &http.Client{
		Timeout: time.Duration(timeOutSecond) * time.Second,
	}

	if nil != transport {
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", u, nil)
	if nil != err {
		return nil, err
	}

	if nil != header {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("StatusCode=%d Status=%s", resp.StatusCode, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

func (*httpUtil) PostClient(u string, header map[string]string, inBody []byte,
	timeOutSecond uint, transport *http.Transport) ([]byte, error) {
	if 0 == timeOutSecond {
		timeOutSecond = 30
	}

	client := &http.Client{
		Timeout: time.Duration(timeOutSecond) * time.Second,
	}

	if nil != transport {
		client.Transport = transport
	}

	req, err := http.NewRequest("POST", u, bytes.NewReader(inBody))
	if nil != err {
		return nil, errors.New(fmt.Sprintf("url=%s err=%s", u, err.Error()))
	}

	if nil == header || 0 == len(header) {
		req.Header.Set("Content-Type", "application/json")
	} else {
		if _, isOk := header["Content-Type"]; !isOk {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("StatusCode=%d Status=%s", resp.StatusCode, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

const ERRORByRefMethodIsValidErr = "!refMethod.IsValid"
const ERRORByRequestParams = "RequestParams"

func (that *httpUtil) HandlerFun(instance any, methodName string, logError func(string, error)) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		refValue := reflect.ValueOf(instance)
		var refMethod reflect.Value
		refMethod = refValue.MethodByName(methodName)
		if !refMethod.IsValid() && reflect.Ptr == refValue.Kind() {
			refMethod = refValue.Elem().MethodByName(methodName)
		}

		if !refMethod.IsValid() {
			if nil != logError {
				logError(ERRORByRefMethodIsValidErr, errors.New(ERRORByRefMethodIsValidErr))
			}
			return
		}

		refMtdType := refMethod.Type()
		numIn := refMtdType.NumIn()
		methodParams := make([]reflect.Value, numIn)

		if 0 != numIn {
			if err := that.RequestParams(writer, request, &methodParams, refMtdType); nil != err {
				if nil != logError {
					logError(ERRORByRequestParams, err)
				}
				return
			}
		}
		// 响应
		result := refMethod.Call(methodParams)
		if nil != result && 0 < len(result) {
			rsu := result[0]
			switch rsu.Kind() {
			case reflect.Struct, reflect.Map, reflect.Slice:
				marshalData, _ := json.Marshal(rsu.Interface())
				writer.Write(marshalData)
			default:
				writer.Write([]byte(rsu.String()))
			}
		}

	}
}

// RequestParams 请求参数封装
func (that *httpUtil) RequestParams(writer http.ResponseWriter, request *http.Request, methodParams *[]reflect.Value, refMtdType reflect.Type) error {
	for i := 0; i < refMtdType.NumIn(); i++ {
		inType := refMtdType.In(i)
		switch inType.String() {
		case "*http.Request":
			(*methodParams)[i] = reflect.ValueOf(request)
		case "http.ResponseWriter":
			(*methodParams)[i] = reflect.ValueOf(writer)
		default:
			if err := that.RequestToData(request, methodParams, inType, i); nil != err {
				return err
			}
		}
	}
	return nil
}

// RequestToData request 中参数封装进结构体或 map，支持 Content-Type 【application/json || application/x-www-form-urlencoded】
func (that *httpUtil) RequestToData(request *http.Request, methodParams *[]reflect.Value, inType reflect.Type, index int) error {
	if that.ContentTypeIsJSON(request) {
		// 以 【application/json】
		defer request.Body.Close()
		buf, err := io.ReadAll(request.Body)
		if nil != err {
			return errors.New("request to data failure")
		}
		obj := reflect.New(inType)
		if err := json.Unmarshal(buf, obj.Interface()); nil != err {
			return errors.New("request to json Unmarshal data failure")
		}
		(*methodParams)[index] = obj.Elem()

		if reflect.Slice == inType.Kind() {
			return nil
		}
		mp := make(map[string]interface{})
		json.Unmarshal(buf, &mp)
		return that.requiredJSON(obj, mp)
	} else {
		// 其它-以 form 形式读取参数 【application/x-www-form-urlencoded】
		request.ParseForm()
		form := request.Form
		var err error
		if (*methodParams)[index], err = that.formToTypeValue(inType, form); nil != err {
			return err
		}
	}
	return nil
}

func (that *httpUtil) ContentTypeIsJSON(request *http.Request) bool {
	contentType := that.GetContentType(request)
	if strings.Contains(contentType, "application/json") {
		return true
	}
	return false
}

func (that *httpUtil) GetContentType(request *http.Request) string {
	contentType := request.Header.Get("Content-Type")
	if 0 == len(contentType) {
		return request.Header.Get("content-type")
	}
	return contentType
}

func (that *httpUtil) requiredJSON(bean reflect.Value, mp map[string]interface{}) error {
	beanElem := bean.Elem()
	num := beanElem.NumField()
	t := beanElem.Type()
	for i := 0; i < num; i++ {
		f := t.Field(i)
		rd := f.Tag.Get("required")
		json := f.Tag.Get("json")
		val, isExist := mp[json]

		//	字符串默认去除首位空格
		if "string" == f.Type.Name() && isExist {
			trimSpace := f.Tag.Get("trimSpace")
			if "false" != trimSpace {
				field := beanElem.Field(i)
				field.Set(reflect.ValueOf(strings.TrimSpace(field.String())))
				val = strings.TrimSpace(val.(string))
			}
		}

		if "true" == rd {
			//field := beanElem.Field(i)
			if !isExist {
				return errors.New(fmt.Sprintf("missing required parameters by 【%s】", json))
			}
			mi := f.Tag.Get("min")
			if "" != mi {
				// 最小值，作用于字符串长度
				if "string" == f.Type.Name() {
					s := val.(string)
					minLen, _ := strconv.Atoi(mi)
					if minLen > StringToCharacterLen(s) {
						return errors.New(fmt.Sprintf("Parameter 【%s】 minimum length of %d, your %d", json, minLen, StringToCharacterLen(s)))
					}
				}
			}
		}
		if !isExist {
			// 非必传，如果值不存在，直接回退
			continue
		}
		// 最大值，作用于字符串即字符最大长度
		ma := f.Tag.Get("max")
		if "" != ma {
			if "string" == f.Type.Name() {
				s := val.(string)
				maxLen, _ := strconv.Atoi(ma)
				if maxLen < StringToCharacterLen(s) {
					return errors.New(fmt.Sprintf("parameter 【%s】 maximum length %d, yours %d", json, maxLen, StringToCharacterLen(s)))
				}
			}
		}
	}
	return nil
}

// formToTypeValue
// map数据根据reflect.Type转为reflect.Value
// fiType reflect.Type	类型,map、struct支持，其它都为string
// form map[string][]string 数据源 如request.Form
// reflect.Value	值
func (that *httpUtil) formToTypeValue(fiType reflect.Type, form map[string][]string) (reflect.Value, error) {
	fiTypeKind := fiType.Kind()
	switch fiTypeKind {
	case reflect.Map:
		valMap := make(map[string]string, len(form))
		for key, values := range form {
			valMap[key] = StringArrayToString(values)
		}
		return reflect.ValueOf(valMap), nil
	case reflect.Struct:
		stVal := reflect.New(fiType)
		stType := stVal.Type()
		stElem := stType.Elem()
		numFiled := stElem.NumField()
		for i := 0; i < numFiled; i++ {
			tf := stElem.Field(i)
			tagJson := tf.Tag.Get("json")
			required := tf.Tag.Get("required")

			value, isExist := form[tagJson]
			if "true" == required {
				if !isExist {
					return stVal, errors.New(fmt.Sprintf("missing required parameters by 【%s】", tagJson))
				}
			} else {
				if !isExist {
					continue
				}
			}

			sv := StringArrayToString(value)
			if "" == sv && "string" != tf.Type.Name() {
				continue
			}

			if "string" == tf.Type.Name() {
				trimSpace := tf.Tag.Get("trimSpace")
				if "false" != trimSpace {
					sv = strings.TrimSpace(sv)
				}
			}

			// 字符串类型校验长度
			if "true" == required && "string" == tf.Type.Name() {
				mi := tf.Tag.Get("min")
				if "" != mi {
					minLen, _ := strconv.Atoi(mi)
					if minLen > StringToCharacterLen(sv) {
						return stVal, errors.New(fmt.Sprintf("Parameter 【%s】 minimum length of %d, your %d", tagJson, minLen, StringToCharacterLen(sv)))
					}
				}
			}

			ma := tf.Tag.Get("max")
			if "" != ma {
				if "string" == tf.Type.Name() {
					maxLen, _ := strconv.Atoi(ma)
					if maxLen < StringToCharacterLen(sv) {
						return stVal, errors.New(fmt.Sprintf("parameter 【%s】 maximum length %d, yours %d", tagJson, maxLen, StringToCharacterLen(sv)))
					}
				}
			}

			v := StringToType(tf.Type.Name(), sv)
			stVal.Elem().Field(i).Set(reflect.ValueOf(v))

			//for key, value := range form {
			//	if key == tagJson {
			//		if 0 == len(value) {
			//			if "true" == required {
			//				return stVal, errors.New(fmt.Sprintf("missing required parameters %s", tagJson))
			//			}
			//			break
			//		}
			//		sv := stringArrayToString(value)
			//		if "" == sv && tf.Type.Name() != "string" {
			//			break
			//		}
			//		v := stringToType(tf.Type.Name(), sv)
			//		stVal.Elem().Field(i).Set(reflect.ValueOf(v))
			//		break
			//	}
			//}
		}
		return stVal.Elem(), nil
	default:
		vaList := ""
		for _, values := range form {
			vaList = StringArrayToString(values)
		}
		return reflect.ValueOf(vaList), nil
	}
}

// StringToCharacterLen 获取字符串准确字符长度，而非字节长度
func StringToCharacterLen(s string) int {
	return len([]rune(s))
}

// StringArrayToString 字串数组转字串，以,拼接
// strArr []string	字串数组
// string	以【】间隔的值
func StringArrayToString(strArr []string) string {
	str := ""
	for inx, _ := range strArr {
		if 0 == inx {
			str = strArr[inx]
			continue
		}
		str = fmt.Sprintf("%s【】%s", str, strArr[inx])
	}
	return str
}

// StringToType 将string参数转为typeStr指定类型的值
// typeStr string	类型字串	支持int、float、bool、Time
// valueStr string	值
// interface{}	为 nil则失败
func StringToType(typeStr string, valueStr string) interface{} {
	var data interface{}
	var e error
	switch typeStr {
	case "int":
		if "" == valueStr {
			return 0
		}
		data, e = strconv.Atoi(valueStr)
	case "uint":
		if "" == valueStr {
			return uint(0)
		}
		val, e1 := strconv.Atoi(valueStr)
		if nil == e1 {
			data = uint(val)
		} else {
			e = e1
			data = val
		}
	case "int8":
		if "" == valueStr {
			return int8(0)
		}
		data, e = strconv.ParseInt(valueStr, 10, 8)
		if nil == e {
			data = int8(data.(int64))
		}
	case "uint8":
		if "" == valueStr {
			return uint8(0)
		}
		data, e = strconv.ParseUint(valueStr, 10, 8)
		if nil == e {
			data = uint8(data.(uint64))
		}
	case "int16":
		if "" == valueStr {
			return int16(0)
		}
		data, e = strconv.ParseInt(valueStr, 10, 16)
		if nil == e {
			data = int16(data.(int64))
		}
	case "uint16":
		if "" == valueStr {
			return uint16(0)
		}
		data, e = strconv.ParseUint(valueStr, 10, 16)
		if nil == e {
			data = uint16(data.(uint64))
		}
	case "int32":
		if "" == valueStr {
			return int32(0)
		}
		data, e = strconv.ParseInt(valueStr, 10, 32)
		if nil == e {
			data = int32(data.(int64))
		}
	case "uint32":
		if "" == valueStr {
			return uint32(0)
		}
		data, e = strconv.ParseUint(valueStr, 10, 32)
		if nil == e {
			data = uint32(data.(uint64))
		}
	case "int64":
		if "" == valueStr {
			return int64(0)
		}
		data, e = strconv.ParseInt(valueStr, 10, 64)
	case "uint64":
		if "" == valueStr {
			return uint64(0)
		}
		data, e = strconv.ParseUint(valueStr, 10, 64)
	case "bool":
		if "" == valueStr {
			return false
		}
		data, e = strconv.ParseBool(valueStr)
	case "float32":
		if "" == valueStr {
			return 0
		}
		data, e = strconv.ParseFloat(valueStr, 32)
		if nil == e {
			data = float32(data.(float64))
		}
	case "float64":
		if "" == valueStr {
			return 0
		}
		data, e = strconv.ParseFloat(valueStr, 64)
	case "string":
		if "" == valueStr {
			return ""
		}
		data = valueStr
	case "Time":
		if 10 == len(valueStr) {
			data, e = time.Parse("2006-01-02", valueStr)
		} else if 13 == len(valueStr) {
			data, e = time.Parse("2006-01-02 15", valueStr)
		} else if 16 == len(valueStr) {
			data, e = time.Parse("2006-01-02 15:04", valueStr)
		} else if 19 == len(valueStr) {
			data, e = time.Parse("2006-01-02 15:04:05", valueStr)
		} else {
			//data, e = time.Parse("2006-01-02'T'15:04:05.999 Z", valueStr)
		}
		if nil != e {
			e = nil
			data = time.Now()
		}
	}
	if nil != e {
		return nil
	}
	return data
}
