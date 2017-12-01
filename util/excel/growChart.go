package excel

import (
	"mime/multipart"
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
	"strconv"
	"github.com/bragfoo/saman/util/db"
	"github.com/bragfoo/saman/util"
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
	m.CreateTime = getDateFromString(d).Unix()
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
