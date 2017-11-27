package excel

import (
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
	"fmt"
	"mime/multipart"
	"github.com/bragfoo/saman/backend/api/model"
	"strings"
	"strconv"
	"time"
	"github.com/bragfoo/saman/util/db"
	"github.com/bragfoo/saman/util"
)

var dateStr = "2006-01-02"
var testQuery = "SELECT pF.ids AS ids,count(ids) AS sum FROM platformFans pF WHERE 1=1 AND pF.createTime = ?  AND pF.platType= ? GROUP BY pF.ids"
var insFansQuery = "INSERT INTO platformFans (ids, createTime, sum, decrease, increase, platType) VALUES (?,?,?,?,?,?)"
var updateQuery = "UPDATE platformFans p SET p.createTime = ?, p.sum          = ?, p.decrease     = ?, p.increase     = ?, p.platType     = ? WHERE p.ids = ?;"

var videoTestQuery = "SELECT ids FROM video WHERE platIds = ? AND title = ?"
var videoInsertQuery = "INSERT INTO video (ids, platIds, title, link, createTime) VALUES (?, ?, ?, ?, ?)"
var videoUpdateQuery = ""

var playAmountQuery = "INSERT INTO playAmount (ids, videoIds, createTime, sum) VALUES (?,?,?,?);"

func Init() {
	xlsx, err := excelize.OpenFile("/Users/kevin1993/sheet.xlsx")
	if nil != err {
		log.Error(err)
	} else {
		fmt.Println(xlsx.GetCellValue("Sheet1", "B4"))
		fmt.Println(xlsx.GetCellValue("Sheet1", "B5"))
		fmt.Println(xlsx.GetCellValue("Sheet1", "B6"))
	}
}

func main() {
	Init()
}

func OpenSpace(file multipart.File) {
	xlsx, err := excelize.OpenReader(file)
	if nil != err {
		log.Error(err)
	} else {
		GetSpaceData(xlsx)
	}
}

func OpenVideo(file multipart.File) {
	xlsx, err := excelize.OpenReader(file)
	if nil != err {
		log.Error(err)
	} else {
		getVideoData(xlsx)
	}
}

func GetSpaceData(file *excelize.File) {
	var l []model.PlayExcel
	var lSum []model.SumExcel

	dataTime1 := file.GetCellValue("Sheet1", "A2")
	timeStrs := strings.Split(dataTime1, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[1])
	month, _ := strconv.Atoi(t[0])
	itime := time.Date(2017, time.Month(month), day, 0, 0, 0, 0, time.Local)

	for i := 4; i < 7; i++ {
		offset := strconv.Itoa(i)
		l = append(l, model.PlayExcel{
			IType:      "5a169130ef2d1345db33cdc6",
			Title:      file.GetCellValue("Sheet1", "B"+offset),
			Link:       file.GetCellValue("Sheet1", "C"+offset),
			PlayAmount: file.GetCellValue("Sheet1", "D"+offset),
			CreateTime: itime.Unix(),
		})
	}

	for i := 35; i < 38; i++ {
		offset := strconv.Itoa(i)
		l = append(l, model.PlayExcel{
			IType:      "5a1690eeef2d1345d21327e9",
			Title:      file.GetCellValue("Sheet1", "B"+offset),
			Link:       file.GetCellValue("Sheet1", "C"+offset),
			PlayAmount: file.GetCellValue("Sheet1", "D"+offset),
			CreateTime: itime.Unix(),
		})
	}

	for i := 70; i <= 73; i++ {
		offset := strconv.Itoa(i)
		l = append(l, model.PlayExcel{
			IType:      "5a16933cef2d1346121beb0c",
			Title:      file.GetCellValue("Sheet1", "B"+offset),
			Link:       file.GetCellValue("Sheet1", "C"+offset),
			PlayAmount: file.GetCellValue("Sheet1", "D"+offset),
			CreateTime: itime.Unix(),
		})
	}

	for i := 91; i < 94; i++ {
		offset := strconv.Itoa(i)
		l = append(l, model.PlayExcel{
			IType:      "5a169398ef2d13461ea201a4",
			Title:      file.GetCellValue("Sheet1", "B"+offset),
			Link:       "weibo",
			PlayAmount: file.GetCellValue("Sheet1", "C"+offset),
			CreateTime: itime.Unix(),
		})
	}

	d := file.GetCellValue("Sheet1", "B2")
	s := strings.Split(d, "，")
	s1 := strings.Split(s[0], "：")
	s2 := strings.Split(s[1], "：")
	i1, _ := strconv.Atoi(strings.TrimSpace(s1[1]))
	i2, _ := strconv.Atoi(strings.TrimSpace(s2[1]))
	lSum = append(lSum, model.SumExcel{
		Type:       "music",
		Total:      i1,
		Grow:       i2,
		CreateDate: itime.Unix(),
	})

	dataB := file.GetCellValue("Sheet1", "B33")
	stringB := strings.Split(dataB, "，")
	sb1 := strings.Split(stringB[0], "：")
	sb2 := strings.Split(stringB[1], "：")
	ib1, _ := strconv.Atoi(strings.TrimSpace(sb1[1]))
	ib2, _ := strconv.Atoi(strings.TrimSpace(sb2[1]))
	lSum = append(lSum, model.SumExcel{
		Type:       "music",
		Total:      ib1,
		Grow:       ib2,
		CreateDate: itime.Unix(),
	})

	file.GetCellValue("Sheet1", "A22")
	file.GetCellValue("Sheet1", "B22")
	file.GetCellValue("Sheet1", "C22")
	file.GetCellValue("Sheet1", "D22")

	var fans []model.FansExcel
	for i := 22; i < 29; i++ {
		offset := strconv.Itoa(i)
		b, _ := strconv.Atoi(file.GetCellValue("Sheet1", "B"+offset))
		c, _ := strconv.Atoi(file.GetCellValue("Sheet1", "C"+offset))
		d, _ := strconv.Atoi(file.GetCellValue("Sheet1", "D"+offset))
		e, _ := strconv.Atoi(file.GetCellValue("Sheet1", "E"+offset))
		t, _ := time.Parse(dateStr, file.GetCellValue("Sheet1", "A"+offset))
		fans = append(fans, model.FansExcel{
			Increase:    b,
			Cancel:      c,
			NetIncrease: d,
			Total:       e,
			CreateTime:  t.Unix(),
			Type:        "5a169130ef2d1345db33cdc6", //空间音乐
		})
	}

	for i := 57; i < 63; i++ {
		offset := strconv.Itoa(i)
		b, _ := strconv.Atoi(file.GetCellValue("Sheet1", "B"+offset))
		c, _ := strconv.Atoi(file.GetCellValue("Sheet1", "C"+offset))
		d, _ := strconv.Atoi(file.GetCellValue("Sheet1", "D"+offset))
		e, _ := strconv.Atoi(file.GetCellValue("Sheet1", "E"+offset))
		t, _ := time.Parse(dateStr, file.GetCellValue("Sheet1", "A"+offset))
		fans = append(fans, model.FansExcel{
			Increase:    b,
			Cancel:      c,
			NetIncrease: d,
			Total:       e,
			CreateTime:  t.Unix(),
			Type:        "5a1690eeef2d1345d21327e9", //空间体育
		})
	}

	fmt.Println(lSum)
	fmt.Println(l)
	fmt.Println(fans)
	dataSave(fans, l)

}

func getUGCData(file *excelize.File)  {
	for i:=5;i<10 ;i++  {

	}
}

func getVideoData(file *excelize.File) {
	var l []model.PlayExcel

	dataTime1 := file.GetCellValue("Sheet1", "A2")
	timeStrs := strings.Split(dataTime1, "-")
	t := strings.Split(timeStrs[0], ".")
	day, _ := strconv.Atoi(t[1])
	month, _ := strconv.Atoi(t[0])
	itime := time.Date(2017, time.Month(month), day, 0, 0, 0, 0, time.Local)
	//播放量1
	for i := 4; i < 7; i++ {
		offset := strconv.Itoa(i)
		l = append(l, model.PlayExcel{
			IType:      "5a169130ef2d1345db33cdc6",
			Title:      file.GetCellValue("Sheet1", "B"+offset),
			Link:       file.GetCellValue("Sheet1", "C"+offset),
			PlayAmount: file.GetCellValue("Sheet1", "D"+offset),
			CreateTime: itime.Unix(),
		})
	}

	//粉丝
	var fans []model.FansExcel
	for i := 22; i < 29; i++ {
		offset := strconv.Itoa(i)
		b, _ := strconv.Atoi(file.GetCellValue("Sheet1", "B"+offset))
		c, _ := strconv.Atoi(file.GetCellValue("Sheet1", "C"+offset))
		d, _ := strconv.Atoi(file.GetCellValue("Sheet1", "D"+offset))
		e, _ := strconv.Atoi(file.GetCellValue("Sheet1", "E"+offset))
		t, _ := time.Parse(dateStr, file.GetCellValue("Sheet1", "A"+offset))
		fans = append(fans, model.FansExcel{
			Increase:    b,
			Cancel:      c,
			NetIncrease: d,
			Total:       e,
			CreateTime:  t.Unix(),
			Type:        "5a169130ef2d1345db33cdc6", //空间音乐
		})
	}
}

func dataSave(fans []model.FansExcel, plays []model.PlayExcel) {

	stms, _ := db.Prepare(insFansQuery)
	stm, _ := db.Prepare(testQuery)
	stmu, _ := db.Prepare(updateQuery)

	for _, v := range fans {
		r, err := stm.Query(v.CreateTime, v.Type)
		defer r.Close()
		if nil != err {
			stms.Exec(util.GetObjectId(), v.CreateTime, v.Total, v.Cancel, v.Increase, v.Type)
		} else {
			rowLength := 0
			for r.Next() {
				rowLength++
				var ids string
				err := r.Scan(&ids)
				if nil == err {
					stmu.Exec(v.CreateTime, v.Total, v.Cancel, v.Increase, v.Type, ids)
				}
			}
			if rowLength == 0 {
				stms.Exec(util.GetObjectId(), v.CreateTime, v.Total, v.Cancel, v.Increase, v.Type)
			}
		}
	}

	stm, err := db.Prepare(videoTestQuery)
	stmIns, err := db.Prepare(videoInsertQuery)
	//stmUpd, err := db.Prepare(videoUpdateQuery)
	stmPlayAMountUpd, err := db.Prepare(playAmountQuery)

	for _, v := range plays {

		if nil != err {

		} else {
			rows, _ := stm.Query(v.IType, v.Title)
			defer rows.Close()
			rowLength := 0
			for rows.Next() {
				rowLength++
				var ids string
				err := rows.Scan(&ids)
				if nil == err {
					//update
					//stmUpd.Exec()
					//ins to play amount
				}
			}
			if rowLength == 0 {
				//insert
				var ids string
				ids = util.GetObjectId()
				stmIns.Exec(ids, v.IType, v.Title, v.Link, v.CreateTime)
				//ins to play amount
				stmPlayAMountUpd.Exec(util.GetObjectId(), ids, v.CreateTime, v.PlayAmount)
			}
		}
	}
}
