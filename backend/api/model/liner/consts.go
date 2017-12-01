package liner

var GetPlayAmount = "SELECT" +
	"  pPA.ids        AS ids," +
	"  pPA.createTime AS createTime," +
	"  pPA.platType   AS platType," +
	"  pPA.sum        AS sum," +
	"  pPA.grow       AS grow" +
	"  FROM platPlatAmount pPA" +
	"  WHERE 1 = 1"

var GetPlayAmountByPlatIds = "      AND pPA.platType = ? "
