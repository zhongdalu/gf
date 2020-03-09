package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/os/gfile"
	"github.com/zhongdalu/gf/g/os/glog"
)

// 设置日志输出路径
func main() {
	path := "/tmp/glog"
	glog.SetPath(path)
	glog.Println("日志内容")
	list, err := gfile.ScanDir(path, "*")
	g.Dump(err)
	g.Dump(list)
}
