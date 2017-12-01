package video

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
)

func GetVideoPlayAmount() (*sql.Stmt, error) {
	return db.Prepare(GetPlayAmountQuery)
}

func GetVideo() (*sql.Stmt, error) {
	return db.Prepare(GetVideoQuery)
}

func GetVideoSource() (*sql.Stmt, error) {
	return db.Prepare(getVideoSourceQuery)
}

func PostVideo() (*sql.Stmt, error) {
	stm, err := db.Prepare(postVideoQuery)
	if nil != err {
		return nil, err
	} else {
		return stm, nil
	}
}

func PostVideoPlayAmount() (*sql.Stmt, error) {
	stm, err := db.Prepare(postVideoPlayAmountQuery)
	if nil != err {
		return nil, err
	} else {
		return stm, nil
	}
}

func PutVideo() (*sql.Stmt, error) {
	return db.Prepare(putVideoQuery)
}

func PutVideoPlayAmount() (*sql.Stmt, error) {
	return db.Prepare(putVideoPlayAmountQuery)
}
func DelVideo() (*sql.Stmt, error) {
	return db.Prepare(delVideoQuery)
}

func DelVideoPlayAmount() (*sql.Stmt, error) {
	return db.Prepare(delVideoPlayAmount)
}
