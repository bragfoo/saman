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
	model.Like = getStatisticsDaily(getLikeDailyQuery, 1)
	model.ShareSum = getStatisticsDaily(getShareDailyQuery, 1)
	model.CommentSum = getStatisticsDaily(getCommentDailyQuery, 1)
	model.VideoSum = getStatisticsDaily(getVideoPlayDaily, 1)
	model.VideoStay = getStatisticsDaily(getVideoStayDaily, 1)
	model.PicSum = getStatisticsDaily(getPicTotal, 1)
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
		)

		if nil != e {
			log.Error(e)
		}

	}

}

func GetTotal() {
	var model = model.AppUGC{}
	model.Like = getStatisticsTotal(getLikeTotalQuery, 0)
	model.ShareSum = getStatisticsTotal(getShareTotalQuery, 0)
	model.CommentSum = getStatisticsTotal(getCommentTotalQuery, 0)
	model.VideoSum = getStatisticsTotal(getVideoPlayTotal, 0)
	model.VideoStay = getStatisticsTotal(getVideoStayTotal, 0)
	model.PicSum = getStatisticsDaily(getPicTotal, 0)
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
		)

		if nil != e {
			log.Error(e)
		}

	}
}

func GetDailyByOffset(offset int) {
	var model = model.AppUGC{}
	o := offset - 1
	model.Like = getStatisticsDaily(getLikeDailyQuery, offset)
	model.ShareSum = getStatisticsDaily(getShareDailyQuery, offset)
	model.CommentSum = getStatisticsDaily(getCommentDailyQuery, offset)
	model.VideoSum = getStatisticsDaily(getVideoPlayDaily, offset)
	model.VideoStay = getStatisticsDaily(getVideoStayDaily, offset)
	model.PicSum = getStatisticsDaily(getPicDaily, offset)
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
		)

		if nil != e {
			log.Error(e)
		}

	}

}

func GetTotalByOffset(offset int) {
	var model = model.AppUGC{}
	model.Like = getStatisticsTotal(getLikeTotalQuery, offset)
	model.ShareSum = getStatisticsTotal(getShareTotalQuery, offset)
	model.CommentSum = getStatisticsTotal(getCommentTotalQuery, offset)
	model.VideoSum = getStatisticsTotal(getVideoPlayTotal, offset)
	model.VideoStay = getStatisticsTotal(getVideoStayTotal, offset)
	model.PicSum = getStatisticsTotal(getPicTotal, offset)
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
		)

		if nil != e {
			log.Error(e)
		}

	}
}
