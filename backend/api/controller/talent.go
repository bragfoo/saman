package controller

import (
	"github.com/bragfoo/saman/backend/api/global"
	"github.com/gin-gonic/gin"
	"github.com/bragfoo/saman/backend/api/model/talent"
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"github.com/bragfoo/saman/backend/api/common"
	"github.com/bragfoo/saman/backend/api/model"
	"github.com/bragfoo/saman/util"
)

func GetTalent(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Query("ids")
		user := c.Query("user")
		skillType := c.Query("type")
		submitted := c.Query("submitted")
		var con []interface{}
		sql := talent.GetQuery
		if "" != ids {
			sql += talent.WhereIds
			con = append(con, ids)
		}
		if "" != user {
			sql += talent.WhereUser
			con = append(con, user)
		}

		if "" != skillType {
			sql += talent.WhereType
			con = append(con, skillType)
		}
		if "" != submitted {
			sql += talent.WhereSubmitted
			con = append(con, submitted)
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
				var result = []model.Talent{}
				for rows.Next() {
					var m = model.Talent{}
					err := rows.Scan(&m.Ids,
						&m.Type,
						&m.User,
						&m.SkillName,
						&m.CreateTime,
						&m.Submitted,
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

func PostTalent(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.Talent{}
		common.ReadJSON(c, &m)
		stm, err := talent.PostTalnet()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			_, err := stm.Exec(util.GetObjectId(),
				m.User,
				m.Type,
				m.CreateTime,
				m.Submitted,
			)
			common.CheckError(err, c)
		}
	}
}

func PutTalent(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		var m = model.Talent{}
		common.ReadJSON(c, &m)
		stm, err := talent.PutQuery()
		if nil != err {
			log.Error(err)
			common.StandardError(c)
		} else {
			_, err := stm.Exec(m.Type,
				m.User,
				m.CreateTime,
				m.Submitted,
				m.Ids,
			)
			common.CheckError(err, c)
		}
	}
}

func DelTalent(g *global.G) func(context *gin.Context) {
	return func(c *gin.Context) {
		ids := c.Param("ids")
		if "" == ids {
			common.StandardBadRequest(c)
		} else {
			stm, err := talent.DelQuery()
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
