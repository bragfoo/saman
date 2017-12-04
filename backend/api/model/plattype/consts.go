package plat

var postPlatformFansQuery string = "INSERT INTO platformFans (ids, createTime, sum, decrease, increase, platType) VALUES (?,?,?,?,?,?);"

var GetFanQuery = "SELECT" +
	"  pF.ids        AS ids," +
	"  pF.createTime AS createTime," +
	"  pF.sum        AS sum," +
	"  pf.decrease   AS decrease," +
	"  pF.increase   AS increase," +
	"  p.nameChinese AS nameChinese," +
	"  p.ids         AS platIds" +
	"  FROM saman.platformType p LEFT JOIN saman.platformFans pF ON pF.platType = p.ids" +
	"  WHERE 1=1 "

var getFanQueryByPlatIdsQuery = GetFanQuery +
	" AND p.ids = ?"

var putQuery = "UPDATE saman.platformFans p " +
	"  SET p.createTime = ?," +
	"  p.sum          = ?," +
	"  p.increase     = ?," +
	"  p.decrease     = ?," +
	"  p.platType     = ?" +
	"  WHERE p.ids = ?"

var delQuery = "DELETE FROM saman.platformFans WHERE ids = ?"

var WherePeriod = "  AND pF.createTime >? AND pF.createTime < ?"
