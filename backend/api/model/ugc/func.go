package ugc

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
)

var postugcQuery = "INSERT INTO appUGC (ids, createTime, `like`, commentSum, shareSum, picSum, videoSum) VALUES (?,?,?,?,?,?,?)"
var getQuery = "SELECT" +
	"  a.ids        AS ids," +
	"  a.createTime AS createTime," +
	"  a.`like`     AS `like`," +
	"  a.commentSum AS commentSum," +
	"  a.shareSum   AS shareSum," +
	"  a.picSum     AS picSum," +
	"  a.videoSum   AS videoSUm" +
	"  FROM saman.appUGC a"
var getByIdsQuery = getQuery + " WHERE a.ids = ? "

func GetUGC() (*sql.Stmt, error) {
	stm, err := db.Prepare(getQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetUGCByIds() (*sql.Stmt, error) {
	stm, err := db.Prepare(getByIdsQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func PostUGC() (*sql.Stmt, error) {
	stm, err := db.Prepare(postugcQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, err
	}
}
