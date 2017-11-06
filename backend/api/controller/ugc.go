package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/ugc"
)

func GetUGC(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := ugc.GetUGC()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.AppUGC
			for rows.Next() {
				var model = model.AppUGC{}
				rows.Scan(&model.Ids,
					&model.CreateTime,
					&model.Like,
					&model.CommentSum,
					&model.ShareSum,
					&model.PicSum,
					&model.VideoSum,
				)
				result = append(result, model)
			}
			context.JSON(http.StatusOK, result)
		}
	}
}
