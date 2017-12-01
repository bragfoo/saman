package main

import (
	"github.com/xuri/excelize"
	"github.com/bragfoo/saman/backend/api/model"
	"strings"
	"strconv"
	"time"
	"fmt"
)

func main() {
	//read()
	test()
}

func read() {
	f, _ := excelize.OpenFile("/Users/kevin1993/project/xlsx/mt.xlsx")
	readFansColumn(f, "D3", "A3", "59fae294ef2d1314e0ea2a57")
}

func readFansColumn(f *excelize.File, axis string, timeAxis string, platType string) model.FansExcel {
	var m model.FansExcel
	m.Total = processSum(f.GetCellValue("Sheet1", axis))
	m.Type = platType
	m.CreateTime = getDate(f, "Sheet1", timeAxis).Unix()
	return m
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

func getDate(f *excelize.File, sheet string, axis string) time.Time {
	sheetTime := f.GetCellValue(sheet, axis)
	timeStrs := strings.Split(sheetTime, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[1])
	month, _ := strconv.Atoi(t[0])
	myTime := time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.Local)
	return myTime
}

func test() {
	f, _ := excelize.OpenFile("/Users/kevin1993/project/xlsx/videoSum.xlsx")
	rows := f.GetRows("A图")
	for _, v := range rows {
		fmt.Println(v[0])
		fmt.Println(v[1])
	}
}
