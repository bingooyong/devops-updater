package cron

import (
	"math/rand"
	"ops.updater/g"
	"time"
)

func SleepRandomDuration() {
	ns := int64(g.Config().Interval) * 1000000000
	// 以当前时间为随机数种子，如果所有ops-updater在同一时间启动，系统时间是相同的，那么随机种子就是一样的
	// 问题不大，批量ssh去启动ops-updater的话也是一个顺次的过程
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d := time.Duration(r.Int63n(ns)) * time.Nanosecond
	time.Sleep(d)
}
