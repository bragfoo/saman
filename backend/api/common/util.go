package common

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

const timePeriod = "  AND %s.createTime >= ? AND %s.createTime <= ? "

func GetTimePeriod(query string, con []interface{}, nickName string) (string, []interface{}) {

	con = append(con, getWeekStartTime().Unix())
	con = append(con, getWeekStartTime().Add(time.Hour * 24 * 6).Unix())
	// statistics are updated everyday at 00:30.

	query += fmt.Sprintf(timePeriod, nickName, nickName)
	return query, con
}

func GetTwoWeekPeriod(query string, con []interface{}, nickName string) (string, []interface{}) {

	con = append(con, getWeekStartTime().Unix())
	con = append(con, getWeekStartTime().Add(time.Hour * 24 * 6 * 2).Unix())
	// statistics are updated everyday at 00:30.

	query += fmt.Sprintf(timePeriod, nickName, nickName)
	return query, con
}

func GetTimePeriodByPeriod(query string, con []interface{}, start interface{}, end interface{}, nickName string) (string, []interface{}) {
	con = append(con, start)
	con = append(con, end)
	query += fmt.Sprintf(timePeriod, nickName, nickName)
	return query, con
}

func GetTimePeriodByoffset(query string, con []interface{}, nickName string, offset time.Duration) (string, []interface{}) {

	con = append(con, getWeekStartTime().Unix())
	con = append(con, getWeekStartTime().Add(time.Hour * 24 * offset).Unix())
	// statistics are updated everyday at 00:30.

	query += fmt.Sprintf(timePeriod, nickName, nickName)
	return query, con
}

func getWeekStartTime() (time.Time) {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	switch t.Weekday() {
	case 0:
		//sunday
		return t.Add(-time.Hour * 24 * 13)
		break
	case 1:
		//monday
		t = t.Add(-time.Hour * 24 * 7)
		break
	case 2:
		t = t.Add(-time.Hour * 24 * 8)
		break
	case 3:
		t = t.Add(-time.Hour * 24 * 9)
		break
	case 4:
		t = t.Add(-time.Hour * 24 * 10)
		break
	case 5:
		t = t.Add(-time.Hour * 24 * 11)
		break
	case 6:
		t = t.Add(-time.Hour * 24 * 12)
		break
	default:
		t = time.Now()
	}
	return t
}

func TimePeriod(query string, con []interface{}, nickName string, context *gin.Context) (string, []interface{}) {
	start := context.Query("start")
	end := context.Query("end")
	if "" != start && "" != end {
		return GetTimePeriodByPeriod(query, con, start, end, nickName)
	} else {
		return query, con
	}
}
