package time

//所有和时间相关的函数

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//将2006-01-02T15:04格式转换成时间戳
func DataToTime(data string) int {
	s := strings.Split(data, "T")
	toBeCharge := s[0] + " " + s[1] + ":00"                         //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	return int(sr)
}

//将2006-01-02T15:04:05格式转换成时间戳
func DataToTime2(data string) int {
	s := strings.Split(data, "T")
	toBeCharge := s[0] + " " + s[1]                                 //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	return int(sr)
}

//将时间戳转换成2006-01-02T15:04格式的日期
func TimeToData(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	//s := strings.Split(tm.Format("2006-01-02 15:04:05")," ")
	return tm.Format("2006-01-02T15:04")
}

//将时间戳转换成2006-01-02 15:04:05格式的日期
func TimeToData2(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	//s := strings.Split(tm.Format("2006-01-02 15:04:05")," ")
	return tm.Format("2006-01-02 15:04:05")
}

//将时间戳转换成2006-01-02 15:04格式的日期
func TimeToData3(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	//s := strings.Split(tm.Format("2006-01-02 15:04")," ")
	return tm.Format("2006-01-02 15:04")
}

//将时间戳转换成2006/01/02 15:04格式的日期
func TimeToData4(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	//s := strings.Split(tm.Format("2006-01-02 15:04")," ")
	return tm.Format("2006/01/02 15:04")
}

//将时间戳转换成20060102150405格式的日期
func TimeToData5(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("20060102150405")
}

//将时间戳转换成200601021504格式的日期
func TimeToData6(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("200601021504")
}

// 当前的时间字符串 格式 "2006-01-02 15:04:05"
func CurTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 当前的时间字符串 格式 "2006-01-02 15:04:05.000"
func CurTimeMsStr() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

//获取当前GMT时区的时间格式为2006-01-02T15:04:05Z的日期
func CurTimeGMT(z int) string {
	var cstZone = time.FixedZone("CST", z*3600)
	return time.Now().In(cstZone).Format("2006-01-02T15:04:05Z")
}

// 当前的时间戳，单位秒
func CurTimeStamp() int64 {
	return time.Now().Unix()
}

//获取年月日
func GetYMD() (int, int, int) {
	t := time.Now()
	return t.Year(), int(t.Month()), t.Day()
}

//将2006-01-02格式转换成时间戳
func DataToTimeStamp(data string) int {
	timeLayout := "2006-01-02"                                //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, data, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                      //转化为时间戳 类型是int64
	return int(sr)
}

//获取每月的天数
func GetDayByMonth(year, month int) (day int) {
	if month == 2 {
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			//闰年
			day = 29
			return
		} else {
			day = 28
			return
		}
	}
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		day = 31
		return
	}
	day = 30
	return
}

// 返回当前的毫秒数
func CurMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetZeroTimestamp 获取当天凌晨的时间戳
func GetZeroTimestamp(timestamp int64) int64 {
	Time := time.Unix(timestamp, 0)
	return time.Date(Time.Year(), Time.Month(), Time.Day(), 0, 0, 0, 0, time.Local).Unix()
}

// GetDayLastTimestamp获取当天最后的时间戳
func GetDayLastStamp(timestamp int64) int64 {
	Time := time.Unix(timestamp, 0)
	return time.Date(Time.Year(), Time.Month(), Time.Day(), 0, 0, 0, 0, time.Local).Unix() + 86400
}

/*计算天数（主要是用于计算酒店夜审天数）
*夜审时间必须为hh:mm格式
 */
func CalcDay(startstamp, endstamp int64, checktime string) int {
	starTime := time.Unix(int64(startstamp), 0)
	endTime := time.Unix(int64(endstamp), 0)
	s_year := strconv.Itoa(starTime.Year())
	s_month := compare(int(starTime.Month()))
	s_day := compare(starTime.Day())
	e_year := strconv.Itoa(endTime.Year())
	e_month := compare(int(endTime.Month()))
	e_day := compare(endTime.Day())
	smidstr := fmt.Sprintf("%s-%s-%sT%s", s_year, s_month, s_day, checktime)
	smidtime := DataToTime(smidstr)
	//先算天数
	startT := DataToTimeStamp(fmt.Sprintf("%s-%s-%s", s_year, s_month, s_day))
	endT := DataToTimeStamp(fmt.Sprintf("%s-%s-%s", e_year, e_month, e_day))
	Day_ := (endT - startT) / (24 * 60 * 60)
	if startstamp < int64(smidtime) {
		Day_ += 1
	}
	return Day_
}
func compare(a int) string {
	if a < 10 {
		return fmt.Sprintf("0%d", a)
	} else {
		return fmt.Sprintf("%d", a)
	}
}
