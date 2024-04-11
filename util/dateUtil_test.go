package util

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestUnixToDate(t *testing.T) {
	t.Log(time.Now().Unix())
	unix := 1663524183
	t.Log(time.Unix(int64(unix), 0).Format(DateFormat_yyyy_MM_dd_HH_mm_ss))

	t.Log(time.Unix(time.Now().Unix(), 0).Format(DateFormat_yyyy_MM_dd_HH_mm_ss))

	t.Log(GetInstanceByDateUtil().IsToday(int64(unix)))
	t.Log(GetInstanceByDateUtil().IsToday(time.Now().Unix()))

	var s []string
	t.Log(s)
}

func TestDate1(t *testing.T) {
	day := time.Now().Format(DateFormat_yyyy_MM_dd)
	t.Log(day)
	t.Log(time.Now().Year())
	t.Log(time.Now().Month())
	t.Log(time.Now().Day())

	s := "2022-10-12"
	d, _ := time.Parse(DateFormat_yyyy_MM_dd, s)
	t.Log(d)
	t.Log(d.Unix())

}

func TestNowToYYYYMMddHHmmssMust(t *testing.T) {
	t.Log(GetInstanceByDateUtil().NowToYYYYMMddHHmmssMust())

	t.Log(GetInstanceByDateUtil().TodayUnix())

	d := time.Now().Add(-1)
	t.Log(GetInstanceByDateUtil().TimeToZeroUnix(d))
}

func TestTimeUnix(t *testing.T) {
	authDate := 1672363858
	tm := time.Unix(int64(authDate), 0)
	t.Log(tm.String())
}

func TestAddDay0OClockUnix(t *testing.T) {
	d := GetInstanceByDateUtil().AddDay0OClockUnix(-1)
	t.Log(d)
	s := GetInstanceByDateUtil().FormatUnixToYYYYMMddHHmmSSMust(int64(d))
	t.Log(s)

	d2 := GetInstanceByDateUtil().AddDay0OClockUnix(0)
	t.Log(d2)
	s2 := GetInstanceByDateUtil().FormatUnixToYYYYMMddHHmmSSMust(int64(d2))
	t.Log(s2)

	d3 := GetInstanceByDateUtil().AddDay0OClockUnix(-7)
	t.Log(d3)
	s3 := GetInstanceByDateUtil().FormatUnixToYYYYMMddHHmmSSMust(int64(d3))
	t.Log(s3)

	d4 := GetInstanceByDateUtil().AddDay0OClockUnix(-30)
	t.Log(d4)
	s4 := GetInstanceByDateUtil().FormatUnixToYYYYMMddHHmmSSMust(int64(d4))
	t.Log(s4)
}

func TestGetDapengDateBeginTime(t *testing.T) {
	t.Log(GetInstanceByDateUtil().GetDapengDateBeginTime())

	endTime := 1709246853
	dat := time.Unix(int64(endTime), 0)
	dat = dat.AddDate(0, 0, -1)
	datS := dat.Format(DateFormatyyyyMMdd)
	datI, _ := strconv.Atoi(datS)
	t.Log(datI)

	t.Log(time.Minute)

	startTimeStamp := uint64(1709225200)
	timeNow := uint64(time.Now().Unix())
	durationMinute := (timeNow - startTimeStamp) / 60

	t.Log(durationMinute)
	startMinuteDuration := time.Duration(durationMinute) * time.Minute

	// 每次最多5分钟
	endHourDuration := startMinuteDuration - 1*time.Minute
	t.Log(startMinuteDuration)
	t.Log(endHourDuration)

	t.Log(fmt.Sprintf("%v %v", time.Now().Add(-startMinuteDuration), time.Now().Add(-endHourDuration)))

	t.Log(fmt.Sprintf("%v", time.Minute*8))
	t.Log(fmt.Sprintf("%v", time.Minute*3))

}
