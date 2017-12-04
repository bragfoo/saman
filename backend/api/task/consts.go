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
	"  FROM instreet_item_user_like WHERE like_time > ? AND like_time < ?  AND instreet_item_user_like.user_id NOT IN "

var getLikeTotalQuery = "SELECT count(1) AS `likeCount`" +
	"  FROM instreet_item_user_like WHERE like_time < ?  AND instreet_item_user_like.user_id NOT IN  "

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
	"  WHERE ic.created_at > ? AND ic.created_at < ?  AND ic.user_id NOT IN "

var getCommentTotalQuery = "SELECT count(1) AS commentCount" +
	"  FROM instreet_comments ic" +
	"  WHERE ic.created_at < ?  AND ic.user_id NOT IN "

//视频点击
var getVideoPlayDaily = "SELECT sum(ii.click) AS videoPlaySum" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 1 AND ii.created_at > ? AND ii.created_at < ?  AND ii.user_id NOT IN "

var getVideoPlayTotal = "SELECT sum(ii.click) AS videoPlaySum" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 1 AND ii.created_at < ?  AND ii.user_id NOT IN "

//视频播放时长
var getVideoStayTotal = "SELECT sum(ia.action_stay) AS videoStay" +
	"  FROM instreet_actions ia WHERE ia.action_type = 6 AND ia.action_time < ?"

var getVideoStayDaily = "SELECT sum(ia.action_stay) AS videoStay" +
	"  FROM instreet_actions ia" +
	"  WHERE ia.action_type = 6 AND ia.action_time > ? AND ia.action_time < ?"

//图片
var getPicDaily = "SELECT count(1) AS picCount" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 2" +
	"  AND ii.published_at > ? AND ii.published_at < ?  AND ii.user_id NOT IN "

var getPicTotal = "SELECT count(1) AS picCount" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 2 AND" +
	"  ii.published_at < ?  AND ii.user_id NOT IN "

var getVideoUploadBase = "SELECT count(1)" +
	"  FROM instreet_items ii" +
	"  WHERE ii.itemtype_id = 4 AND ii.user_id NOT IN "
var getVideoUploadDaily = "      AND ii.created_at > ? AND ii.created_at < ?"
var getVideoUpoladTotal = " AND ii.created_at < ?"

//插入语句
var insertDailyUGC = "INSERT INTO appUGC (ids, createTime, `like`, commentSum, shareSum, picSum, videoSum, videoStay, videoUpload)" +
	"  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

var insertTotalUGC = "INSERT INTO appUGCDailyTotal (ids, createTime, `like`, commentSum, shareSum, picSum, videoSum, videoStay, videoUpload)" +
	"  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
