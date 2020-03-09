package main

import (
	"github.com/zhongdalu/gf/g/util/gutil"
)

func Test(s *interface{}) {
	//debug.PrintStack()
	gutil.PrintStack()
}

func main() {
	Test(nil)
}
