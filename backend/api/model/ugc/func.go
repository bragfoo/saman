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
	"  a.videoSum   AS videoSum," +
	"  a.videoStay  AS videoStay" +
	"  FROM saman.appUGC a" +
	"  WHERE 1=1 "
var getByIdsQuery = getQuery + " WHERE a.ids = ? "

var putQuery = "UPDATE saman.appUGC a" +
	"  SET a.createTime = ?," +
	"  a.videoSum     = ?," +
	"  a.picSum       = ?," +
	"  a.shareSum     = ?," +
	"  a.commentSum   = ?," +
	"  a.`like`       = ?," +
	"  a.videoStay    = ?" +
	"  WHERE a.ids = ?"

var getTotalQuery = "SELECT" +
	"  a.ids        AS ids," +
	"  a.createTime AS createTime," +
	"  a.`like`     AS `like`," +
	"  a.commentSum AS commentSum," +
	"  a.shareSum   AS shareSum," +
	"  a.picSum     AS picSum," +
	"  a.videoSum   AS videoSum," +
	"  a.videoStay  AS videoStay" +
	"  FROM saman.appUGCDailyTotal a" +
	"  WHERE 1=1 "

var delQuery = "DELETE FROM appUGC WHERE ids = ?"
var WherePeriod = "  AND a.createTime > ? AND a.createTime < ?"

func GetUGC() (*sql.Stmt, error) {
	stm, err := db.Prepare(getQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetUGCTotal() (*sql.Stmt, error) {
	return db.Prepare(getTotalQuery)
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

func PutUGC() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelUGC() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
