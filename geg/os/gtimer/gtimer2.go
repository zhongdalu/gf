package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g/container/gtype"
	"github.com/zhongdalu/gf/g/os/gtimer"
	"time"
)

func main() {
	v := gtype.NewInt()
	//w := gtimer.New(10, 10*time.Millisecond)
	fmt.Println("start:", time.Now())
	for i := 0; i < 1000000; i++ {
		gtimer.AddTimes(time.Second, 2, func() {
			v.Add(1)
		})
	}
	fmt.Println("end  :", time.Now())
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(v.Val(), time.Now())

	//gtimer.AddSingleton(time.Second, func() {
	//    fmt.Println(time.Now().String())
	//})
	//select { }
}
