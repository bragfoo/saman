package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model/channel"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
)

func GetChannel(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		stm, err := channel.GetChannel()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.Channel
			for rows.Next() {
				var m = model.Channel{}
				err := rows.Scan(
					&m.Ids,
					&m.Name,
				)
				if nil == err {
					result = append(result, m)
				}
			}
			common.StandardOk(c, result)
		}
	}
}
