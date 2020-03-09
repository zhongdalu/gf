package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln(r.Cookie.Map())
	})
	s.SetPort(8199)
	s.Run()
}
