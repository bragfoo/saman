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
