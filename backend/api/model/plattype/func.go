package plat

import (
	"github.com/bragfoo/saman/util/db"
	"log"
	"database/sql"
)

func GetPlatformTypeStmt() (*sql.Stmt, error) {
	stm, err := db.Prepare("SELECT p.ids AS ids,p.name AS name,p.nameChinese AS nameChinese FROM saman.platformType p")
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetPlatformFansStmt() (*sql.Stmt, error) {
	var query string
	query = "SELECT" +
		"  pF.ids        AS ids," +
		"  pF.createTime AS createTime," +
		"  pF.sum        AS sum," +
		"  pf.decrease   AS decrease," +
		"  pF.increase   AS increase," +
		"  p.nameChinese AS nameChinese" +
		"  FROM saman.platformType p LEFT JOIN saman.platformFans pF ON pF.type = p.ids;"
	stm, err := db.Prepare(query)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}
