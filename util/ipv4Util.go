package util

type ipv4Util struct {
}

var ipv4UtilInstance ipv4Util

func GetInstanceByIpv4Util() *ipv4Util {
	return &ipv4UtilInstance
}

// IsLanIp ip 是否是局域网 ip
func (*ipv4Util) IsLanIp(ip string) bool {
	if 3 > len(ip) {
		return true
	}
	ip3 := ip[0:3]
	ip4 := ip[0:4]
	if "10." == ip3 || "172." == ip4 || "192." == ip4 {
		return true
	}
	return false
}
