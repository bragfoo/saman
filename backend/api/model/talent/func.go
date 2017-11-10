package talent

import (
	"database/sql"
	"github.com/bragfoo/saman/util/db"
)

var GetQuery = "SELECT" +
	"  t.ids  AS ids," +
	"  t.name AS `name`," +
	"  t.type AS `type`," +
	"  t.user AS user," +
	"  s.name AS skillName" +
	"  FROM saman.talent t " +
	"  LEFT JOIN saman.skill s ON t.type = s.ids " +
	"  WHERE 1=1 "
var postQuery = "INSERT INTO saman.talent (ids, user, type, name) VALUES (?, ?, ?, ?)"
var putQuery = "UPDATE saman.talent t" +
	"  SET t.type = ?," +
	"  t.user   = ?," +
	"  t.name   = ?" +
	"  WHERE t.ids = ?"
var delQuery = "DELETE FROM saman.talent WHERE ids = ?"

//con
var WhereIds = "  AND t.ids = ? "
var WhereUser = "  AND t.user = ? "
var WhereType = "  AND t.type = ? "
var WhereName = "  AND t.name = ? "

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
