package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model/liner"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"net/http"
	"time"
)

func GetLinerPlayAmountByPlat(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		platIds := c.Query("platIds")

		var con []interface{}
		var sql = liner.GetPlayAmount
		if "" != platIds {
			sql += liner.GetPlayAmountByPlatIds
			con = append(con, platIds)
		}

		// constant start at 2017-6-1
		if true {
			sql += liner.GetTime
			con = append(con, time.Date(2017, time.August, 1, 0, 0, 0, 0, time.Local).Unix())
		}
		stm, err := db.Prepare(sql)
		defer stm.Close()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			rows, err := stm.Query(con...)
			defer rows.Close()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				var result []model.PlatPlayAmount
				for rows.Next() {
					var m = model.PlatPlayAmount{}
					err := rows.Scan(&m.Ids, &m.CreateTime, &m.PlatType, &m.Sum, &m.Grow)
					if nil == err {
						result = append(result, m)
					}
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}

func GetGrowPercentage(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		platIds := c.Query("platIds")
		if "" != platIds {
			m := liner.GetPlatPercentage(platIds)
			c.JSON(http.StatusOK, m)
		} else {
			common.StandardBadRequest(c)
		}
	}
}
