package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/os/gfile"
	"github.com/zhongdalu/gf/g/os/glog"
)

func main() {
	path := "/tmp/glog-cat"
	glog.SetPath(path)
	glog.Stdout(false).Cat("cat1").Cat("cat2").Println("test")
	list, err := gfile.ScanDir(path, "*", true)
	g.Dump(err)
	g.Dump(list)
}
