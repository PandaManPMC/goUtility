package util

import "testing"

func TestIsIpRange(t *testing.T) {

	t.Log(GetInstanceByIpUtil().IsIpRange("1.0.8.0", "1.0.15.255", "1.0.11.22"))
	t.Log(GetInstanceByIpUtil().IsIpRange("1.0.8.0", "1.0.15.255", "1.0.15.223"))

	t.Log(GetInstanceByIpUtil().IsIpRange("1.0.8.0", "1.0.15.255", "2.0.11.22"))
	t.Log(GetInstanceByIpUtil().IsIpRange("1.0.8.0", "1.0.15.255", "1.0.7.223"))
	t.Log(GetInstanceByIpUtil().IsIpRange("1.0.8.0", "1.0.15.255", "1.1.11.22"))
	t.Log(GetInstanceByIpUtil().IsIpRange("1.0.8.0", "1.0.15.255", "1.1.18.22"))

	var arr []string
	t.Log(arr)
}
