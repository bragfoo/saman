package mobile

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"log"
)

var postEventsQuery string = "INSERT INTO saman.mobileData (ids, createTime, active, launch, channel, systemType) VALUES (?, ?, ?, ?, ?, ?);"

var getMobileDataQuery string = "SELECT" +
	"  m.ids        AS ids," +
	"  m.createTime AS createTime," +
	"  m.active     AS active," +
	"  m.launch     AS launch," +
	"  c.name       AS channel," +
	"  m.systemType AS systemType" +
	"  FROM saman.mobileData m LEFT JOIN saman.channel c ON m.channel = c.ids"

func GetChannel() (*sql.Stmt, error) {
	var query string
	query = "SELECT c.ids AS ids,c.name AS name FROM saman.channel c"
	stm, err := db.Prepare(query)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetMobileData() (*sql.Stmt, error) {

	stm, err := db.Prepare(getMobileDataQuery)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func PostEvents() (*sql.Stmt, error) {
	stm, err := db.Prepare(postEventsQuery)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}
