package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/event"
	"github.com/siddontang/go/log"
	"io/ioutil"
	"github.com/json-iterator/go"
	"github.com/bragfoo/saman/util"
)

func GetEvents(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := event.GetEvent()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.Event
			for rows.Next() {
				var model = model.Event{}
				rows.Scan(&model.Ids,
					&model.Name,
					&model.StartDate,
					&model.EndDate,
					&model.TotalPeople,
					&model.TotalWork,
					&model.UploadPeople)
				result = append(result, model)
			}
			context.JSON(http.StatusOK, result)
		}
	}
}

func PostEvents(g *global.G) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := c.Request.Body
		defer body.Close()
		b, _ := ioutil.ReadAll(body)
		var m = model.Event{}

		jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(b, &m)
		log.Info(m)
		stm, err := event.PostEvents()
		if nil != err {
			c.Status(http.StatusInternalServerError)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.Name,
				m.StartDate,
				m.EndDate,
				m.TotalPeople,
				m.TotalWork,
				m.UploadPeople)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}
