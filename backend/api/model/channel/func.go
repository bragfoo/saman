package channel

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
)

var getQuery string = "SELECT c.ids AS ids,c.name AS name FROM saman.channel c"
var postQuery string = "INSERT INTO saman.channel (ids, name) VALUES (?, ?);"
var putQuery string = "UPDATE saman.channel c SET c.name = ? WHERE c.ids = ?;"
var delQuery string = ""

func GetChannel() (*sql.Stmt, error) {
	return db.Prepare(getQuery)
}

func PostChannel() (*sql.Stmt, error) {
	return db.Prepare(postQuery)
}

func PutChannel() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelChannel() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
