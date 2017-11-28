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

func getTotalPlayAmount(f *excelize.File, sheet string, axis string) (sum int, grow int) {
	str := f.GetCellValue(sheet, axis)
	sumStrs := strings.Split(str, "、")
	total := strings.Split(sumStrs[0], "：")[1]
	grows := strings.Split(sumStrs[1], "：")[1]
	return processSum(total), processSum(grows)
}

func processSum(sum string) int {
	var convertStr string
	if strings.Contains(sum, ",") {
		convertStr = strings.Replace(sum, ",", "", -1)
	} else if strings.Contains(sum, "万") {
		convertStr = strings.Replace(sum, "万", "0000", -1)

	}
	ret, _ := strconv.Atoi(convertStr)
	return ret
}
