package main

import (
	"github.com/zhongdalu/gf/g/os/gfsnotify"
	"github.com/zhongdalu/gf/g/os/glog"
)

func main() {
	//path := `D:\temp`
	path := "/Users/john/Temp"
	_, err := gfsnotify.Add(path, func(event *gfsnotify.Event) {
		glog.Println(event)
	})
	if err != nil {
		glog.Fatal(err)
	} else {
		select {}
	}
}
