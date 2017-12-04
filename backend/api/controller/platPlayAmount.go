package controller

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"net/http"
	"github.com/bragfoo/saman/backend/api/model/platPlayAmount"
)

func GetPlatPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {

		platIds := context.Query("platIds")

		var con []interface{}
		sql := platPlayAmount.GetPlatPlayAmount
		if "" != platIds {
			sql += platPlayAmount.WhereByPlatIds
			con = append(con, platIds)
		}
		s, c := common.GetTimePeriod(sql, con, "pPA")
		stm, err := db.Prepare(s)
		defer stm.Close()
		if nil != err {
			log.Error(err)
			common.StandardError(context)
		} else {
			rows, err := stm.Query(c...)
			if nil != err {
				log.Error(err)
				common.StandardError(context)
			} else {
				var result []model.PlatPlayAmount
				for rows.Next() {
					var model = model.PlatPlayAmount{}
					err := rows.Scan(&model.Ids,
						&model.Sum,
						&model.PlatType,
						&model.CreateTime,
						&model.Grow,
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
