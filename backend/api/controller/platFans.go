package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/bragfoo/saman/backend/api/model/plattype"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"net/http"
	"github.com/bragfoo/saman/util"
	"github.com/bragfoo/saman/util/db"
)

func GetPlatFans(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		platIds := context.Query("platIds")
		if "" != platIds {
			sql := plat.GetFanQueryByPlatIdsQuery
			var con []interface{}
			con = append(con, platIds)
			sql, con = common.TimePeriod(sql, con, "", context)
			stm, err := db.Prepare(sql)
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				rows, err := stm.Query(con...)
				if nil != err {
					log.Error(err)
					common.StandardError(context)
				} else {
					var result []model.PlatformFans
					for rows.Next() {
						var model = model.PlatformFans{}
						err := rows.Scan(&model.Ids,
							&model.CreateTime,
							&model.Sum,
							&model.Decrease,
							&model.Increase,
							&model.Type,
							&model.PlatIds)
						if nil == err {
							result = append(result, model)
						}
					}
					context.JSON(http.StatusOK, result)
				}
			}
		} else {
			stm, err := plat.GetPlatformFansStmt()
			if nil != err {
				log.Error(err)
				context.Status(http.StatusInternalServerError)
			} else {
				rows, err := stm.Query()
				defer rows.Close()
				if nil != err {
					log.Error(err)
					common.StandardError(context)
				} else {
					var result []model.PlatformFans
					for rows.Next() {
						var model = model.PlatformFans{}
						err := rows.Scan(&model.Ids,
							&model.CreateTime,
							&model.Sum,
							&model.Decrease,
							&model.Increase,
							&model.Type,
							&model.PlatIds)
						if nil == err {
							result = append(result, model)
						}
					}
					context.JSON(http.StatusOK, result)
				}
			}
		}
	}
}

func PostPlatFans(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.PlatformFans{}
		common.ReadJSON(c, &m)
		stm, err := plat.PostPlatformFans()
		if nil != err {
			log.Error(err)
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.CreateTime,
				m.Sum,
				m.Decrease,
				m.Increase,
				m.PlatIds)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}

func PutPlatFans(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.PlatformFans{}
		common.ReadJSON(context, &m)
		stm, err := plat.PutPlatformFans()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec(m.CreateTime,
				m.Sum,
				m.Increase,
				m.Decrease,
				m.PlatIds,
				m.Ids)
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func DelPlatFans(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := plat.DelPlatformFans()
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
