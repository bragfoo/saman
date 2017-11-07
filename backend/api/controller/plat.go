package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/plattype"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
	"io/ioutil"
	"github.com/json-iterator/go"
	"github.com/bragfoo/saman/backend/api/common"
)

func GetPlatType(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		ids := context.Query("ids")
		if "" != ids {
			log.Info(ids)
		} else
		{
			stm, err := plat.GetPlatformTypeStmt()
			if nil != err {
				log.Error(err)
				context.Status(http.StatusInternalServerError)
			} else {
				rows, err := stm.Query()
				if err != nil {
					log.Error(err)
					common.StandardJoke(context)
				}
				defer rows.Close()
				var result []model.PlatformType
				for rows.Next() {
					var model = model.PlatformType{}
					err := rows.Scan(&model.Ids, &model.Name, &model.NameChinese)
					if nil == err {
						result = append(result, model)
					}
				}
				context.JSON(http.StatusOK, result)
			}
		}
	}
}

func GetPlatFans(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		platIds := context.Query("platIds")
		if "" != platIds {
			stm, err := plat.GetPlatformFansByPlatIds()
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				rows, err := stm.Query(platIds)
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
		body := c.Request.Body
		defer body.Close()
		b, _ := ioutil.ReadAll(body)
		var m = model.PlatformFans{}

		jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(b, &m)
		log.Info(m)
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
