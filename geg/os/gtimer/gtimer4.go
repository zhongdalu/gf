package main

import (
	"github.com/zhongdalu/gf/g/os/glog"
	"github.com/zhongdalu/gf/g/os/gtimer"
	"time"
)

func main() {
	interval := time.Second
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing1")
	})
	gtimer.AddTimes(interval, 2, func() {
		glog.Println("doing2")
	})

	select {}
}
