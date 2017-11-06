package ugc

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"log"
)

var postugcQuery = "INSERT INTO appUGC (ids, createTime, `like`, commentSum, shareSum, picSum, videoSum) VALUES (?,?,?,?,?,?,?)"

func GetUGC() (*sql.Stmt, error) {
	var query string
	query = "SELECT" +
		"  a.ids        AS ids," +
		"  a.createTime AS createTime," +
		"  a.`like`     AS `like`," +
		"  a.commentSum AS commentSum," +
		"  a.shareSum   AS shareSum," +
		"  a.picSum     AS picSum," +
		"  a.videoSum   AS videoSUm" +
		"  FROM saman.appUGC a"
	stm, err := db.Prepare(query)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func PostUGC() (*sql.Stmt, error) {
	stm, err := db.Prepare(postugcQuery)
	if nil != err {
		return nil, err
	} else {
		return stm, err
	}
}
