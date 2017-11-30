package excel

var testPlatPlayAmount = "SELECT pPA.ids" +
	"  FROM platPlatAmount pPA" +
	"  WHERE pPA.platType = ? AND pPA.createTime = ?"

var updatePlatPlayAmount = "UPDATE platPlatAmount pPA" +
	"  SET pPA.grow = ?," +
	"  pPA.sum    = ?" +
	"  WHERE pPA.ids = ?;"

var insertPlatPlayAMount = "INSERT INTO platPlatAmount (ids, createTime, platType, sum, grow) VALUES (?, ?, ?, ?, ?);"

var testVideo = "SELECT v.ids" +
	"  FROM video v" +
	"  WHERE v.title = ?"

var insVideo = "INSERT INTO video (ids, platIds, title, link, createTime) VALUES (?,?,?,?,?);"

var testPlayAmount = "SELECT pA.ids" +
	"  FROM playAmount pA" +
	"  WHERE pA.videoIds = ? AND pA.createTime = ?"
var updPlayAmount = "UPDATE playAmount pA" +
	"  SET  pA.sum = ?" +
	"  WHERE pA.ids = ?"
var insPlayAMount = "INSERT INTO playAmount (ids, videoIds, createTime, sum) VALUES (?,?,?,?);"

var testFans = "SELECT pF.ids" +
	"  FROM platformFans pF" +
	"  WHERE pF.createTime = ? AND pF.platType = ?"
var updFans = "UPDATE platformFans pF" +
	"  SET pF.sum = ?" +
	"  WHERE pF.ids = ?"
var insFans = "INSERT INTO platformFans (ids, createTime, sum, decrease, increase, platType) VALUES (?, ?, ?, ?, ?, ?)"
