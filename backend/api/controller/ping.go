package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bragfoo/saman/backend/api/global"
)

func GetPing(g *global.G) func(*gin.Context) {
	return func(c *gin.Context) {
		pong := "pong"
		c.String(http.StatusOK, pong)
	}
}
