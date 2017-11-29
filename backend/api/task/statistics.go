package task

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
)

func getStatisticsDaily(query string, offset int) (int) {
	var ret = 0
	stm, err := db.PrepareExt(query, &extDbConf)
	if nil != err {
		log.Error(err)
	} else {
		rows, err := stm.Query(getTime(offset), getTime(offset-1))
		if nil != err {
			log.Error(err)
		} else {
			rows.Scan(&ret)
		}
	}
	return ret
}

func getStatisticsTotal(query string) (int) {
	var ret = 0
	stm, err := db.PrepareExt(query, &extDbConf)
	if nil != err {
		log.Error(err)
	} else {
		rows, err := stm.Query(getTime(0))
		if nil != err {
			log.Error(err)
		} else {
			rows.Scan(&ret)
		}
	}
	return ret
}
