package main

import (
	"github.com/zhongdalu/gf/g/os/glog"
)

func main() {
	l := glog.New()
	l.SetPrefix("[API]")
	l.Println("hello world")
	l.Error("error occurred")
}
