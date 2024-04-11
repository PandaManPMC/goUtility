package util

import (
	"github.com/google/uuid"
	"os/exec"
	"strings"
)

type uUIDUtil struct {
}

var uUIDUtilInstance uUIDUtil

func GetInstanceByUUIDUtil() *uUIDUtil {
	return &uUIDUtilInstance
}

// GetByTimeUUID 获取基于时间戳的 uuid
func (*uUIDUtil) GetByTimeUUID() string {
	u1, _ := uuid.NewUUID()
	return strings.ReplaceAll(u1.String(), "-", "")
}

// GetByRandomUUID 获取基于随机数的 uuid
func (*uUIDUtil) GetByRandomUUID() string {
	u4 := uuid.New()
	return strings.ReplaceAll(u4.String(), "-", "")
}

// UnixUUID 基于 unix、linux 操作系统命令生成 uuid
func (*uUIDUtil) UnixUUID() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
