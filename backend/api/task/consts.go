package task

import "github.com/bragfoo/saman/util/config"

var extDbConf = config.DbType{
	Username: "instreet",
	Password: "instreetCoding@1by1",
	Address:  "rm-2ze33sny9l33w8pp8.mysql.rds.aliyuncs.com",
	Dbname:   "instreet4",
	Protocol: "tcp",
	Port:     "3306",
}

//用户点赞
var getLikeDailyQuery = "SELECT count(1) AS `likeCount`" +
	"  FROM instreet_item_user_like WHERE like_time > ? AND like_time < ?"

var getLikeTotalQuery = "SELECT count(1) AS `likeCount`" +
	"  FROM instreet_item_user_like WHERE AND like_time < ?"

//用户分享
var getShareDailyQuery = "SELECT count(1) AS videoCount" +
	"  FROM instreet_actions ia" +
	"  WHERE ia.action_type = 2 AND action_time > ? AND action_time < ?"

var getShareTotalQuery = "SELECT count(1) AS videoCount" +
	"  FROM instreet_actions ia" +
	"  WHERE ia.action_type = 2 AND action_time < ?"

//用户评论
var getCommentDailyQuery = "SELECT count(1) AS commentCount" +
	"  FROM instreet_comments ic" +
	"  WHERE ic.created_at > ? AND ic.created_at < ?"

var getCommentTotalQuery = "SELECT count(1) AS commentCount" +
	"  FROM instreet_comments ic" +
	"  WHERE ic.created_at < ?"

//视频点击
var getVideoPlayDaily = "SELECT sum(ii.click) AS videoPlaySum" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 1 AND ii.created_at > ? AND ii.created_at < ?"

var getVideoPlatTotal = "SELECT sum(ii.click) AS videoPlaySum" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 1 ii.created_at < ?"

//视频播放时长
var getVideoStayTotal = "SELECT sum(ia.action_stay) AS videoStay" +
	"  FROM instreet_actions ia WHERE ia.action_type = 6"

var getVideoStayDaily = "SELECT sum(ia.action_stay) AS videoStay" +
	"  FROM instreet_actions ia" +
	"  WHERE ia.action_type = 6 AND ia.action_time > ? AND ia.action_time < ?"
