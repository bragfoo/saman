package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bragfoo/saman/backend/api/controller"
	"github.com/bragfoo/saman/backend/api/global"
)

func Server(router *gin.RouterGroup, g *global.G) {
	conf := g.Conf

	v1 := router.Group("/" + conf.Api.Version)
	// dashboard
	v1.GET("/ping", controller.GetPing(g))
}
