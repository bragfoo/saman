package mobile

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
)

func GetMobileData() (*sql.Stmt, error) {

	stm, err := db.Prepare(GetMobileDataQuery)
	if nil != err {
		log.Error(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetMobileDataByChannel() (*sql.Stmt, error) {
	return db.Prepare(GetByChannelQuery)
}

func PutMobileData() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelMObileData() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}

func PostMobileData() (*sql.Stmt, error) {
	return db.Prepare(postQuery)
}
