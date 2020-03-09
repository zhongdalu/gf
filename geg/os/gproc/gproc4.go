package main

import (
	"github.com/zhongdalu/gf/g/os/genv"
	"github.com/zhongdalu/gf/g/os/glog"
	"github.com/zhongdalu/gf/g/os/gproc"
	"os"
	"time"
)

// 查看父子进程的环境变量
func main() {
	time.Sleep(5 * time.Second)
	glog.Printf("%d: %v", gproc.Pid(), genv.All())
	p := gproc.NewProcess(os.Args[0], os.Args, os.Environ())
	p.Start()
}
