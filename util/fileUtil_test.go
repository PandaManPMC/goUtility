package util

import (
	"fmt"
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	t.Log(GetInstanceByFileUtil().IsExists("/home/ubuntu/webapps/files/a0/a0/4hylki60lmaynh7kmo21gsro6z1o6pxn.png"))
	t.Log(GetInstanceByFileUtil().IsExists("/home/ubuntu/webapps/files/a0/a0/4hylki60lmaynh7kmo21gsro6z1o6pxn"))
	t.Log(GetInstanceByFileUtil().IsExists("D:\\tempWork"))

	s := "/home/ubuntu/webapps/files"
	s2 := "a0/a0/4hylki60lmaynh7kmo21gsro6z1o6pxn.png"
	s3 := fmt.Sprintf("%s%c%s", s, os.PathSeparator, s2)
	t.Log(s3)
	t.Log(GetInstanceByFileUtil().IsExists(s3))

	var lst []string
	t.Log(lst)

	lst = append(lst, "1111")
	t.Log(lst)

}
