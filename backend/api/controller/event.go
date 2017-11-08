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
	"github.com/bragfoo/saman/backend/api/common"
)

func GetEvent(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		ids := context.Query("ids")
		if "" != ids {
			stm, err := event.GetEventById()
			if nil != err {
				common.StandardError(context)
			} else {
				var model = model.Event{}
				err := stm.QueryRow(ids).Scan(&model.Ids,
					&model.Name,
					&model.StartDate,
					&model.EndDate,
					&model.TotalPeople,
					&model.TotalWork,
					&model.UploadPeople)
				if nil != err {
					log.Error(err)
					common.StandardNotFound(context)
				} else {
					common.StandardOk(context, model)
				}
			}
		} else {
			stm, err := event.GetEvent()
			if nil != err {
				common.StandardError(context)
			} else {
				rows, _ := stm.Query()
				defer rows.Close()
				var result []model.Event
				for rows.Next() {
					var model = model.Event{}
					err := rows.Scan(&model.Ids,
						&model.Name,
						&model.StartDate,
						&model.EndDate,
						&model.TotalPeople,
						&model.TotalWork,
						&model.UploadPeople)
					if nil == err {
						result = append(result, model)
					}
				}
				context.JSON(http.StatusOK, result)
			}
		}
	}
}

func PostEvent(g *global.G) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := c.Request.Body
		defer body.Close()
		b, _ := ioutil.ReadAll(body)
		var m = model.Event{}

		jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(b, &m)
		log.Info(m)
		stm, err := event.PostEvent()
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

func PutEvent(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.Event{}
		common.ReadJSON(context, &m)
		stm, err := event.PutEvent()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec(m.Name,
				m.UploadPeople,
				m.TotalWork,
				m.TotalPeople,
				m.EndDate,
				m.StartDate,
				m.Ids,
			)
			if nil != err {
				common.StandardError(context)
			} else {
				common.StandardSuccess(context)
			}
		}
	}
}

func DelEvent(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := event.DeleteEvent()
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
