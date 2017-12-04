package task

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
	"fmt"
	"github.com/bragfoo/saman/util"
)

func getStatisticsDaily(query string, offset int) (int64) {
	o := offset - 1
	log.Info(query)
	var t int64
	stm, _ := db.PrepareExt(query, &extDbConf)
	rows, err := stm.Query(getTime(offset), getTime(o))
	defer rows.Close()
	if nil != err {
		log.Error(err)
	} else {
		for rows.Next() {
			rows.Scan(&t)
			fmt.Println(t)
		}
	}

	return t
}

func getStatisticsTotal(query string, offset int) (int64) {
	var t int64
	log.Info(query)
	stm, _ := db.PrepareExt(query, &extDbConf)
	rows, err := stm.Query(getTime(offset))
	defer rows.Close()
	if nil != err {
		log.Error(err)
	} else {
		for rows.Next() {
			rows.Scan(&t)
			fmt.Println(t)
		}
	}

	return t
}

func GetDaily() {
	var model = model.AppUGC{}
	model.Like = getStatisticsDaily(getLikeDailyQuery+GetVestStr(), 1)
	model.ShareSum = getStatisticsDaily(getShareDailyQuery, 1)
	model.CommentSum = getStatisticsDaily(getCommentDailyQuery+GetVestStr(), 1)
	model.VideoSum = getStatisticsDaily(getVideoPlayDaily+GetVestStr(), 1)
	model.VideoStay = getStatisticsDaily(getVideoStayDaily, 1)
	model.PicSum = getStatisticsDaily(getPicDaily+GetVestStr(), 1)
	model.VideoUpload = getStatisticsDaily(getVideoUploadBase+GetVestStr()+getVideoUploadDaily, 1)
	stm, err := db.Prepare(insertDailyUGC)
	if nil != err {
		log.Error(err)
	} else {
		_, e := stm.Exec(
			util.GetObjectId(),
			getTimeInTime(0).Unix(),
			model.Like,
			model.CommentSum,
			model.ShareSum,
			model.PicSum,
			model.VideoSum,
			model.VideoStay,
			model.VideoUpload,
		)

		if nil != e {
			log.Error(e)
		}

	}

}

func GetTotal() {
	var model = model.AppUGC{}
	model.Like = getStatisticsTotal(getLikeTotalQuery+GetVestStr(), 0)
	model.ShareSum = getStatisticsTotal(getShareTotalQuery, 0)
	model.CommentSum = getStatisticsTotal(getCommentTotalQuery+GetVestStr(), 0)
	model.VideoSum = getStatisticsTotal(getVideoPlayTotal+GetVestStr(), 0)
	model.VideoStay = getStatisticsTotal(getVideoStayTotal, 0)
	model.PicSum = getStatisticsTotal(getPicTotal+GetVestStr(), 0)
	model.VideoUpload = getStatisticsTotal(getVideoUploadBase+GetVestStr()+getVideoUpoladTotal, 0)
	stm, err := db.Prepare(insertTotalUGC)
	if nil != err {
		log.Error(err)
	} else {
		_, e := stm.Exec(
			util.GetObjectId(),
			getTimeInTime(0).Unix(),
			model.Like,
			model.CommentSum,
			model.ShareSum,
			model.PicSum,
			model.VideoSum,
			model.VideoStay,
			model.VideoUpload,
		)

		if nil != e {
			log.Error(e)
		}

	}
}

func GetDailyByOffset(offset int) {
	var model = model.AppUGC{}
	o := offset - 1
	model.Like = getStatisticsDaily(getLikeDailyQuery+GetVestStr(), offset)
	model.ShareSum = getStatisticsDaily(getShareDailyQuery, offset)
	model.CommentSum = getStatisticsDaily(getCommentDailyQuery+GetVestStr(), offset)
	model.VideoSum = getStatisticsDaily(getVideoPlayDaily+GetVestStr(), offset)
	model.VideoStay = getStatisticsDaily(getVideoStayDaily, offset)
	model.PicSum = getStatisticsDaily(getPicDaily+GetVestStr(), offset)
	model.VideoUpload = getStatisticsDaily(getVideoUploadBase+GetVestStr()+getVideoUploadDaily, offset)
	stm, err := db.Prepare(insertDailyUGC)
	if nil != err {
		log.Error(err)
	} else {
		_, e := stm.Exec(
			util.GetObjectId(),
			getTimeInTime(o).Unix(),
			model.Like,
			model.CommentSum,
			model.ShareSum,
			model.PicSum,
			model.VideoSum,
			model.VideoStay,
			model.VideoUpload,
		)

		if nil != e {
			log.Error(e)
		}

	}

}

func GetTotalByOffset(offset int) {
	var model = model.AppUGC{}
	model.Like = getStatisticsTotal(getLikeTotalQuery+GetVestStr(), offset)
	model.ShareSum = getStatisticsTotal(getShareTotalQuery, offset)
	model.CommentSum = getStatisticsTotal(getCommentTotalQuery+GetVestStr(), offset)
	model.VideoSum = getStatisticsTotal(getVideoPlayTotal+GetVestStr(), offset)
	model.VideoStay = getStatisticsTotal(getVideoStayTotal, offset)
	model.PicSum = getStatisticsTotal(getPicTotal+GetVestStr(), offset)
	model.VideoUpload = getStatisticsTotal(getVideoUploadBase+GetVestStr()+getVideoUpoladTotal, offset)
	stm, err := db.Prepare(insertTotalUGC)
	if nil != err {
		log.Error(err)
	} else {
		_, e := stm.Exec(
			util.GetObjectId(),
			getTimeInTime(offset).Unix(),
			model.Like,
			model.CommentSum,
			model.ShareSum,
			model.PicSum,
			model.VideoSum,
			model.VideoStay,
			model.VideoUpload,
		)

		if nil != e {
			log.Error(e)
		}

	}
}
