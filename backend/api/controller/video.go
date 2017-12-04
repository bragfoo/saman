package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/video"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/util"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/util/db"
)

func GetVideo(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {

		platIds := context.Query("platIds")

		var con []interface{}
		sql := video.GetVideoQuery
		if "" != platIds {
			sql += video.VideoWherePlatIds
			con = append(con, platIds)
		}
		//s, c := common.GetTimePeriodByoffset(sql, con, "v",9)
		stm, err := db.Prepare(sql)
		defer stm.Close()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			rows, err := stm.Query(con...)
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				var result []model.Video
				for rows.Next() {
					var model = model.Video{}
					err := rows.Scan(&model.Ids,
						&model.Title,
						&model.Link,
						&model.CreateTime,
						&model.PlatIds,
						&model.VideoIds,
					)
					if nil == err {
						result = append(result, model)
					}
				}
				context.JSON(http.StatusOK, result)
			}
		}
	}
}

func GetVideoSource(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := video.GetVideoSource()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.VideoPlayAmount
			for rows.Next() {
				var model = model.VideoPlayAmount{}
				err := rows.Scan(&model.Ids,
					&model.Title,
					&model.Link,
					&model.CreateTime,
					&model.PlatIds,
				)
				if nil == err {
					result = append(result, model)
				}
			}
			context.JSON(http.StatusOK, result)
		}
	}
}

func PostVideo(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.Video{}
		common.ReadJSON(c, &m)
		stm, err := video.PostVideo()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.VideoIds,
				m.PlatIds,
				m.Title,
				m.Link,
				m.CreateTime,
			)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}

func PutVideo(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.Video{}
		common.ReadJSON(context, &m)
		stm, err := video.PutVideo()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec(m.PlatIds, m.Title, m.Link, m.CreateTime, m.PlatIds, m.Ids)
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func DelVideo(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := video.DelVideo()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				_, err := stm.Exec(ids)
				if nil != err {
					log.Error(err)
					common.StandardError(c)
				} else {
					common.StandardSuccess(c)
				}
			}
		}
	}
}
