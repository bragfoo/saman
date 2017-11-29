package task

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
	"fmt"
	"github.com/bragfoo/saman/util"
)

func getStatisticsDaily(query string, offset int) (int64) {
	log.Info(query)
	var t int64
	stm, _ := db.PrepareExt(query, &extDbConf)
	rows, err := stm.Query(getTime(1), getTime(0))
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

func getStatisticsTotal(query string) (int64) {
	var t int64
	log.Info(query)
	stm, _ := db.PrepareExt(query, &extDbConf)
	rows, err := stm.Query(getTime(0))
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
	model.Like = getStatisticsTotal(getLikeTotalQuery)
	model.ShareSum = getStatisticsTotal(getShareTotalQuery)
	model.CommentSum = getStatisticsTotal(getCommentTotalQuery)
	model.VideoSum = getStatisticsTotal(getVideoPlayTotal)
	model.VideoStay = getStatisticsTotal(getVideoStayTotal)
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
