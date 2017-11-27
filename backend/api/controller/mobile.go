package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/mobile"
	"github.com/bragfoo/saman/backend/api/model/event"
	"github.com/bragfoo/saman/util"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/util/db"
)

func GetMobileData(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		channelIds := context.Query("channelIds")
		systemType := context.Query("systemType")

		var con []interface{}
		sql := mobile.GetMobileDataQuery
		if "" != channelIds {
			sql += mobile.GetByChannelQuery
			con = append(con, channelIds)
		}

		if "" != systemType {
			sql += mobile.GetBySystemType
			con = append(con, systemType)
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
				var result []model.MobileData
				for rows.Next() {
					var model = model.MobileData{}
					err := rows.Scan(&model.Ids,
						&model.CreateTime,
						&model.Active,
						&model.Launch,
						&model.Channel,
						&model.SystemType,
						&model.ChannelIds,
					)
					if nil == err {
						result = append(result, model)
					}
				}
				context.JSON(http.StatusOK,result)
			}
		}

		//if "" != ids {
		//
		//} else if "" == ids && "" != channelIds {
		//	stm, err := mobile.GetMobileDataByChannel()
		//	if nil != err {
		//		log.Error(err)
		//		common.StandardError(context)
		//	} else {
		//		rows, err := stm.Query(channelIds)
		//		defer rows.Close()
		//		if nil != err {
		//			log.Error(err)
		//			common.StandardError(context)
		//		} else {
		//			var result []model.MobileData
		//			for rows.Next() {
		//				var model = model.MobileData{}
		//				err := rows.Scan(&model.Ids,
		//					&model.CreateTime,
		//					&model.Active,
		//					&model.Launch,
		//					&model.Channel,
		//					&model.SystemType,
		//					&model.ChannelIds,
		//				)
		//				if nil == err {
		//					result = append(result, model)
		//				}
		//			}
		//			context.JSON(http.StatusOK, result)
		//		}
		//	}
		//} else {
		//	stm, err := mobile.GetMobileData()
		//	if nil != err {
		//		context.Status(http.StatusInternalServerError)
		//	} else {
		//		rows, _ := stm.Query()
		//		defer rows.Close()
		//		var result []model.MobileData
		//		for rows.Next() {
		//			var model = model.MobileData{}
		//			err := rows.Scan(&model.Ids,
		//				&model.CreateTime,
		//				&model.Active,
		//				&model.Launch,
		//				&model.Channel,
		//				&model.SystemType,
		//				&model.ChannelIds,
		//			)
		//			if nil == err {
		//				result = append(result, model)
		//			}
		//		}
		//		context.JSON(http.StatusOK, result)
		//	}
		//}
	}
}

func PostMobileData(g *global.G) func(c *gin.Context) {
	return func(c *gin.Context) {
		var m = model.MobileData{}
		common.ReadJSON(c, &m)
		stm, err := event.PostEvent()
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

func PutMobileData(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		var m = model.MobileData{}
		common.ReadJSON(context, &m)
		stm, err := mobile.PutMobileData()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			_, err := stm.Exec(m.CreateTime,
				m.ChannelIds,
				m.SystemType,
				m.Launch,
				m.Active,
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

func DelMobileData(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := mobile.DelMObileData()
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
