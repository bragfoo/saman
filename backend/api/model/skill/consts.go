package skill

var GetQuery string = "SELECT" +
	"  s.ids  AS ids," +
	"  s.name AS `name`" +
	"  FROM saman.skill s" +
	"  WHERE 1=1 "

var WhereIds = "  AND s.ids = ?"
var WhereName = "  AND s.name = ?"
var putQuery string = "UPDATE saman.skill s" +
	"  SET s.name = ?" +
	"  WHERE s.ids = ?"
var postQuery string = "INSERT INTO saman.skill (ids, name) VALUES (?,?)"
var delQuery string = "DELETE FROM saman.skill WHERE ids = ?"
