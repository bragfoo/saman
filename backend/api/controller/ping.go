package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bragfoo/saman/backend/api/global"
	"github.com/bragfoo/saman/util/db"
	"log"
)

func GetPing(g *global.G) func(*gin.Context) {
	return func(c *gin.Context) {
		pong := "pong"
		c.String(http.StatusOK, pong)
	}
}

func GetDbTest(g *global.G) func(*gin.Context) {
	return func(c *gin.Context) {
		row, err := db.Query("SELECT * FROM platformFans;")
		if nil != err {
			log.Fatal(err)
			c.String(http.StatusInternalServerError, "error")
		} else {
			str, _ := row.Columns()
			c.String(http.StatusOK, str[0])
		}
		row.Close()
	}
}
func GetDbTest2(g *global.G) func(*gin.Context) {
	return func(c *gin.Context) {
		Db := db.GetInstance()
		err := Db.Handle.Ping()
		if nil != err {
			log.Fatal(err)
		} else {
			if nil != err {
				log.Fatal(err)
				c.String(http.StatusInternalServerError, "error")
			} else {
			}
		}
	}
}
