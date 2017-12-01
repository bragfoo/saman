package channel

var getQuery string = "SELECT c.ids AS ids,c.name AS name FROM saman.channel c"
var postQuery string = "INSERT INTO saman.channel (ids, name) VALUES (?, ?);"
var putQuery string = "UPDATE saman.channel c SET c.name = ? WHERE c.ids = ?;"
var delQuery string = ""
