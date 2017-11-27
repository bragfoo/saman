package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/util/excel"
)

func Upload(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		uploadFile, _ := c.FormFile("file")
		file, _ := uploadFile.Open()
		excel.OpenSpace(file)
	}
}
