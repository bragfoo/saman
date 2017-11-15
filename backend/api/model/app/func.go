package app

import (
	"github.com/bragfoo/saman/util/db"
	"database/sql"
)

var getQuery string = "SELECT" +
	"  a.ids         AS ids," +
	"  a.createTime  AS createTime," +
	"  a.activeUser  AS activeUser," +
	"  a.picUpload   AS picUpload," +
	"  a.talentSum   AS talentSum," +
	"  a.videoUpload AS videoUpload" +
	"  FROM saman.appData a " +
	"  WHERE 1=1 "
var postQuery string = "INSERT INTO saman.appData (ids, createTime, picUpload, videoUpload, talentSum, activeUser) VALUES (?,?,?,?,?,?)"
var putQuery string = "UPDATE saman.appData a" +
	"  SET a.createTime = ?," +
	"  a.activeUser   = ?," +
	"  a.picUpload    = ?," +
	"  a.talentSum    = ?," +
	"  a.videoUpload  = ?" +
	"  WHERE a.ids = ?"
var delQuery string = "DELETE FROM saman.appData WHERE ids = ?"

var WherePeriod string = "AND a.createTime > ? AND a.createTime < ?"

func GetAppData() (*sql.Stmt, error) {
	return db.Prepare(getQuery)
}

func PostAppData() (*sql.Stmt, error) {
	return db.Prepare(postQuery)
}

func PutAppData() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelAppData() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
