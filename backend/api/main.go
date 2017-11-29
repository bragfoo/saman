package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/koding/multiconfig"

	"github.com/bragfoo/saman/backend/api/global"
	api "github.com/bragfoo/saman/backend/api/router"
	"github.com/bragfoo/saman/util/config"
	"github.com/bragfoo/saman/backend/api/task"
)

var g *global.G

func initConfig() {
	// config
	confPath := "conf/config.toml"
	m := multiconfig.NewWithPath(confPath)
	conf := &config.ConfType{}

	m.MustLoad(conf) // Panic's if there is any error
	g = &global.G{Conf: conf}
}

func main() {
	initConfig()
	conf := g.Conf

	r := gin.Default()
	// load api router
	if conf.Api.Enable {
		api.Server(r.Group("/"+conf.Api.Prefix), g)
	}

	port := fmt.Sprintf(":%d", conf.Api.Port)

	//task.Exec()
	task.RunCrons()
	r.Run(port)

}
