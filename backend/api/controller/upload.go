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
		excel.OpenNewVideo(file)
	}
}

func UploadFans(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		uploadFile, _ := c.FormFile("file")
		file, _ := uploadFile.Open()
		excel.OpenFans(file)
	}
}

func UploadVideo(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		uploadFile, _ := c.FormFile("file")
		file, _ := uploadFile.Open()
		excel.OpenPlatformVideo(file)
	}
}

func UploadCharts(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		uploadFile, _ := c.FormFile("file")
		file, _ := uploadFile.Open()
		excel.OpenCharts(file)
	}
}
