package common

import (
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
	"github.com/json-iterator/go"
)

const (
	EQ = "eq"
	GT = "gt"
	LT = "lt"
)

type Payload struct {
	rows map[string]string
	cols []Cols
}

type Cols struct {
	colName string
	value   string
	cmp     string
}

func ReadJSON(context *gin.Context, obj interface{}) {
	body := context.Request.Body
	defer body.Close()
	b, _ := ioutil.ReadAll(body)
	jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(b, &obj)
	log.Debug(obj)
}

func process() {
}
