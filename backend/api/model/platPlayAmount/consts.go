package platPlayAmount

var GetPlatPlayAmount = "SELECT" +
	"  pPA.ids        AS ids," +
	"  pPA.sum        AS sum," +
	"  pT.nameChinese AS platType," +
	"  pPA.createTime AS createTime," +
	"  pPA.grow       AS grow" +
	"  FROM platPlatAmount pPA" +
	"  JOIN platformType pT ON pPA.platType = pT.ids" +
	"  WHERE 1=1 "
var WhereByPlatIds = "  AND pPA.platType = ?"

var GetWeeklyPlayAmount = "SELECT" +
	"  sum(pA.sum) AS sum," +
	"  v.platIds AS pIds," +
	"  pT.nameChinese AS name" +
	"  FROM playAmount pA LEFT JOIN video v ON pA.videoIds = v.ids" +
	"  LEFT JOIN platformType pT ON v.platIds = pT.ids" +
	"  GROUP BY pIds,sum,name"
