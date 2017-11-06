package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/plattype"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model/event"
	"io/ioutil"
	"github.com/json-iterator/go"
)

func GetPlatType(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := plat.GetPlatformTypeStmt()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
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

func GetPlatFans(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := plat.GetPlatformFansStmt()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.PlatformFans
			for rows.Next() {
				var model = model.PlatformFans{}
				err := rows.Scan(&model.Ids,
					&model.CreateTime,
					&model.Sum,
					&model.Decrease,
					&model.Increase,
					&model.Type)
				if nil != err {
					log.Error(err)
				}
			}
			context.JSON(http.StatusOK, result)
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
		stm, err := event.PostEvents()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.CreateTime,
				m.Sum,
				m.Decrease,
				m.Increase,
				m.Type)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}
