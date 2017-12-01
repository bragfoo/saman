package app

var getQuery string = "SELECT" +
	"  a.ids         AS ids," +
	"  a.createTime  AS createTime," +
	"  a.activeUser  AS activeUser," +
	"  a.picUpload   AS picUpload," +
	"  a.talentSum   AS talentSum," +
	"  a.videoUpload AS videoUpload" +
	"  FROM saman.appData a " +
	"  WHERE 1=1 "
var postQuery string = "INSERT INTO saman.appData (ids, createTime, picUpload, videoUpload, talentSum, activeUser) VALUES (?,?,?,?,?,?)"
var putQuery string = "UPDATE saman.appData a" +
	"  SET a.createTime = ?," +
	"  a.activeUser   = ?," +
	"  a.picUpload    = ?," +
	"  a.talentSum    = ?," +
	"  a.videoUpload  = ?" +
	"  WHERE a.ids = ?"
var delQuery string = "DELETE FROM saman.appData WHERE ids = ?"

var WherePeriod string = "AND a.createTime > ? AND a.createTime < ?"
