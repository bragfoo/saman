package talent

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
)

func GetTalent() (*sql.Stmt, error) {
	return db.Prepare(GetQuery)
}

func PostTalnet() (*sql.Stmt, error) {
	return db.Prepare(postQuery)
}

func PutQuery() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func DelQuery() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
