package skill

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
)

func GetSkill() (*sql.Stmt, error) {
	return db.Prepare(GetQuery)
}

func PutSkill() (*sql.Stmt, error) {
	return db.Prepare(putQuery)
}

func PostSkill() (*sql.Stmt, error) {
	return db.Prepare(postQuery)
}

func DelSkill() (*sql.Stmt, error) {
	return db.Prepare(delQuery)
}
