package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g"
)

func main() {
	fmt.Println(g.Config().Get("log-path"))
}
