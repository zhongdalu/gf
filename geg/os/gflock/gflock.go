package main

import (
	"time"

	"github.com/zhongdalu/gf/g/os/gflock"
	"github.com/zhongdalu/gf/g/os/glog"
	"github.com/zhongdalu/gf/g/os/gproc"
)

func main() {
	l := gflock.New("demo.lock")
	l.Lock()
	glog.Printf("locked by pid: %d", gproc.Pid())
	time.Sleep(10 * time.Second)
	l.UnLock()
	glog.Printf("unlocked by pid: %d", gproc.Pid())
}
