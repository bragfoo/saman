package video

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"log"
)

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
