package util

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"text/template"
)

func TestUserAgent(t *testing.T) {
	// ua.Mobile() ua.Engine()

	s := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
	ua := GetInstanceByHttpUtil().UserAgent(s)
	fmt.Printf("%v\n", ua.Mobile())      // => true
	fmt.Printf("%v\n", ua.Bot())         // => false
	fmt.Printf("%v\n", ua.Mozilla())     // => "5.0"
	fmt.Printf("Model=%v\n", ua.Model()) // => "Nexus One"

	t.Log("------------- Platform")

	fmt.Printf("%v\n", ua.Platform()) // => "Linux"
	fmt.Printf("%v\n", ua.OS())       // => "Android 2.3.7"
	t.Log("------------- Engine")

	name, version := ua.Engine()
	fmt.Printf("%v\n", name)    // => "AppleWebKit"
	fmt.Printf("%v\n", version) // => "533.1"
	t.Log("------------- Browser")

	name, version = ua.Browser()
	fmt.Printf("%v\n", name)    // => "Android"
	fmt.Printf("%v\n", version) // => "4.0"

	t.Log("-------------")
	// Let's see an example with a bot.

	ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	fmt.Printf("%v\n", ua.Bot()) // => true

	name, version = ua.Browser()
	fmt.Printf("%v\n", name)    // => Googlebot
	fmt.Printf("%v\n", version) // => 2.1
	t.Log("-----------------")

	info := GetInstanceByHttpUtil().UserAgentInfo(s)
	t.Log(info)
}

// TestAppUserAgent
func TestAppUserAgent(t *testing.T) {
	// 所有参数中间要去除空格，每个参数都不要超过 30 个字符
	// 1.固定 MobileApp 空格 接下一个参数
	// 2.DeviceSystem   string 设备系统 填入 Android 或 IOS
	// 3.DeviceEngine   string 设备内核 填入 设备唯一标识 如设备 MAC 地址或设备序列号
	// 4.DevicePlatform string 设备平台 填入 设备型号 如 SM-S9010
	// "MobileApp DeviceSystem DevicePlatform DeviceEngine DevicePlatform"

}

func TestXss(t *testing.T) {
	s1 := "abc动感超人"
	t.Log(template.HTMLEscapeString(s1))
	t.Log(template.JSEscapeString(s1))

	s1 = "<div><img src='http://gogs.ytbxs1812.com:10300/repo-avatars/1'></div>"
	t.Log(template.HTMLEscapeString(s1))
	t.Log(template.JSEscapeString(s1))

	s1 = "<script>alert(1)</script>"
	t.Log(template.HTMLEscapeString(s1))
	t.Log(template.JSEscapeString(s1))
	t.Log(len(s1))
	t.Log(len(template.HTMLEscapeString(s1)))
	t.Log(len(template.JSEscapeString(s1)))

	s1 = "script_username"
	t.Log(template.HTMLEscapeString(s1))
	t.Log(template.JSEscapeString(s1))

	s1 = "alert(1);"
	t.Log(template.HTMLEscapeString(s1))
	t.Log(template.JSEscapeString(s1))
}

func TestDetectContentType(t *testing.T) {

	f, _ := os.Open("C:\\Users\\LaoGou\\Pictures\\74.png")
	b := make([]byte, 512)
	f.Read(b)
	s := http.DetectContentType(b)
	t.Log(s)
	fmt.Println(s)

}
