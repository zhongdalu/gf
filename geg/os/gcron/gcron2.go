package main

import (
	"github.com/zhongdalu/gf/g/os/gcron"
	"github.com/zhongdalu/gf/g/os/glog"
	"time"
)

func test() {
	glog.Println(111)
}

func main() {
	_, err := gcron.AddOnce("@every 2s", test)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
}
