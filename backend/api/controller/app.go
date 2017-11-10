package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/util"
	"github.com/bragfoo/saman/backend/api/model/app"
)

func GetAppData(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		stm, err := app.GetAppData()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.AppData
			for rows.Next() {
				var m = model.AppData{}
				err := rows.Scan(
					&m.Ids,
					&m.CreateTime,
					&m.ActiveUser,
					&m.PicUpload,
					&m.TalentSum,
					&m.VideoUpload,
				)
				if nil == err {
					result = append(result, m)
				}
			}
			common.StandardOk(c, result)
		}
	}
}

func PostAppData(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.AppData{}
		common.ReadJSON(c, &m)
		stm, err := app.PostAppData()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.CreateTime,
				m.PicUpload,
				m.VideoUpload,
				m.TalentSum,
				m.ActiveUser,
			)
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				common.StandardSuccess(c)
			}
		}
	}
}

func PutAppData(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.AppData{}
		common.ReadJSON(c, &m)
		stm, err := app.PutAppData()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			_, err := stm.Exec(m.CreateTime,
				m.ActiveUser,
				m.PicUpload,
				m.TalentSum,
				m.VideoUpload,
				m.Ids,
			)
			common.CheckError(err, c)
		}
	}
}

func DelAppData(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := app.DelAppData()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				_, err := stm.Exec(ids)
				common.CheckError(err, c)
			}
		}
	}
}
