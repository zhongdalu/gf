package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g/os/grpool"
	"github.com/zhongdalu/gf/g/os/gtime"
	"sync"
	"time"
)

func main() {
	start := gtime.Millisecond()
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		grpool.Add(func() {
			time.Sleep(time.Second)
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println(grpool.Size())
	fmt.Println("time spent:", gtime.Millisecond()-start)
}
