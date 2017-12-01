package mobile

var postQuery string = "INSERT INTO saman.mobileData (ids, createTime, active, launch, channel, systemType) VALUES (?, ?, ?, ?, ?, ?);"

var putQuery string = "UPDATE saman.mobileData m" +
	"  SET m.createTime = ?," +
	"  m.active       = ?," +
	"  m.launch       = ?," +
	"  m.systemType   = ?," +
	"  m.channel      = ?" +
	"  WHERE m.ids = ?;"

var delQuery string = "DELETE FROM saman.mobileData WHERE ids = ?"

var GetMobileDataQuery string = "SELECT" +
	"  m.ids        AS ids," +
	"  m.createTime AS createTime," +
	"  m.active     AS active," +
	"  m.launch     AS launch," +
	"  c.name       AS channel," +
	"  m.systemType AS systemType," +
	"  c.ids        AS channelIds" +
	"  FROM saman.mobileData m LEFT JOIN saman.channel c ON m.channel = c.ids" +
	"  WHERE 1=1"

var updateQuery = "UPDATE saman.mobileData m" +
	"  SET m.createTime = ?," +
	"  m.channel      = ?," +
	"  m.systemType   = ?," +
	"  m.launch       = ?," +
	"  m.active       = ?" +
	"  WHERE m.ids = ? "

var GetByChannelQuery = "  AND m.channel = ?"

var GetBySystemType = "  AND m.systemType = ?"

var WherePeriod = "  AND m.createTime > ? AND m.createTime < ?"
