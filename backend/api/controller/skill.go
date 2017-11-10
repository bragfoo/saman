package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/util"
	"github.com/bragfoo/saman/backend/api/model/skill"
)

func GetSkill(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Query("ids")
		name := c.Query("name")
		var con []interface{}
		sql := skill.GetQuery
		if "" != ids {
			sql += skill.WhereIds
			con = append(con, ids)
		}

		if "" != name {
			sql += skill.WhereName
			con = append(con, name)
		}

		stm, err := db.Prepare(sql)
		defer stm.Close()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			rows, err := stm.Query(con...)
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				var result = []model.Skill{}
				for rows.Next() {
					var m = model.Skill{}
					err := rows.Scan(&m.Ids,
						&m.Name,
					)
					if nil == err {
						result = append(result, m)
					}
				}
				common.StandardOk(c, result)

			}
		}

	}
}

func PostSkill(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.Skill{}
		common.ReadJSON(c, &m)
		stm, err := skill.PostSkill()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.Name,
			)
			common.CheckError(err, c)
		}
	}
}

func PutSkill(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.Skill{}
		common.ReadJSON(c, &m)
		stm, err := skill.PutSkill()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			_, err := stm.Exec(m.Name,
				m.Ids,
			)
			common.CheckError(err, c)
		}
	}
}

func DelSkill(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := skill.DelSkill()
			if nil != err {
				log.Error(err)
				common.StandardError(c)
			} else {
				_, err := stm.Exec(ids)
				common.CheckError(err, c)
			}
		}
	}
}
