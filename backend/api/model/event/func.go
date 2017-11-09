package event

import (
	"github.com/bragfoo/saman/util/db"
	"database/sql"
	"log"
)

var getEventQuery string = "SELECT" +
	"  e.ids          AS ids," +
	"  e.name         AS name," +
	"  e.startDate    AS startDate," +
	"  e.endDate      AS endDate," +
	"  e.totalPeople  AS totalPeople," +
	"  e.totalWork    AS totalWork," +
	"  e.uploadPeople AS uploadPeople" +
	"  FROM saman.event e"
var getEventByIdQuery string = getEventQuery + " WHERE e.ids = ?"

var updateQuery string = "UPDATE saman.event e" +
	"  SET e.name = ?, e.uploadPeople = ?, e.totalWork = ?, e.totalPeople = ?, e.endDate = ?, e.startDate = ?" +
	"  WHERE e.ids = ?"

var delQuery string = "DELETE FROM saman.event WHERE ids = ?"

func GetEvent() (*sql.Stmt, error) {
	return db.Prepare(getEventQuery)
}

func GetEventById() (*sql.Stmt, error) {
	return db.Prepare(getEventByIdQuery)
}

func PostEvent() (*sql.Stmt, error) {
	var query string
	query = "INSERT INTO saman.event (ids, name, startDate, endDate, totalPeople, totalWork, uploadPeople)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?);"
	stm, err := db.Prepare(query)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func PutEvent() (*sql.Stmt, error) {
	return db.Prepare(updateQuery)
}

func DeleteEvent() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
