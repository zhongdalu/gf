package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/util/gconv"
)

func main() {
	type User struct {
		Id   int    `json:"uid"`
		Name string `my-tag:"nick-name" json:"name"`
	}
	user := &User{
		Id:   1,
		Name: "john",
	}
	g.Dump(gconv.Map(user, "my-tag"))
}
