package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/mobile"
	"io/ioutil"
	"github.com/json-iterator/go"
	"github.com/bragfoo/saman/backend/api/model/event"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
)

func GetMobileData(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		ids := context.Query("ids")
		if "" != ids {

		} else {
			stm, err := mobile.GetMobileData()
			if nil != err {
				context.Status(http.StatusInternalServerError)
			} else {
				rows, _ := stm.Query()
				defer rows.Close()
				var result []model.MobileData
				for rows.Next() {
					var model = model.MobileData{}
					err := rows.Scan(&model.Ids,
						&model.CreateTime,
						&model.Active,
						&model.Launch,
						&model.Channel,
						&model.SystemType,
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

func PostMobileData(g *global.G) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := c.Request.Body
		defer body.Close()
		b, _ := ioutil.ReadAll(body)
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		var m = model.MobileData{}
		log.Info(m)
		json.Unmarshal(b, &m)

		stm, err := event.PostEvents()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.CreateTime,
				m.Active,
				m.Launch,
				m.Channel,
				m.SystemType)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}
