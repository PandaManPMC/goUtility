package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mssola/useragent"
	"io"
	"net/http"
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
func (*httpUtil) GetRequestIp(req *http.Request) string {
	ip := func(req *http.Request) string {
		// 优先使用 X-Forwarded-For
		fIp := req.Header.Get("X-Forwarded-For")
		if "" != fIp && !strings.Contains(fIp, "[") {
			// x.x.x.x,xx.xx.x.x,x.x.x.xx ...
			if strings.Contains(fIp, ",") {
				ips := strings.Split(fIp, ",")
				return ips[0]
			}
			return fIp
		}

		rIp := req.RemoteAddr
		// RemoteAddr=[::1] or 127.0.0.1
		if "" != rIp && !strings.Contains(rIp, "[") && !strings.HasPrefix(rIp, "127.") {
			return rIp
		}

		xIp := req.Header.Get("X-Real-IP")
		if "" != xIp && !strings.Contains(xIp, "[") {
			return xIp
		}

		remoteAddr := req.Header.Get("Remote_addr")
		if "" != remoteAddr && !strings.Contains(remoteAddr, "[") {
			return remoteAddr
		}

		return req.RemoteAddr
	}(req)
	if strings.Contains(ip, ":") {
		return strings.Split(ip, ":")[0]
	}
	return strings.Trim(ip, " ")
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
		return nil, errors.New(fmt.Sprintf("StatusCode=%d", resp.StatusCode))
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

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}
