package main

import (
	"github.com/zhongdalu/gf/g"
	_ "github.com/zhongdalu/gf/geg/frame/mvc/controller/demo"
	_ "github.com/zhongdalu/gf/geg/frame/mvc/controller/stats"
)

func main() {

	//g.Server().SetDumpRouteMap(false)
	g.Server().SetPort(8199)
	g.Server().Run()

}
