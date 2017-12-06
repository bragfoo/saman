package excel

import (
	"mime/multipart"
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
	"strconv"
	"time"
)

func OpenNewVideo(file multipart.File) {
	xlsx, err := excelize.OpenReader(file)
	if nil != err {
		log.Error(err)
	} else {
		getNewVideoFromData(xlsx)
	}
}

func getNewVideoFromData(f *excelize.File) {
	var l []model.PlayExcel
	var platSum []model.PlatPlayAmount
	var fans []model.FansExcel

	l = append(l, readPlatform(f, 4, "5a169130ef2d1345db33cdc6", sheet)...)  // 空间音乐
	l = append(l, readPlatform(f, 35, "5a1690eeef2d1345d21327e9", sheet)...) // 空间体育
	l = append(l, readPlatform(f, 70, "5a16933cef2d1346121beb0c", sheet)...) // 秒拍
	l = append(l, readPlatformWeibo(f, 91, "5a169398ef2d13461ea201a4", sheet)...) // 微博

	platSum = append(platSum, readPlatformSum(f, sheet, "B2", "5a169130ef2d1345db33cdc6")) // 空间音乐
	platSum = append(platSum, readPlatformSum(f, sheet, "B33", "5a1690eeef2d1345d21327e9")) // 空间体育
	platSum = append(platSum, readPlatformSum(f, sheet, "B68", "5a16933cef2d1346121beb0c")) // 秒拍

	fans = append(fans, getNewVideoFans(f, sheet, 22, "5a169130ef2d1345db33cdc6")...) // 空间音乐
	fans = append(fans, getNewVideoFans(f, sheet, 57, "5a1690eeef2d1345d21327e9")...) // 空间体育

	saveVideoData(l, platSum)
	saveFans(fans)
}

func getNewVideoFans(f *excelize.File, sheet string, startAxis int, platType string) []model.FansExcel {
	var fans []model.FansExcel
	for i := startAxis; i <= startAxis+7; i++ {
		offset := strconv.Itoa(i)
		b, _ := strconv.Atoi(f.GetCellValue(sheet, "B"+offset))
		c, _ := strconv.Atoi(f.GetCellValue(sheet, "C"+offset))
		d, _ := strconv.Atoi(f.GetCellValue(sheet, "D"+offset))
		e, _ := strconv.Atoi(f.GetCellValue(sheet, "E"+offset))
		t, _ := time.Parse(dateStr, f.GetCellValue(sheet, "A"+offset))
		fans = append(fans, model.FansExcel{
			Increase:    b,
			Cancel:      c,
			NetIncrease: d,
			Total:       e,
			CreateTime:  t.Unix(),
			Type:        platType,
		})
	}
	return fans
}
