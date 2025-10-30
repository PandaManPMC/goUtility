package util

import (
	"fmt"
	"strconv"
	"time"
)

type dateUtil struct {
}

const (
	// 北京时间
	TimeZoneBeijing = "Asia/Shanghai"

	DateFormatyyyyMMdd             = "20060102"
	DateFormatyyyyMMddHH           = "2006010215"
	DateFormat_yyyy_MM_dd          = "2006-01-02"
	DateFormat_yyyy_MM_dd_HH_mm_ss = "2006-01-02 15:04:05"
	DateFormat_yyyy_MM_dd_HH       = "2006-01-02 15"
	DateFormat_HH                  = "15"
	DateFormat_Minute              = "04"
	DateFormatyyyyMM               = "200601"
	DateFormatyyyyMMddHHmmss       = "20060102150405"   // 年月日时分秒 无分隔符
	DateFormtyyyyMMddHHmm          = "2006-01-02 15:04" // 年月日时：分
	DaySecond                      = 86400              // 一天秒数
)

var dateUtilInstance dateUtil

func GetInstanceByDateUtil() *dateUtil {
	return &dateUtilInstance
}

// IsToday 指定时间戳 unix 是今天 true 否则 false
func (*dateUtil) IsToday(unix int64) bool {
	d := time.Unix(unix, 0).Format(DateFormat_yyyy_MM_dd)
	n := time.Unix(time.Now().Unix(), 0).Format(DateFormat_yyyy_MM_dd)
	if d == n {
		return true
	}
	return false
}

// TodayUnix 今天 0 点 0 时 时间戳（本地时区 time.Local）
func (*dateUtil) TodayUnix() uint64 {
	now := time.Now()
	// 截断到天级别，获取当天 0 点
	today := now.Truncate(24 * time.Hour)
	timestamp := today.Unix()
	return uint64(timestamp)
}

// AddDay0OClockUnix 增加 days 的天数的0点 unix 时间
func (*dateUtil) AddDay0OClockUnix(days int) uint64 {
	d := time.Now().AddDate(0, 0, days)
	s := d.Format(DateFormat_yyyy_MM_dd)
	n, _ := time.ParseInLocation(DateFormat_yyyy_MM_dd, s, time.Local)
	return uint64(n.Unix())
}

func (*dateUtil) TodayUnixInt64() int64 {
	s := time.Now().Format(DateFormat_yyyy_MM_dd)
	n, _ := time.ParseInLocation(DateFormat_yyyy_MM_dd, s, time.Local)
	return n.Unix()
}

// TimeToZeroUnix 将 time 转为 00:00:00 的 unix 时间戳（本地时区 time.Local）
func (*dateUtil) TimeToZeroUnix(date time.Time) uint64 {
	s := date.Format(DateFormat_yyyy_MM_dd)
	n, _ := time.ParseInLocation(DateFormat_yyyy_MM_dd, s, time.Local)
	return uint64(n.Unix())
}

// ParseToYYYYMMddMust 将字符串 2022-10-12 格式化成 time.Time （使用 Parse 默认 UTC 时区）
func (*dateUtil) ParseToYYYYMMddMust(date string) time.Time {
	t, _ := time.Parse(DateFormat_yyyy_MM_dd, date)
	return t
}

// ParseToYYYYMMddHHmmSsMust 将字符串 2022-10-12 12:11:33 格式化成 time.Time （使用 Parse 默认 UTC 时区）
func (*dateUtil) ParseToYYYYMMddHHmmSsMust(date string) time.Time {
	t, _ := time.Parse(DateFormat_yyyy_MM_dd_HH_mm_ss, date)
	return t
}

// FormatUnixToYYYYMMddMust 将时间戳格式化成字符串 2022-10-12
func (*dateUtil) FormatUnixToYYYYMMddMust(t int64) string {
	y := time.Unix(t, 0)
	return y.Format(DateFormat_yyyy_MM_dd)
}

// FormatUnixToYYYYMMddHHMust 将时间戳格式化成字符串 2022-10-12 08
func (*dateUtil) FormatUnixToYYYYMMddHHMust(t int64) string {
	y := time.Unix(t, 0)
	return y.Format(DateFormat_yyyy_MM_dd_HH)
}

// FormatUnixToYYYYMMddIntegerMust 将时间戳格式化成整数 20221012
func (*dateUtil) FormatUnixToYYYYMMddIntegerMust(t int64) int {
	y := time.Unix(t, 0)
	s := y.Format(DateFormatyyyyMMdd)
	i, _ := strconv.Atoi(s)
	return i
}

// FormatUnixToYYYYMMddHHIntegerMust 将时间戳格式化成整数 2022101220
func (*dateUtil) FormatUnixToYYYYMMddHHIntegerMust(t int64) int {
	y := time.Unix(t, 0)
	s := y.Format(DateFormatyyyyMMddHH)
	i, _ := strconv.Atoi(s)
	return i
}

// FormatUnixToHHIntegerMust 将时间戳格式化成整数 12 【点】
func (*dateUtil) FormatUnixToHHIntegerMust(t int64) int {
	y := time.Unix(t, 0)
	s := y.Format(DateFormat_HH)
	i, _ := strconv.Atoi(s)
	return i
}

// FormatUnixToMinuteIntegerMust 将时间戳格式化成整数 5 【分】
func (*dateUtil) FormatUnixToMinuteIntegerMust(t int64) int {
	y := time.Unix(t, 0)
	s := y.Format(DateFormat_Minute)
	i, _ := strconv.Atoi(s)
	return i
}

// FormatUnixToYYYYMMddHHmmSSMust 将时间戳格式化成字符串 2022-10-12 11:22:11
func (*dateUtil) FormatUnixToYYYYMMddHHmmSSMust(t int64) string {
	y := time.Unix(t, 0)
	return y.Format(DateFormat_yyyy_MM_dd_HH_mm_ss)
}

// ParseToYYYYMMddLocalMust 将字符串 2022-10-12 格式化成 time.Time 使用本地时区
func (*dateUtil) ParseToYYYYMMddLocalMust(date string) time.Time {
	n, _ := time.ParseInLocation(DateFormat_yyyy_MM_dd, date, time.Local)
	return n
}

// ParseToYYYYMMddHHmmssLocalMust 将字符串 2022-10-12 12:09:00 格式化成 time.Time 使用本地时区
func (*dateUtil) ParseToYYYYMMddHHmmssLocalMust(date string) time.Time {
	n, _ := time.ParseInLocation(DateFormat_yyyy_MM_dd_HH_mm_ss, date, time.Local)
	return n
}

// NowToYYYYMMddHHmmssMust 将当前时间格式化成 2022-10-12 10:48:11
func (*dateUtil) NowToYYYYMMddHHmmssMust() string {
	return time.Now().Format(DateFormat_yyyy_MM_dd_HH_mm_ss)
}

// GetCurrent24HoursUnix 获取 24 小时内时间戳，beginTime 是24小时前 unix 时间,endTime 是当前 unix 时间戳
func (*dateUtil) GetCurrent24HoursUnix() (beginTime, endTime uint64) {
	endTime = uint64(time.Now().Unix())
	beginTime = uint64(time.Now().AddDate(0, 0, -1).Unix())
	return
}

// GetNowUnixUint64 获取当前时间戳 unix  的 uint64 类型
func (*dateUtil) GetNowUnixUint64() uint64 {
	return uint64(time.Now().Unix())
}

// GetCurrentToDaysUnix 获取当前在 days 偏移后的时间戳，如 days = -1 则是一天前
func (*dateUtil) GetCurrentToDaysUnix(days int) int64 {
	return time.Now().AddDate(0, 0, days).Unix()
}

// GetCurrentToMonthsUnix 获取当前在 months 偏移后的时间戳，如 months = -1 则是一月前
func (*dateUtil) GetCurrentToMonthsUnix(months int) int64 {
	return time.Now().AddDate(0, months, 0).Unix()
}

// GetCurrentToYearsUnix 获取当前在 years 偏移后的时间戳，如 years = -1 则是一年前
func (*dateUtil) GetCurrentToYearsUnix(years int) int64 {
	return time.Now().AddDate(years, 0, 0).Unix()
}

// GetDiffDays 获取t1,t2相差的天数， t2 - t1
func (that *dateUtil) GetDiffDays(t1, t2 int64) int {
	beginTime := time.Unix(t1, 0)
	endTime := time.Unix(t2, 0)
	if endTime.Unix() > that.TodayUnixInt64() {
		endTime = time.Unix(that.TodayUnixInt64(), 0)
	}
	beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, time.Local)
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.Local)
	return int(endTime.Sub(beginTime).Hours() / 24)
}

// GetLastMonths 获取上个月的时间戳, 返回月份 200601
func (that *dateUtil) GetLastMonths(months int) string {
	return time.Now().AddDate(0, months, 0).Format(DateFormatyyyyMM)
}

// ParseToLocalMust date 格式化 format 本地时区
func (*dateUtil) ParseToLocalMust(date, format string) (time.Time, error) {
	n, err := time.ParseInLocation(format, date, time.Local)
	return n, err
}

// ParseToyyyyMMddHHmmssLocalMust 将字符串 20221012120900 格式化成 time.Time 使用本地时区
func (*dateUtil) ParseToyyyyMMddHHmmssLocalMust(date string) time.Time {
	n, _ := time.ParseInLocation(DateFormatyyyyMMddHHmmss, date, time.Local)
	return n
}

// GetLastMonthRangeByTime 根据时间 获取上月第一天，最后一天日期。时间格式 20230915
// firstDay： 20230801
// lastDay: 20230831
func (*dateUtil) GetLastMonthRangeByTime(t int64) (uint64, uint64) {
	now := time.Unix(t, 0)
	// 计算上个月的年份和月份
	lastMonth := now.AddDate(0, -1, 0)
	// 获取上个月的第一天和最后一天
	firstDayOfLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfLastMonth := firstDayOfLastMonth.AddDate(0, 1, -1)
	firstDay, _ := strconv.Atoi(firstDayOfLastMonth.Format(DateFormatyyyyMMdd))
	lastDay, _ := strconv.Atoi(lastDayOfLastMonth.Format(DateFormatyyyyMMdd))
	return uint64(firstDay), uint64(lastDay)
}

// GetLastMonthByTime 根据时间获取上个月
// 200601
func (*dateUtil) GetLastMonthByTime(t int64, months int) string {
	now := time.Unix(t, 0)
	return now.AddDate(0, months, 0).Format(DateFormatyyyyMM)
}

// GetMonthByDay 根据日期获取时间 20230915 202309
func (*dateUtil) GetMonthByDay(date string) string {
	n, _ := time.ParseInLocation(DateFormatyyyyMMddHHmmss, date, time.Local)
	return n.Format(DateFormatyyyyMM)
}

// ParseToyyyyMMddHHmmLocalMust 解析时间 yyyyMMddHHmm
func (*dateUtil) ParseToyyyyMMddHHmmLocalMust(date string) int64 {
	n, _ := time.ParseInLocation(DateFormtyyyyMMddHHmm, date, time.Local)
	return n.Unix()
}

// GetDiffDaysByYyyyMMdd t2 - t1 相差的天数
func (that *dateUtil) GetDiffDaysByYyyyMMdd(beginTime, endTime uint64) int {
	begin, _ := time.ParseInLocation(DateFormatyyyyMMdd, fmt.Sprintf("%d", beginTime), time.Local)
	end, _ := time.ParseInLocation(DateFormatyyyyMMdd, fmt.Sprintf("%d", endTime), time.Local)
	return that.GetDiffDays(begin.Unix(), end.Unix())
}

// ParseToYyyyMMddMust 将字符串 20221012 格式化成 time.Time （使用 Parse 默认 UTC 时区）
func (*dateUtil) ParseToYyyyMMddMust(day uint64) time.Time {
	t, _ := time.Parse(DateFormatyyyyMMdd, fmt.Sprintf("%d", day))
	return t
}

// GetDayTarget unix 时间戳 -> yyyyMMdd
func (that *dateUtil) GetDayTarget(unix int64) int64 {
	return int64(that.FormatUnixToYYYYMMddIntegerMust(unix))
}

// GetUnixRangeFromDayTarget 根据日期获得这一天的起始结束 unix 时间错，如 dayTarget = 20240101
func (*dateUtil) GetUnixRangeFromDayTarget(dayTarget int64) (start int64, end int64, err error) {
	// 将 uint64 转换为字符串
	dateStr := strconv.FormatInt(dayTarget, 10)
	if len(dateStr) != 8 {
		return 0, 0, fmt.Errorf("invalid date format. It must be yyyyMMdd")
	}

	// 解析为 time.Time
	t, err := time.Parse("20060102", dateStr)
	if err != nil {
		return 0, 0, err
	}

	// 起始时间：当天 00:00:00
	startTime := t
	// 结束时间：当天 23:59:59
	endTime := t.Add(24*time.Hour - time.Second)

	return startTime.Unix(), endTime.Unix(), nil
}

// GetRecentDates 获取 -years -months -days 至今的 dayTarget， months = -3 则3个月前至今的 dayTarget
func (*dateUtil) GetRecentDates(years, months, days int) []int {
	var result []int
	end := time.Now()
	start := end.AddDate(years, months, days)

	for t := start; !t.After(end); t = t.AddDate(0, 0, 1) {
		dateInt := t.Year()*10000 + int(t.Month())*100 + t.Day()
		result = append(result, dateInt)
	}

	return result
}

// GetRecentDatesByEnd 获取 -years -months -days 至 end 的 dayTarget， months = -3 则3个月前至今的 dayTarget
func (that *dateUtil) GetRecentDatesByEnd(years, months, days int, end time.Time) []int {
	start := end.AddDate(years, months, days)
	return that.GetRecentDatesByBeginEnd(start, end)
}

// GetRecentDatesByBeginEnd 获取 begin 至 end 的 dayTarget， months = -3 则3个月前至今的 dayTarget
func (*dateUtil) GetRecentDatesByBeginEnd(begin, end time.Time) []int {
	var result []int

	for t := begin; !t.After(end); t = t.AddDate(0, 0, 1) {
		dateInt := t.Year()*10000 + int(t.Month())*100 + t.Day()
		result = append(result, dateInt)
	}

	return result
}

// HourList 获得日的 小时列表 [00,01,...]
func (*dateUtil) HourList(isToday bool) []string {
	hourLst := make([]string, 0)
	maxHour := 24
	if isToday {
		maxHour = time.Now().Hour()
	}

	for i := 0; i <= maxHour; i++ {
		hour := fmt.Sprintf("%d", i)
		if 1 == len(hour) {
			hour = fmt.Sprintf("0%s", hour)
		}
		hourLst = append(hourLst, hour)
	}
	return hourLst
}
