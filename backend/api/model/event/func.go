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

var updateQuery string = ""

func GetEvent() (*sql.Stmt, error) {
	stm, err :=
		db.Prepare(getEventQuery)
	if nil != err {
		log.Fatal(err)
		return nil, err
	} else {
		return stm, nil
	}
}

func GetEventById() (*sql.Stmt, error) {
	stm, err := db.Prepare(getEventByIdQuery)
	if nil != err {
		return nil, err
	} else {
		return stm, nil
	}
}

func PostEvents() (*sql.Stmt, error) {
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
