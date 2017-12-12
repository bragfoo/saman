package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bragfoo/saman/util"
)

func GetIds(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, util.GetObjectId())
	}
}
