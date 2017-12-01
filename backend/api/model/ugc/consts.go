package ugc

var postugcQuery = "INSERT INTO appUGC (ids, createTime, `like`, commentSum, shareSum, picSum, videoSum) VALUES (?,?,?,?,?,?,?)"
var getQuery = "SELECT" +
	"  a.ids        AS ids," +
	"  a.createTime AS createTime," +
	"  a.`like`     AS `like`," +
	"  a.commentSum AS commentSum," +
	"  a.shareSum   AS shareSum," +
	"  a.picSum     AS picSum," +
	"  a.videoSum   AS videoSum," +
	"  a.videoStay  AS videoStay" +
	"  FROM saman.appUGC a" +
	"  WHERE 1=1 "
var getByIdsQuery = getQuery + " WHERE a.ids = ? "

var putQuery = "UPDATE saman.appUGC a" +
	"  SET a.createTime = ?," +
	"  a.videoSum     = ?," +
	"  a.picSum       = ?," +
	"  a.shareSum     = ?," +
	"  a.commentSum   = ?," +
	"  a.`like`       = ?," +
	"  a.videoStay    = ?" +
	"  WHERE a.ids = ?"

var getTotalQuery = "SELECT" +
	"  a.ids        AS ids," +
	"  a.createTime AS createTime," +
	"  a.`like`     AS `like`," +
	"  a.commentSum AS commentSum," +
	"  a.shareSum   AS shareSum," +
	"  a.picSum     AS picSum," +
	"  a.videoSum   AS videoSum," +
	"  a.videoStay  AS videoStay" +
	"  FROM saman.appUGCDailyTotal a" +
	"  WHERE 1=1 "

var delQuery = "DELETE FROM appUGC WHERE ids = ?"
var WherePeriod = "  AND a.createTime > ? AND a.createTime < ?"
