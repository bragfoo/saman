package task

import "time"

func getTime(sub int) (time.Time) {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()-sub, 0, 0, 0, 0, time.Local)
}
