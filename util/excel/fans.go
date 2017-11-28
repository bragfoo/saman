package excel

import (
	"mime/multipart"
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
)

func OpenFans(file multipart.File) {
	xlsx, err := excelize.OpenReader(file)
	if nil != err {
		log.Error(err)
	} else {
		getVideoFromFans(xlsx) // Sheet2's data
		getFansData(xlsx)
	}
}

func getFansData(f *excelize.File) {
	var l []model.FansExcel
	l = append(l, readFansColumn(f, "B3", "A3", "59fae20cef2d1314e0ea2a55")) //toutiaoMusic
	l = append(l, readFansColumn(f, "C3", "A3", "59fae276ef2d1314e0ea2a56")) //toutiaoSports
	l = append(l, readFansColumn(f, "D3", "A3", "59fae294ef2d1314e0ea2a57")) //tMusic
	l = append(l, readFansColumn(f, "E3", "A3", "59fae2a8ef2d1314e0ea2a58")) //tSports
	l = append(l, readFansColumn(f, "F3", "A3", "59fae2beef2d1314e0ea2a59")) //qzone
	l = append(l, readFansColumn(f, "G3", "A3", "59fae2d4ef2d1314e0ea2a5a")) //acfun
	l = append(l, readFansColumn(f, "H3", "A3", "59fae2ecef2d1314e0ea2a5b")) //bilibili
	l = append(l, readFansColumn(f, "I3", "A3", "59fae313ef2d1314e0ea2a5c")) //iqiyi
	l = append(l, readFansColumn(f, "J3", "A3", "59fae32cef2d1314e0ea2a5d")) //youku
	l = append(l, readFansColumn(f, "K3", "A3", "59fae351ef2d1314e0ea2a5f")) //renren
	saveFans(l)
}

func readFansColumn(f *excelize.File, axis string, timeAxis string, platType string) model.FansExcel {
	var m model.FansExcel
	m.Total = processSum(f.GetCellValue("Sheet1", axis))
	m.Type = platType
	m.CreateTime = getDate(f, "Sheet1", timeAxis).Unix()
	return m
}

func getVideoFromFans(f *excelize.File) {
	var l []model.PlayExcel
	var platSum []model.PlatPlayAmount

	l = append(l, readPlatform(f, 24, "59fae20cef2d1314e0ea2a55", "Sheet2")...) //toutiaoMusic
	l = append(l, readPlatform(f, 49, "59fae276ef2d1314e0ea2a56", "Sheet2")...) //toutiaoSports
	l = append(l, readPlatform(f, 72, "59fae294ef2d1314e0ea2a57", "Sheet2")...) //tMusic
	l = append(l, readPlatform(f, 94, "59fae2a8ef2d1314e0ea2a58", "Sheet2")...) //tSports

	platSum = append(platSum, readPlatformSum(f, "Sheet2", "B24", "59fae20cef2d1314e0ea2a55")) // toutiaoMusic
	platSum = append(platSum, readPlatformSum(f, "Sheet2", "B49", "59fae276ef2d1314e0ea2a56")) // toutiaoSports
	platSum = append(platSum, readPlatformSum(f, "Sheet2", "B72", "59fae294ef2d1314e0ea2a57")) // tMusic
	platSum = append(platSum, readPlatformSum(f, "Sheet2", "B94", "59fae2a8ef2d1314e0ea2a58")) // tSports
	saveVideoData(l, platSum)
}

func saveFans(l []model.FansExcel) {

}
