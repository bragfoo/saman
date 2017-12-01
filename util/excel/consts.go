package excel

const testPlatPlayAmount = "SELECT pPA.ids" +
	"  FROM platPlatAmount pPA" +
	"  WHERE pPA.platType = ? AND pPA.createTime = ?"

const updatePlatPlayAmount = "UPDATE platPlatAmount pPA" +
	"  SET pPA.grow = ?," +
	"  pPA.sum    = ?" +
	"  WHERE pPA.ids = ?;"

const insertPlatPlayAMount = "INSERT INTO platPlatAmount (ids, createTime, platType, sum, grow) VALUES (?, ?, ?, ?, ?);"

const testVideo = "SELECT v.ids" +
	"  FROM video v" +
	"  WHERE v.title = ?"

const insVideo = "INSERT INTO video (ids, platIds, title, link, createTime) VALUES (?,?,?,?,?);"

const testPlayAmount = "SELECT pA.ids" +
	"  FROM playAmount pA" +
	"  WHERE pA.videoIds = ? AND pA.createTime = ?"
const updPlayAmount = "UPDATE playAmount pA" +
	"  SET  pA.sum = ?" +
	"  WHERE pA.ids = ?"
const insPlayAMount = "INSERT INTO playAmount (ids, videoIds, createTime, sum) VALUES (?,?,?,?);"

const testFans = "SELECT pF.ids" +
	"  FROM platformFans pF" +
	"  WHERE pF.createTime = ? AND pF.platType = ?"
const updFans = "UPDATE platformFans pF" +
	"  SET pF.sum = ?" +
	"  WHERE pF.ids = ?"
const insFans = "INSERT INTO platformFans (ids, createTime, sum, decrease, increase, platType) VALUES (?, ?, ?, ?, ?, ?)"
