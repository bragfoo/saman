package talent

var GetQuery = "SELECT" +
	"  t.ids  AS ids," +
	"  t.type AS `type`," +
	"  t.user AS user," +
	"  s.name AS skillName," +
	"  t.createTime AS createTime," +
	"  t.submitted AS submitted" +
	"  FROM saman.talent t " +
	"  LEFT JOIN saman.skill s ON t.type = s.ids " +
	"  WHERE 1=1 "
var postQuery = "INSERT INTO saman.talent (ids, user, type, createTime, submitted) VALUES (?, ?, ?, ?, ?)"
var putQuery = "UPDATE saman.talent t" +
	"  SET t.type = ?," +
	"  t.user   = ?," +
	"  t.createTime  = ?," +
	"  t.submitted = ?" +
	"  WHERE t.ids = ?"
var delQuery = "DELETE FROM saman.talent WHERE ids = ?"

//con
var WhereIds = "  AND t.ids = ? "
var WhereUser = "  AND t.user = ? "
var WhereType = "  AND t.type = ? "
var WhereName = "  AND t.name = ? "
var WhereSubmitted = "  AND t.submitted = ? "
var WherePeriod = "  AND t.createTime > ? AND t.createTime < ?"
