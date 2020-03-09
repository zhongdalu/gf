package main

import (
	"github.com/zhongdalu/gf/g/os/glog"
	"github.com/zhongdalu/gf/g/os/gtime"
	"github.com/zhongdalu/gf/g/os/gtimer"
	"time"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		glog.SetDebug(false)
	})
	for {
		glog.Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
