package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/ugc"
	"io/ioutil"
	"github.com/json-iterator/go"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
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
				)
				if nil != err {
					log.Error(err)
					common.StandardNotFound(context)
				} else {
					common.StandardOk(context, model)
				}
			}
		} else {
			stm, err := ugc.GetUGC()
			if nil != err {
				context.Status(http.StatusInternalServerError)
			} else {
				rows, _ := stm.Query()
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
		body := c.Request.Body
		defer body.Close()
		b, _ := ioutil.ReadAll(body)
		var m = model.AppUGC{}
		jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(b, &m)
		log.Info(m)
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
				m.VideoSum)
			if nil != err {
				log.Error(err)
				c.Status(http.StatusInternalServerError)
			} else {
				c.Status(http.StatusOK)
			}
		}
	}
}
