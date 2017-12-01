package app

import (
	"github.com/bragfoo/saman/util/db"
	"database/sql"
)

func GetAppData() (*sql.Stmt, error) {
	return db.Prepare(getQuery)
}

func PostAppData() (*sql.Stmt, error) {
	return db.Prepare(postQuery)
}

func PutAppData() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelAppData() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
