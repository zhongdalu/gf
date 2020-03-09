package main

import (
	"github.com/zhongdalu/gf/g/os/glog"
	"time"
)

func main() {
	glog.SetAsync(true)
	for i := 0; i < 10; i++ {
		glog.Async().Print("async log", i)
	}
	time.Sleep(time.Second)
}
