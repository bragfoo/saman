package excel

import (
	"mime/multipart"
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
	"strconv"
	"github.com/bragfoo/saman/util/db"
	"github.com/bragfoo/saman/util"
	"fmt"
	"strings"
)

func OpenCharts(file multipart.File) {
	xlsx, err := excelize.OpenReader(file)
	if nil != err {
		log.Error(err)
	} else {
		getCharts(xlsx)
	}
}

func getCharts(file *excelize.File) {
	var l []model.WeekGrow
	l = append(l, readCharts(file, "头条音乐图", "59fae20cef2d1314e0ea2a55", 1)...)
	l = append(l, readCharts(file, "头条体育图", "59fae276ef2d1314e0ea2a56", 1)...)
	l = append(l, readCharts(file, "腾讯音乐图", "59fae294ef2d1314e0ea2a57", 1)...)
	l = append(l, readCharts(file, "腾讯体育图", "59fae2a8ef2d1314e0ea2a58", 1)...)
	l = append(l, readCharts(file, "空间音乐图", "5a169130ef2d1345db33cdc6", 3)...)
	l = append(l, readCharts(file, "空间体育图", "5a1690eeef2d1345d21327e9", 0)...)
	l = append(l, readCharts(file, "秒拍图", "5a16933cef2d1346121beb0c", 0)...)
	l = append(l, readCharts(file, "A图", "59fae2d4ef2d1314e0ea2a5a", 1)...)
	l = append(l, readCharts(file, "B图", "59fae2ecef2d1314e0ea2a5b", 1)...)
	l = append(l, readCharts(file, "爱奇艺图", "59fae313ef2d1314e0ea2a5c", 1)...)
	l = append(l, readCharts(file, "优酷图", "59fae32cef2d1314e0ea2a5d", 1)...)
	l = append(l, readCharts(file, "人人图", "59fae351ef2d1314e0ea2a5f", 1)...)
	l = append(l, readCharts(file, "UC图", "5a2103a6ef2d13067d1b9e23", 1)...)
	l = append(l, readCharts(file, "斗鱼图", "59fae33fef2d1314e0ea2a5e", 1)...)

	saveCharts(l)

	var list []model.PlayExcel
	var platSum []model.PlatPlayAmount

	var li []model.PlayExcel
	var pl model.PlatPlayAmount

	li, pl = getLastWeekPlayed(file, "59fae20cef2d1314e0ea2a55", "头条音乐")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae276ef2d1314e0ea2a56", "头条体育")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae294ef2d1314e0ea2a57", "腾讯音乐")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae2a8ef2d1314e0ea2a58", "腾讯体育")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "5a169130ef2d1345db33cdc6", "公共空间音乐")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "5a1690eeef2d1345d21327e9", "公共空间体育")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "5a16933cef2d1346121beb0c", "秒拍")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae2d4ef2d1314e0ea2a5a", "A站")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae2ecef2d1314e0ea2a5b", "B站")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae313ef2d1314e0ea2a5c", "爱奇艺")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae2a8ef2d1314e0ea2a58", "人人")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "5a2103a6ef2d13067d1b9e23", "UC")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae33fef2d1314e0ea2a5e", "斗鱼")
	list = append(list, li...)
	platSum = append(platSum, pl)
	li, pl = getLastWeekPlayed(file, "59fae32cef2d1314e0ea2a5d", "优酷")
	list = append(list, li...)
	platSum = append(platSum, pl)

	saveVideoData(list, platSum)

}

func readCharts(f *excelize.File, sheet string, platIds string, offset int) []model.WeekGrow {
	var l []model.WeekGrow
	rows := f.GetRows(sheet)
	for k, v := range rows {
		if k <= offset {
			continue
		}
		if len(v) >= 2 {
			if "" == v[0] || "" == v[1] {
				continue
			}
			l = append(l, readChartsColumn(v[0], v[1], platIds))
		}

	}
	return l
}

func readChartsColumn(d string, s string, p string) model.WeekGrow {
	var m = model.WeekGrow{}
	sum, _ := strconv.Atoi(s)
	//m.CreateTime = getDateFromString(d).Unix()
	m.CreateTime = getDateFromStringWithYear(d).Unix()
	m.Grow = sum
	m.PlatIds = p
	return m

}

func saveCharts(l []model.WeekGrow) {
	for _, v := range l {
		updateCharts(v)
	}
}

func updateCharts(amount model.WeekGrow) {
	stmIns, _ := db.Prepare(insertPlatPlayAMount)
	defer stmIns.Close()
	stmUpd, _ := db.Prepare(updatePlatPlayAmount)
	defer stmUpd.Close()
	var lineSum = 0
	rows, err := db.Query(testPlatPlayAmount, amount.PlatIds, amount.CreateTime)
	defer rows.Close()
	if nil != err {
		log.Error(err)
	} else {
		for rows.Next() {
			lineSum++
			var m = model.PlatPlayAmount{}
			err := rows.Scan(&m.Ids)
			if nil == err {
				stmUpd.Exec(amount.Grow, 0, m.Ids)
			}
		}
	}
	if lineSum == 0 {
		stmIns.Exec(util.GetObjectId(), amount.CreateTime, amount.PlatIds, 0, amount.Grow)
	}
}

func getLastWeekPlayed(f *excelize.File, plat string, sheet string) ([]model.PlayExcel, model.PlatPlayAmount) {
	var createTime int64
	var total int
	var grow int
	var list []model.PlayExcel
	var platAmount = model.PlatPlayAmount{}
	d := f.GetRows(sheet)
	l := len(d)
	offset := l - 5
	for k, v := range d {
		if k < offset {
			continue
		} else if k == offset {
			createTime, total, grow = getMetaData(v)
		} else if k == offset+1 {
			continue
		} else {
			list = append(list, readPlayVideo(v, plat, createTime))
		}
	}

	platAmount.CreateTime = createTime
	platAmount.Sum = total
	platAmount.Grow = grow
	platAmount.PlatType = plat

	fmt.Println(list)
	fmt.Println(platAmount)

	return list, platAmount
}

func readPlayVideo(s []string, plat string, createTime int64) (model.PlayExcel) {
	return model.PlayExcel{
		IType:      plat,
		Title:      s[1],
		Link:       s[2],
		PlayAmount: processSum(s[3]),
		CreateTime: createTime,
	}
}

func getMetaData(s []string) (createTime int64, sum int, grow int) {
	var str = s[1]
	var strTime = s[0]
	var sumStrs []string

	if strings.Contains(str, "、") {
		sumStrs = strings.Split(str, "、")
	} else {
		sumStrs = strings.Split(str, "，")
	}
	total := strings.Split(sumStrs[0], "：")[1]
	grows := strings.Split(sumStrs[1], "：")[1]
	time := getDateFromString(strTime)
	return time.Unix(), processSum(total), processSum(grows)
}
