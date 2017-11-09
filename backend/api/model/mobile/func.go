package mobile

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
	"log"
)

var postEventsQuery string = "INSERT INTO saman.mobileData (ids, createTime, active, launch, channel, systemType) VALUES (?, ?, ?, ?, ?, ?);"

var putQuery string = "UPDATE saman.mobileData m" +
	"  SET m.createTime = ?," +
	"  m.active       = ?," +
	"  m.launch       = ?," +
	"  m.systemType   = ?," +
	"  m.channel      = ?" +
	"  WHERE m.ids = ?;"

var delQuery string = "DELETE FROM saman.mobileData WHERE ids = ?"

var getMobileDataQuery string = "SELECT" +
	"  m.ids        AS ids," +
	"  m.createTime AS createTime," +
	"  m.active     AS active," +
	"  m.launch     AS launch," +
	"  c.name       AS channel," +
	"  m.systemType AS systemType," +
	"  c.ids        AS channelIds" +
	"  FROM saman.mobileData m LEFT JOIN saman.channel c ON m.channel = c.ids"

var updateQuery = "UPDATE saman.mobileData m" +
	"  SET m.createTime = ?," +
	"  m.channel      = ?," +
	"  m.systemType   = ?," +
	"  m.launch       = ?," +
	"  m.active       = ?" +
	"  WHERE m.ids = ? "

var getByChannelQuery = getMobileDataQuery + "  WHERE m.channel = ?"

func GetMobileData() (*sql.Stmt, error) {

	stm, err := db.Prepare(getMobileDataQuery)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetMobileDataByChannel() (*sql.Stmt, error) {
	return db.Prepare(getByChannelQuery)
}

func PutMobileData() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelMObileData() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
