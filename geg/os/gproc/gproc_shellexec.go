package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g/os/gproc"
)

// 执行shell指令
func main() {
	r, err := gproc.ShellExec(`sleep 3s; echo "hello gf!";`)
	fmt.Println("result:", r)
	fmt.Println(err)
}
