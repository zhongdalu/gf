package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g/os/gcmd"
)

func doEcho() {
	fmt.Println("do echo")
}

func main() {
	fmt.Println(gcmd.Value.GetAll())

	fmt.Println(gcmd.Value.Get(1))

	gcmd.BindHandle("echo", doEcho)
	gcmd.RunHandle("echo")

	gcmd.AutoRun()
}
