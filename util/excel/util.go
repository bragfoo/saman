package excel

import (
	"github.com/xuri/excelize"
	"time"
	"strings"
	"strconv"
)

func getDate(f *excelize.File, sheet string, axis string) time.Time {
	sheetTime := f.GetCellValue(sheet, axis)
	timeStrs := strings.Split(sheetTime, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[1])
	month, _ := strconv.Atoi(t[0])
	myTime := time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.Local)
	return myTime
}

func getDateWithYear(f *excelize.File, sheet string, axis string) time.Time {
	// 2016.2.13-2.19
	sheetTime := f.GetCellValue(sheet, axis)
	timeStrs := strings.Split(sheetTime, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[2])
	month, _ := strconv.Atoi(t[1])
	year, _ := strconv.Atoi(t[0])
	myTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return myTime
}

func getDateFromString(s string) time.Time {
	timeStrs := strings.Split(s, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[1])
	month, _ := strconv.Atoi(t[0])
	myTime := time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.Local)
	return myTime
}
func getDateFromStringWithYear(s string) time.Time {
	timeStrs := strings.Split(s, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[2])
	month, _ := strconv.Atoi(t[1])
	year, _ := strconv.Atoi(t[0])
	myTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return myTime
}

func getTotalPlayAmount(f *excelize.File, sheet string, axis string) (sum int, grow int) {
	str := f.GetCellValue(sheet, axis)
	var sumStrs []string
	if strings.Contains(str, "、") {
		sumStrs = strings.Split(str, "、")
	} else {
		sumStrs = strings.Split(str, "，")
	}
	total := strings.Split(sumStrs[0], "：")[1]
	grows := strings.Split(sumStrs[1], "：")[1]
	return processSum(total), processSum(grows)
}

func processSum(sum string) int {
	var convertStr string
	if strings.Contains(sum, ",") {
		convertStr = strings.Replace(sum, ",", "", -1)
	} else if strings.Contains(sum, "万") {
		if strings.Contains(sum, ".") {
			s := strings.Replace(sum, "万", "", -1)
			f, _ := strconv.ParseFloat(s, 64)
			return int(f * 10000)
		} else {
			convertStr = strings.Replace(sum, "万", "0000", -1)
		}
	} else {
		convertStr = sum
	}
	ret, _ := strconv.Atoi(convertStr)
	return ret
}
