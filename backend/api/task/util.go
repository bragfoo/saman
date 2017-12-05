package task

import (
	"time"
	"strings"
	"github.com/siddontang/go/log"
)

func getTimeInTime(sub int) (time.Time) {
	now := time.Now()
	d:=time.Hour*24*time.Duration(sub)
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	//ti := time.Date(t.Year(), t.Month(), t.Day()-sub, 0, 0, 0, 0, time.Local)
	ti := t.Add(time.Duration(-d))
	log.Info("insert time:")
	log.Info(ti)
	return ti
}

func getTime(sub int) (string) {
	t := time.Now()
	ti := time.Date(t.Year(), t.Month(), t.Day()-sub, 0, 0, 0, 0, time.Local)
	ret := strings.Replace(ti.String(), " +0800 CST", "", -1)
	log.Info("get time:")
	log.Info(ret)
	return ret
}
