package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/container/garray"
	"strings"
)

func main() {
	array := garray.NewSortedStringArray()
	array.Add("/api/ctl/show")
	array.Add("/api/ctl/post")
	array.Add("/api/obj/rest")
	array.Add("/api/handler")
	array.Add("/api/obj/delete")
	array.Add("/api/obj/show")
	array.Add("/api/obj/my-show")
	array.Add("/api/*")
	array.Add("/api/ctl/rest")
	array.Add("/api/ctl/my-show")
	g.Dump(array.Slice())

	fmt.Println(strings.Compare("/api/ctl/post", "/api/*"))
	fmt.Println(strings.Compare("/api/*", "/api/ctl/my-show"))
}
