package util

import (
	"fmt"
	"strconv"
	"strings"
)

type ipUtil struct {
}

var ipUtilInstance ipUtil

func GetInstanceByIpUtil() *ipUtil {
	return &ipUtilInstance
}

// IsIpRange ip 范围校验
func (*ipUtil) IsIpRange(ipMin, ipMax, ipTarget string) bool {
	minLst := strings.Split(ipMin, ".")
	maxLst := strings.Split(ipMax, ".")
	targetLst := strings.Split(ipTarget, ".")

	if len(minLst) != len(maxLst) || len(minLst) != len(targetLst) {
		return false
	}

	for i := 0; i < len(maxLst); i++ {
		maxIp, _ := strconv.ParseInt(maxLst[i], 10, 64)
		minIp, _ := strconv.ParseInt(minLst[i], 10, 64)
		tarIp, _ := strconv.ParseInt(targetLst[i], 10, 64)
		if maxIp < tarIp || minIp > tarIp {
			return false
		}
	}

	return true
}

// ChinaProvince 处理香港、台湾、澳门
func (*ipUtil) ChinaProvince(country string) string {
	province := country
	if "Hong Kong" == country || "Taiwan" == country || "Macao" == country {
		province = fmt.Sprintf("China %s", country)
	}
	return province
}
