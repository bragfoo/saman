package util

import (
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
)

func AnalyzeRequest(context *gin.Context) (*common.Payload) {
	var payload common.Payload
	log.Info("Request")
	log.Info(context.Request)
	log.Info("Params")
	log.Info(context.Params)

	return &payload
}
