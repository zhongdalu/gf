package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/os/glog"
)

func main() {
	glog.Debug(g.Map{"uid": 100, "name": "john"})

	type User struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	glog.Debug(User{100, "john"})
}
