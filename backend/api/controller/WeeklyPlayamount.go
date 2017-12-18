package controller

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"net/http"
	"github.com/bragfoo/saman/backend/api/model/platPlayAmount"
	"github.com/gin-gonic/gin"
)

func GetWeeklyPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {

		var con []interface{}
		sql := platPlayAmount.GetWeeklyPlayAmount
		s, c := common.GetTimePeriod(sql, con, "pA")
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
				var result []interface{}
				for rows.Next() {
					var m = make(map[string]interface{})
					var sum int64
					var platIds string
					var platName string
					err := rows.Scan(&sum,
						&platIds,
						&platName,
					)
					m["sum"] = sum
					m["platIds"] = platIds
					m["platName"] = platName

					if nil == err {
						result = append(result, m)
					}
				}
				context.JSON(http.StatusOK, result)
			}
		}
	}

}
