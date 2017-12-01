package plat

import (
	"github.com/bragfoo/saman/util/db"
	"database/sql"
	"github.com/siddontang/go/log"
)

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
