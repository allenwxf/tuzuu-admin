package time

import (
	"fmt"
	"strings"
	"testing"
)

/*
*测试Log函数
 */
func Test_DataToTimeStamp(t *testing.T) {
	t1 := DataToTimeStamp("2018-02-07")
	fmt.Println(t1)
	year, month, _ := GetYMD()
	fmt.Println(GetDayByMonth(year, month))
	//fmt.Println(GetYMD())
}

func TestGetZeroTimeStamp(t *testing.T) {
	str := ""
	strslice := strings.Split(str, "|")
	strslice = append(strslice, "yuansudong")

	fmt.Println(strings.Join(strslice, "|"))
}
func Test_Calc(t *testing.T) {
	mid := "02:00"
	start := DataToTime("2018-08-06T17:26")
	end := DataToTime("2018-08-08T01:30")
	days := CalcDay(int64(start), int64(end), mid)
	fmt.Println(days)
}
func Test_TimeToData(t *testing.T) {
	fmt.Println(TimeToData4(int(CurTimeStamp())))
}
func Test_TimeToData5(t *testing.T) {
	fmt.Println(TimeToData5(1534746585))
}
func Test_DataToTime2(t *testing.T) {
	fmt.Println(DataToTime2("2018-09-05T17:31:25"))
}

func Test_DTimeToData3(t *testing.T) {
	fmt.Println(TimeToData3(1545026460))
}
