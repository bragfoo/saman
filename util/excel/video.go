package excel

import (
	"mime/multipart"
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
	"strconv"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/util/db"
	"github.com/bragfoo/saman/util"
	"time"
)

var sheet = "Sheet1"

func OpenPlatformVideo(f multipart.File) {
	xlsx, err := excelize.OpenReader(f)
	if nil != err {
		log.Error(err)
	} else {
		getVideoPlatformData(xlsx)
	}
}

func getVideoPlatformData(f *excelize.File) {
	var l []model.PlayExcel
	var platSum []model.PlatPlayAmount
	l = append(l, readPlatform(f, 4, "59fae2ecef2d1314e0ea2a5b", sheet)...)  // bilibili
	l = append(l, readPlatform(f, 11, "59fae313ef2d1314e0ea2a5c", sheet)...) // iqiyi
	l = append(l, readPlatform(f, 18, "59fae32cef2d1314e0ea2a5d", sheet)...) // youku
	l = append(l, readPlatform(f, 25, "5a1cfaf5ef2d13bca79f17ce", sheet)...) //yidianzixun
	l = append(l, readPlatform(f, 32, "5a1cfb70ef2d13bcb31bec01", sheet)...) //uc
	l = append(l, readPlatform(f, 39, "59fae351ef2d1314e0ea2a5f", sheet)...) //renren

	platSum = append(platSum, readPlatformSum(f, sheet, "B2", "59fae2ecef2d1314e0ea2a5b"))  //bilibili
	platSum = append(platSum, readPlatformSum(f, sheet, "B9", "59fae313ef2d1314e0ea2a5c"))  //iqiyi
	platSum = append(platSum, readPlatformSum(f, sheet, "B16", "59fae32cef2d1314e0ea2a5d")) //youku
	platSum = append(platSum, readPlatformSum(f, sheet, "B23", "5a1cfaf5ef2d13bca79f17ce")) //yidianzixun
	platSum = append(platSum, readPlatformSum(f, sheet, "B30", "5a1cfb70ef2d13bcb31bec01")) //uc
	platSum = append(platSum, readPlatformSum(f, sheet, "B37", "59fae351ef2d1314e0ea2a5f")) //renren
	saveVideoData(l, platSum)
}

func saveVideoData(pe []model.PlayExcel, pa []model.PlatPlayAmount) {
	for _, v := range pe {
		saveVideo(v)
	}
	for _, v := range pa {
		savePlatPlayAmount(v)
	}
}

func readPlatform(file *excelize.File, start int, platType string, sheet string) ([]model.PlayExcel) {
	var l []model.PlayExcel

	for i := start; i < start+3; i++ {
		offset := strconv.Itoa(i)
		if "" == file.GetCellValue(sheet, "B"+offset) {
			break
		} else {
			l = append(l, model.PlayExcel{
				IType:      platType,
				Title:      file.GetCellValue(sheet, "B"+offset),
				Link:       file.GetCellValue(sheet, "C"+offset),
				PlayAmount: processSum(file.GetCellValue(sheet, "D"+offset)),
				CreateTime: getDate(file, sheet, "A"+offset).Unix(),
			})
		}

	}

	return l
}
func readPlatformWeibo(file *excelize.File, start int, platType string, sheet string) ([]model.PlayExcel) {
	var l []model.PlayExcel

	for i := start; i < start+3; i++ {
		offset := strconv.Itoa(i)
		if "" == file.GetCellValue(sheet, "B"+offset) {
			break
		} else {
			l = append(l, model.PlayExcel{
				IType:      platType,
				Title:      file.GetCellValue(sheet, "B"+offset),
				Link:       file.GetCellValue(sheet, "C"+offset),
				PlayAmount: processSum(file.GetCellValue(sheet, "C"+offset)),
				CreateTime: getDate(file, sheet, "A"+offset).Unix(),
			})
		}

	}

	return l
}

func readPlatformSum(f *excelize.File, sheet string, axis string, platType string) model.PlatPlayAmount {
	sum, grow := getTotalPlayAmount(f, sheet, axis)
	var date time.Time
	if "Sheet2" == sheet {
		date = getDate(f, sheet, "A22")
	} else {
		date = getDate(f, sheet, "A2")
	}

	return model.PlatPlayAmount{
		CreateTime: date.Unix(),
		PlatType:   platType,
		Sum:        sum,
		Grow:       grow,
	}
}

func savePlatPlayAmount(amount model.PlatPlayAmount) {
	stmIns, _ := db.Prepare(insertPlatPlayAMount)
	defer stmIns.Close()
	stmUpd, _ := db.Prepare(updatePlatPlayAmount)
	defer stmUpd.Close()
	var lineSum = 0
	rows, err := db.Query(testPlatPlayAmount, amount.PlatType, amount.CreateTime)
	defer rows.Close()
	if nil != err {
		log.Error(err)
	} else {
		for rows.Next() {
			lineSum++
			var m = model.PlatPlayAmount{}
			err := rows.Scan(&m.Ids)
			if nil == err {
				stmUpd.Exec(amount.Grow, amount.Sum, m.Ids)
			}
		}
	}
	if lineSum == 0 {
		stmIns.Exec(util.GetObjectId(), amount.CreateTime, amount.PlatType, amount.Sum, amount.Grow)
	}
}

func saveVideo(e model.PlayExcel) {
	var videoLineSum = 0
	rows, err := db.Query(testVideo, e.Title)
	defer rows.Close()
	if nil != err {
		log.Error(err)
	} else {
		for rows.Next() {
			videoLineSum++
			var m = model.Video{}
			err := rows.Scan(&m.Ids)
			if nil == err {
				updatePlayAmount(m.Ids, e)
			}
		}
	}
	if videoLineSum == 0 {
		savePlayAmount(e)
	}
}

func savePlayAmount(e model.PlayExcel) {
	var ids string
	ids = util.GetObjectId()
	stmInsVideo, _ := db.Prepare(insVideo)
	defer stmInsVideo.Close()
	stmInsPlayAmount, _ := db.Prepare(insPlayAMount)
	defer stmInsPlayAmount.Close()
	_, err := stmInsVideo.Exec(ids, e.IType, e.Title, e.Link, e.CreateTime)
	if nil != err {
		log.Error(err)
	}
	_, errIns := stmInsPlayAmount.Exec(util.GetObjectId(), ids, e.CreateTime, e.PlayAmount)
	if nil != errIns {
		log.Error(errIns)
	}
}

func updatePlayAmount(videoIds string, e model.PlayExcel) {
	stmTest, _ := db.Prepare(testPlayAmount)
	defer stmTest.Close()
	stmUpd, _ := db.Prepare(updPlayAmount)
	defer stmUpd.Close()
	stmIns, _ := db.Prepare(insPlayAMount)
	defer stmIns.Close()
	var lineSum = 0
	rows, err := stmTest.Query(videoIds, e.CreateTime)
	defer rows.Close()
	if nil != err {
		log.Error(err)
	} else {
		for rows.Next() {
			lineSum++
			var m = model.PlayAmount{}
			err := rows.Scan(&m.Ids)
			if nil == err {
				log.Info(e.PlayAmount)
				stmUpd.Exec(e.PlayAmount, videoIds, e.CreateTime)
			}
		}
	}
	if lineSum == 0 {
		stmIns.Exec(util.GetObjectId(), videoIds, e.CreateTime, e.PlayAmount)
	}
}
