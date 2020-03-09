package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g/util/gutil"
)

func main() {
	gutil.TryCatch(func() {
		fmt.Println(1)
		gutil.Throw("error")
		fmt.Println(2)
	}, func(err interface{}) {
		fmt.Println(err)
	})
}
