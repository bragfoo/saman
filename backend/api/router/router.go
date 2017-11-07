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
	v1.GET("/events", controller.GetEvents(g))
	v1.GET("/mobileData", controller.GetMobileData(g))
	v1.GET("/appUGC", controller.GetUGC(g))
	v1.GET("/videoPlayAmount", controller.GetVideoPlayAmount(g))
	v1.GET("/video", controller.GetVideoPlayAmount(g))

	// POST
	v1.POST("/platformFans", controller.PostPlatFans(g))
	v1.POST("/events", controller.PostEvents(g))
	v1.POST("/mobileData", controller.PostMobileData(g))
	v1.POST("/appUGC", controller.PostUGC(g))
	v1.POST("/videoPlayAmount", controller.PostVideoPlayAmount(g))
	v1.POST("/video", controller.PostVideo(g))

	//PUT
	v1.PUT("/platformFans", controller.PutPlatFans(g))
	v1.PUT("/events", controller.PutEvents(g))
	v1.PUT("/mobileData", controller.PutMobileData(g))
	v1.PUT("/appUGC", controller.PutUGC(g))
	v1.PUT("/videoPlayAmount", controller.PutVideoPlayAmount(g))
	v1.PUT("/video", controller.PutVideo(g))

	//DELETE
	v1.DELETE("/platformFans", controller.GetPing(g))
	v1.DELETE("/events", controller.GetPing(g))
	v1.DELETE("/mobileData", controller.GetPing(g))
	v1.DELETE("/appUGC", controller.GetPing(g))
	v1.DELETE("/videoPlayAmount", controller.GetPing(g))
	v1.DELETE("/video", controller.GetPing(g))

	//DELETE_BY_ID
	v1.DELETE("/platformFans/:ids", controller.DelPlatFans(g))
	v1.DELETE("/events/:ids", controller.DelEvents(g))
	v1.DELETE("/mobileData/:ids", controller.DelMobileData(g))
	v1.DELETE("/appUGC/:ids", controller.DelUGC(g))
	v1.DELETE("/videoPlayAmount/:ids", controller.DelVideoPlayAmount(g))
	v1.DELETE("/video/:ids", controller.DelVideo(g))

	//test
	test := router.Group("/test")
	test.GET("/one", controller.GetPlatType(g))
}
