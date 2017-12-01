package event

var getEventQuery string = "SELECT" +
	"  e.ids          AS ids," +
	"  e.name         AS name," +
	"  e.startDate    AS startDate," +
	"  e.endDate      AS endDate," +
	"  e.totalPeople  AS totalPeople," +
	"  e.totalWork    AS totalWork," +
	"  e.uploadPeople AS uploadPeople" +
	"  FROM saman.event e" +
	"  WHERE 1=1 "
var getEventByIdQuery string = getEventQuery + " WHERE e.ids = ?"

var updateQuery string = "UPDATE saman.event e" +
	"  SET e.name = ?, e.uploadPeople = ?, e.totalWork = ?, e.totalPeople = ?, e.endDate = ?, e.startDate = ?" +
	"  WHERE e.ids = ?"

var delQuery string = "DELETE FROM saman.event WHERE ids = ?"

var WhereStartDate string = "  AND e.startData > ? "
var WhereEndDate string = "  AND e.endDate < ?"
