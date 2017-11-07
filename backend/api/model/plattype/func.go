package plat

import (
	"github.com/bragfoo/saman/util/db"
	"database/sql"
	"github.com/siddontang/go/log"
)

var postPlatformFansQuery string = "INSERT INTO platformFans (ids, createTime, sum, decrease, increase, platType) VALUES (?,?,?,?,?,?);"

var getFanQuery = "SELECT" +
	"  pF.ids        AS ids," +
	"  pF.createTime AS createTime," +
	"  pF.sum        AS sum," +
	"  pf.decrease   AS decrease," +
	"  pF.increase   AS increase," +
	"  p.nameChinese AS nameChinese," +
	"  p.ids         AS platIds" +
	"  FROM saman.platformType p LEFT JOIN saman.platformFans pF ON pF.platType = p.ids"

var getFanQueryByPlatIdsQuery = getFanQuery +
	" WHERE p.ids = ?"

var putQuery = "UPDATE saman.platformFans p " +
	"  SET p.createTime = ?," +
	"  p.sum          = ?," +
	"  p.increase     = ?," +
	"  p.decrease     = ?," +
	"  p.platType     = ?" +
	"  WHERE p.ids = ?"

var delQuery = "DELETE FROM saman.platformFans WHERE ids = ?"

func GetPlatformTypeStmt() (*sql.Stmt, error) {
	stm, err := db.Prepare("SELECT p.ids AS ids,p.name AS name,p.nameChinese AS nameChinese FROM saman.platformType p")
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetPlatformFansStmt() (*sql.Stmt, error) {
	stm, err := db.Prepare(getFanQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}
func GetPlatformFansByPlatIds() (*sql.Stmt, error) {
	stm, err := db.Prepare(getFanQueryByPlatIdsQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func PostPlatformFans() (*sql.Stmt, error) {
	stm, err := db.Prepare(postPlatformFansQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func PutPlatformFans() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelPlatformFans() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
