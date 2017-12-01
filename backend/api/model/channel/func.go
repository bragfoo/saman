package channel

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
)



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
