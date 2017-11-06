package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/plattype"
	"github.com/bragfoo/saman/backend/api/util"
)

func GetPlatType(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		util.AnalyzeRequest(context)
		stm, err := plat.GetPlatformTypeStmt()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.PlatformType
			for rows.Next() {
				var model = model.PlatformType{}
				rows.Scan(&model.Ids, &model.Name, &model.NameChinese)
				result = append(result, model)
			}
			context.JSON(http.StatusOK, result)
		}
	}
}

func GetPlatFans(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		util.AnalyzeRequest(context)
		stm, err := plat.GetPlatformFansStmt()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.PlatformFans
			for rows.Next() {
				var model = model.PlatformFans{}
				rows.Scan(&model.Ids,
					&model.CreateTime,
						&model.Sum,
							&model.Decrease,
								&model.Increase,
									&model.Type)
				result = append(result, model)
			}
			context.JSON(http.StatusOK, result)
		}
	}
}
