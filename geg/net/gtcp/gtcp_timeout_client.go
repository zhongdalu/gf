package main

import (
	"github.com/zhongdalu/gf/g/net/gtcp"
	"github.com/zhongdalu/gf/g/os/glog"
	"github.com/zhongdalu/gf/g/os/gtime"
	"time"
)

func main() {
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.Send([]byte(gtime.Now().String())); err != nil {
		glog.Error(err)
	}

	time.Sleep(time.Minute)
}
