package main

import (
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gmlock"
	"sync"
	"time"
)

// 内存锁基本使用
func main() {
	key := "lock"
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			gmlock.Lock(key)
			glog.Println(i)
			time.Sleep(time.Second)
			gmlock.Unlock(key)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
