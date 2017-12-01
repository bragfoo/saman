package event

import (
	"github.com/bragfoo/saman/util/db"
	"database/sql"
	"github.com/siddontang/go/log"
)



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
		log.Error(err)
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
