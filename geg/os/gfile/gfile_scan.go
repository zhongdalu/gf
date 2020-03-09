package main

import (
	"github.com/zhongdalu/gf/g/os/gfile"
	"github.com/zhongdalu/gf/g/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/home/john/Documents", "*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
