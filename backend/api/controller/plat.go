package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/plattype"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
)

func GetPlatType(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		ids := context.Query("ids")
		if "" != ids {
			log.Info(ids)
		} else
		{
			stm, err := plat.GetPlatformTypeStmt()
			if nil != err {
				log.Error(err)
				context.Status(http.StatusInternalServerError)
			} else {
				rows, err := stm.Query()
				if err != nil {
					log.Error(err)
					common.StandardJoke(context)
				}
				defer rows.Close()
				var result []model.PlatformType
				for rows.Next() {
					var model = model.PlatformType{}
					err := rows.Scan(&model.Ids, &model.Name, &model.NameChinese)
					if nil == err {
						result = append(result, model)
					}
				}
				common.ResponseList(context, result)
			}
		}
	}
}
