package ugc

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
)

func GetUGC() (*sql.Stmt, error) {
	stm, err := db.Prepare(GetQuery)
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
	stm, err := db.Prepare(GetByIdsQuery)
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
