package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/siddontang/go/log"
)

func StandardError(context *gin.Context) {
	context.Status(http.StatusInternalServerError)
}

func StandardOk(context *gin.Context, obj interface{}) {
	context.JSON(http.StatusOK, obj)
}

func StandardSuccess(context *gin.Context) {
	context.Status(http.StatusOK)
}

func StandardNotFound(context *gin.Context) {
	context.Status(http.StatusNotFound)
}

func StandardJoke(context *gin.Context) {
	context.String(http.StatusInternalServerError, "are u kidding me?")
}

func StandardBadRequest(context *gin.Context) {
	context.Status(http.StatusBadRequest)
}

func CheckError(err error, context *gin.Context) {
	if nil != err {
		log.Error(err)
		StandardError(context)
	} else {
		StandardSuccess(context)
	}
}
