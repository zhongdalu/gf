package main

import (
	"fmt"
	"github.com/gogf/gf/g/net/gtcp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"time"
)

func main() {
	// Server
	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("receive:", data)
		}
	}).Run()

	time.Sleep(time.Second)

	// Client
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 10000; i++ {
		if err := conn.SendPkg([]byte(gconv.String(i))); err != nil {
			glog.Error(err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
