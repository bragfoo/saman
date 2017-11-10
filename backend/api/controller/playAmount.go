package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/bragfoo/saman/backend/api/model/video"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/util/db"
)

func GetVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {

		videoIds := context.Query("videoIds")
		platIds := context.Query("platIds")

		var con []interface{}
		sql := video.GetPlayAmountQuery
		if "" != videoIds {
			sql += video.WhereVideoIds
			con = append(con, videoIds)
		}
		if "" != platIds {
			sql += video.WherePlatIds
			con = append(con, platIds)
		}
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
				var result []model.VideoPlayAmount
				for rows.Next() {
					var model = model.VideoPlayAmount{}
					err := rows.Scan(&model.Ids,
						&model.Title,
						&model.Link,
						&model.CreateTime,
						&model.Sum,
						&model.NameChinese,
						&model.CreateTime,
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
func PostVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.PlayAmount{}
		common.ReadJSON(c, &m)
		stm, err := video.PostVideoPlayAmount()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.VideoIds,
				m.CreateTime,
				m.Sum)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}

func PutVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.VideoPlayAmount{}
		common.ReadJSON(context, &m)
		stm, err := video.PutVideoPlayAmount()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec()
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func DelVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := video.DelVideoPlayAmount()
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
