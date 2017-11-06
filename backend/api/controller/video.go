package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/backend/api/model/video"
)

func GetVideoPlayAmount(g *global.G) func(context *gin.Context) {
	return func(context *gin.Context) {
		stm, err := video.GetVideo()
		if nil != err {
			context.Status(http.StatusInternalServerError)
		} else {
			rows, _ := stm.Query()
			defer rows.Close()
			var result []model.VideoPlayAmount
			for rows.Next() {
				var model = model.VideoPlayAmount{}
				rows.Scan(&model.Ids,
					&model.Title,
					&model.Link,
					&model.CreateTime,
					&model.Sum,
					&model.NameChinese,
				)
				result = append(result, model)
			}
			context.JSON(http.StatusOK, result)
		}
	}
}
