package task

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/model"
	"time"
	"github.com/bragfoo/saman/util/config"
)

var query = "SELECT" +
	"  ids.type     AS type," +
	"  ids.day_time AS createTime," +
	"  ids.total    AS total" +
	"  FROM instreet_day_statistics ids" +
	"  WHERE 1 = 1"
var dbConf = config.DbType{
	Username: "instreet",
	Password: "instreetCoding@1by1",
	Address:  "rm-2ze33sny9l33w8pp8.mysql.rds.aliyuncs.com",
	Dbname:   "instreet4",
	Protocol: "tcp",
	Port:     "3306",
}

func GetAppStatistics() {

	stm, err := db.PrepareExt(query, &dbConf)
	if nil != err {
		log.Error(err)
	} else {
		rows, err := stm.Query()
		if nil != err {
			log.Error(err)
		} else {
			for rows.Next() {
				var model = model.DayStatistics{}
				rows.Scan(&model.Type, &model.CreateTime, &model.Total)
				log.Info(model)
			}
		}
	}
}

func Exec() {
	ticker := time.NewTicker(time.Hour * 6)
	go func() {
		for t := range ticker.C {
			log.Info("start timer process at:")
			log.Info(t)
			GetAppStatistics()
		}
	}()
}
