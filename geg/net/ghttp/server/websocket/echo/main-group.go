package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
	"github.com/zhongdalu/gf/g/os/gfile"
	"github.com/zhongdalu/gf/g/os/glog"
)

func ws(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		if err = ws.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func main() {
	s := g.Server()
	s.Group().Bind([]ghttp.GroupItem{
		{"ALL", "/ws", ws},
	})

	s.SetServerRoot(gfile.MainPkgPath())
	s.SetPort(8199)
	s.Run()
}
