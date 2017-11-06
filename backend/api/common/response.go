package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StandardError(context *gin.Context) {
	context.Status(http.StatusInternalServerError)
}

func StandardOk(context *gin.Context, obj interface{}) {
	context.JSON(http.StatusOK, obj)
}

func StandardNotFound(context *gin.Context)  {
	context.Status(http.StatusNotFound)
}

func StandardJoke(context *gin.Context)  {
	context.String(http.StatusInternalServerError,"are u kidding me?")
}
