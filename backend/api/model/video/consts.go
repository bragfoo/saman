package video

var getVideoPlayAmount string = "SELECT p.ids AS ids,p.videoIds AS videoIds,p.createTime AS createtime,p.videoIds AS videoIds FROM playAmount p ;"
var postVideoQuery string = "INSERT INTO video (ids, videoIds, platIds, title, link, createTime) VALUES (?,?,?,?,?,?)"
var postVideoPlayAmountQuery string = "INSERT INTO playAmount (ids, videoIds, createTime, sum) VALUES (?,?,?,?);"
var putVideoPlayAmountQuery string = "UPDATE saman.playAmount p" +
	"  SET p.createTime = ?," +
	"  p.sum          = ?," +
	"  p.videoIds     = ?" +
	"  WHERE p.ids = ?;"
var putVideoQuery string = "UPDATE saman.video v" +
	"  SET " +
	"  v.platIds    = ?," +
	"  v.title      = ?," +
	"  v.link       = ?," +
	"  v.createTime = ?," +
	"  v.platIds    = ?" +
	"  WHERE v.ids = ?;"

var delVideoQuery string = "DELETE FROM saman.video WHERE ids = ?"
var delVideoPlayAmount string = "DELETE FROM saman.playAmount WHERE ids = ?"

var GetVideoQuery string = "SELECT" +
	"  v.ids          AS ids," +
	"  v.title        AS title," +
	"  v.link         AS link," +
	"  v.createTime   AS createTime," +
	"  v.platIds      AS platIds," +
	"  v.videoIds     AS videoIds" +
	"  FROM saman.video v" +
	"  WHERE 1=1 "

var VideoWherePlatIds = "  AND v.platIds = ? "

var getVideoSourceQuery string = "SELECT" +
	"  v.ids          AS ids," +
	"  v.title        AS title," +
	"  v.link         AS link," +
	"  v.createTime   AS createTime," +
	"  v.platIds      AS platIds" +
	"  FROM saman.video v" +
	"  WHERE v.videoIds = ''"

var GetPlayAmountQuery = "SELECT" +
	"  pA.ids          AS ids," +
	"  v.title        AS title," +
	"  v.link         AS link," +
	"  pA.createTime  AS createTime," +
	"  pA.sum         AS sum," +
	"  pt.nameChinese AS nameChinese," +
	"  pA.createTime   AS createTime," +
	"  v.platIds      AS platIds," +
	"  v.ids          AS videoIds" +
	"  FROM saman.video v LEFT JOIN saman.platformType pt ON v.platIds = pt.ids" +
	"  LEFT JOIN saman.playAmount pA ON pA.videoIds = v.ids " +
	"  WHERE 1=1 "
var WhereVideoIds = "  AND pA.videoIds = ?"

var WherePlatIds = "  AND v.platIds = ?"
