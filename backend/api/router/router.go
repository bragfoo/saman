package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bragfoo/saman/backend/api/controller"
	"github.com/bragfoo/saman/backend/api/global"
)

func Server(router *gin.RouterGroup, g *global.G) {
	conf := g.Conf

	v1 := router.Group("/" + conf.Api.Version)
	// GET
	v1.GET("/platformFans", controller.GetPlatFans(g))
	v1.GET("platformType", controller.GetPlatType(g))
	v1.GET("/event", controller.GetEvent(g))
	v1.GET("/mobileData", controller.GetMobileData(g))
	v1.GET("/appUGC", controller.GetUGC(g))
	v1.GET("/playAmount", controller.GetVideoPlayAmount(g))
	v1.GET("/video", controller.GetVideo(g))
	v1.GET("/videoSource", controller.GetVideoSource(g))
	v1.GET("/channel", controller.GetChannel(g))
	v1.GET("/appData", controller.GetAppData(g))
	v1.GET("/talent", controller.GetTalent(g))
	v1.GET("/skill", controller.GetSkill(g))

	// liner
	v1.GET("/liner/playAmount", controller.GetLinerPlayAmountByPlat(g))

	// POST
	v1.POST("/platformFans", controller.PostPlatFans(g))
	v1.POST("/event", controller.PostEvent(g))
	v1.POST("/mobileData", controller.PostMobileData(g))
	v1.POST("/appUGC", controller.PostUGC(g))
	v1.POST("/playAmount", controller.PostVideoPlayAmount(g))
	v1.POST("/video", controller.PostVideo(g))
	v1.POST("/channel", controller.PostChannel(g))
	v1.POST("/appData", controller.PostAppData(g))
	v1.POST("/talent", controller.PostTalent(g))
	v1.POST("/skill", controller.PostSkill(g))

	v1.POST("/upload", controller.Upload(g))

	//PUT
	v1.PUT("/platformFans", controller.PutPlatFans(g))
	v1.PUT("/event", controller.PutEvent(g))
	v1.PUT("/mobileData", controller.PutMobileData(g))
	v1.PUT("/appUGC", controller.PutUGC(g))
	v1.PUT("/playAmount", controller.PutVideoPlayAmount(g))
	v1.PUT("/video", controller.PutVideo(g))
	v1.PUT("/channel", controller.PutChannel(g))
	v1.PUT("/appData", controller.PutAppData(g))
	v1.PUT("/talent", controller.PutTalent(g))
	v1.PUT("/skill", controller.PutSkill(g))

	//DELETE
	v1.DELETE("/platformFans", controller.GetPing(g))
	v1.DELETE("/event", controller.GetPing(g))
	v1.DELETE("/mobileData", controller.GetPing(g))
	v1.DELETE("/appUGC", controller.GetPing(g))
	v1.DELETE("/playAmount", controller.GetPing(g))
	v1.DELETE("/video", controller.GetPing(g))
	v1.DELETE("/channel", controller.GetPing(g))
	v1.DELETE("/appData", controller.GetPing(g))
	v1.DELETE("/talent", controller.GetPing(g))
	v1.DELETE("/skill", controller.GetPing(g))

	//DELETE_BY_ID
	v1.DELETE("/platformFans/:ids", controller.DelPlatFans(g))
	v1.DELETE("/event/:ids", controller.DelEvent(g))
	v1.DELETE("/mobileData/:ids", controller.DelMobileData(g))
	v1.DELETE("/appUGC/:ids", controller.DelUGC(g))
	v1.DELETE("/playAmount/:ids", controller.DelVideoPlayAmount(g))
	v1.DELETE("/video/:ids", controller.DelVideo(g))
	v1.DELETE("/channel/:ids", controller.DelChannel(g))
	v1.DELETE("/appData/:ids", controller.DelAppData(g))
	v1.DELETE("/talent/:ids", controller.DelTalent(g))
	v1.DELETE("/skill/:ids", controller.DelSkill(g))

	//upload
	upload := v1.Group("/upload")
	upload.POST("/fans", controller.UploadFans(g))
	upload.POST("/video", controller.UploadVideo(g))
	upload.POST("/newMedia", controller.Upload(g))

	//test
	test := router.Group("/test")
	test.GET("/one", controller.GetPlatType(g))
}
