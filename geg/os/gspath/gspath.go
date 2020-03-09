package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/os/gspath"
	"github.com/zhongdalu/gf/g/os/gtime"
	"time"
)

func main() {
	sp := gspath.New()
	path := "/Users/john/Temp"
	rp, err := sp.Add(path)
	fmt.Println(err)
	fmt.Println(rp)
	fmt.Println(sp)

	gtime.SetInterval(5*time.Second, func() bool {
		g.Dump(sp.AllPaths())
		return true
	})

	select {}
}
