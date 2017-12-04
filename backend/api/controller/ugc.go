package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/ugc"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/util/db"
)

func GetUGC(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		ids := context.Query("ids")
		if "" != ids {
			stm, err := ugc.GetUGCByIds()
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				var model = model.AppUGC{}
				err := stm.QueryRow(ids).Scan(&model.Ids,
					&model.CreateTime,
					&model.Like,
					&model.CommentSum,
					&model.ShareSum,
					&model.PicSum,
					&model.VideoSum,
					&model.VideoSum,
					&model.VideoUpload,
				)
				if nil != err {
					log.Error(err)
					common.StandardNotFound(context)
				} else {
					common.StandardOk(context, model)
				}
			}
		} else {
			sql := ugc.GetQuery
			var con []interface{}
			s, c := common.GetTimePeriod(sql, con, "a")
			stm, err := db.Prepare(s)
			if nil != err {
				context.Status(http.StatusInternalServerError)
			} else {
				rows, _ := stm.Query(c...)
				defer rows.Close()
				var result []model.AppUGC
				for rows.Next() {
					var model = model.AppUGC{}
					err := rows.Scan(&model.Ids,
						&model.CreateTime,
						&model.Like,
						&model.CommentSum,
						&model.ShareSum,
						&model.PicSum,
						&model.VideoSum,
						&model.VideoStay,
						&model.VideoUpload,
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

func PostUGC(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.AppUGC{}
		common.ReadJSON(c, &m)
		stm, err := ugc.PostUGC()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.CreateTime,
				m.Like,
				m.CommentSum,
				m.ShareSum,
				m.PicSum,
				m.VideoSum,
				m.VideoStay,
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

func PutUGC(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.AppUGC{}
		common.ReadJSON(context, &m)
		stm, err := ugc.PutUGC()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec(m.CreateTime,
				m.Like,
				m.CommentSum,
				m.ShareSum,
				m.PicSum,
				m.VideoSum,
				m.VideoStay,
			)
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func DelUGC(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := ugc.DelUGC()
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
