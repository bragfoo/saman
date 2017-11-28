package excel

import (
	"mime/multipart"
	"github.com/xuri/excelize"
	"github.com/siddontang/go/log"
)

func OpenMobile(f multipart.File) {
	xlsx, err := excelize.OpenReader(f)
	if nil != err {
		log.Error(err)
	} else {
		getMobileData(xlsx)
	}
}

// read2struct
func getMobileData(f *excelize.File) {
	saveMobileData()
}

// save2db
func saveMobileData() {

}
