package common

import "time"

const timePeriod = "  AND createTime > ? AND createTime < ? "

func GetTimePeriod(query string, con []interface{}, startTime int64, endTime int64) (string, []interface{}) {

	if startTime != 0 && endTime != 0 {
		con = append(con, startTime)
		con = append(con, endTime)
		query += timePeriod
		return query, con
	} else {
		time.Now()
		return query, con
	}
}

func GetWeekStartTime(t time.Time) (time.Time) {

	switch t.Weekday() {
	case 0:
		//sunday

		break
	case 1:
		break
	case 2:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		break
	case 6:
		break
	}

	return time.Now()

}
