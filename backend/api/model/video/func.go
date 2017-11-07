package video

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"log"
)

var getVideoPlayAmount string = "SELECT p.ids AS ids,p.videoIds AS videoIds,p.createTime AS createtime,p.videoIds AS videoIds FROM playAmount p ;"
var postVideoQuery string = "INSERT INTO video (ids, videoIds, platIds, title, link) VALUES (?,?,?,?,?)"
var postVideoPlayAmountQuery string = "INSERT INTO playAmount (ids, videoIds, createTime, sum) VALUES (?,?,?,?);"
var putVideoPlayAmountQuery string = "UPDATE saman.playAmount p" +
	"  SET p.createTime = ?," +
	"  p.sum          = ?," +
	"  p.videoIds     = ?" +
	"  WHERE p.ids = ?;"
var putVideoQuery string = "UPDATE saman.video v" +
	"  SET v.videoIds = ?," +
	"  v.platIds    = ?," +
	"  v.title      = ?," +
	"  v.link       = ?" +
	"  WHERE v.ids = ?;"

var delVideoQuery string = "DELETE FROM saman.video WHERE ids = ?"
var delVideoPlayAmount string = "DELETE FROM saman.playAmount WHERE ids = ?"

func GetVideo() (*sql.Stmt, error) {
	var query string
	query = "SELECT" +
		"  v.ids          AS ids," +
		"  v.title        AS title," +
		"  v.link         AS link," +
		"  pA.createTime  AS createTime," +
		"  pA.sum         AS sum," +
		"  pt.nameChinese AS nameChinese" +
		"  FROM saman.video v LEFT JOIN saman.platformType pt ON v.platIds = pt.ids" +
		"  LEFT JOIN saman.playAmount pA ON pA.videoIds = v.videoIds;"
	stm, err := db.Prepare(query)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetVideoPlayAmount() (*sql.Stmt, error) {
	return db.Prepare(getVideoPlayAmount)
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
