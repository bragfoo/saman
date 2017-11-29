package task

import (
	"time"
	"strings"
	"github.com/siddontang/go/log"
)

func getTimeInTime(sub int) (time.Time) {
	t := time.Now()
	ti := time.Date(t.Year(), t.Month(), t.Day()-sub, 0, 0, 0, 0, time.Local)
	log.Info(ti)
	return ti
}

func getTime(sub int) (string) {
	t := time.Now()
	ti := time.Date(t.Year(), t.Month(), t.Day()-sub, 0, 0, 0, 0, time.Local)
	ret := strings.Replace(ti.String(), " +0800 CST", "", -1)
	log.Info(ret)
	return ret
}
