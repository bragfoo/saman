package liner

import (
	"github.com/bragfoo/saman/util/db"
	"github.com/siddontang/go/log"
	"fmt"
	"time"
)

func GetPlatPercentage(platIds string) map[int64]string {
	return processMap(getMap(platIds))
}

func getMap(platIds string) map[int64]int {
	stm, err := db.Prepare(GetPlatGrowPercentage)
	if nil != err {
		log.Error(err)
		return nil
	} else {
		rows, _ := stm.Query(platIds)
		defer rows.Close()
		totalMap := make(map[int64]int)
		for rows.Next() {
			var createTime int64
			var grow int
			rows.Scan(&createTime, &grow)
			totalMap[createTime] = grow
		}
		return totalMap
	}
}

func processMap(m map[int64]int) map[int64]string {
	totalMap := make(map[int64]string)
	for k, v := range m {
		dateTime := time.Unix(k, 0)
		before := dateTime.Add(-time.Hour * 24 * 7)
		if 0 != m[before.Unix()] {
			thisWeek := float64(v)
			lastWeek := float64(m[before.Unix()])
			result := (thisWeek - lastWeek) / lastWeek
			totalMap[k] = processPercent(result)
		}
	}
	return totalMap
}

func processPercent(f float64) string {
	s := fmt.Sprintf("%.2f", f*100)
	s += "%"
	return s
}
