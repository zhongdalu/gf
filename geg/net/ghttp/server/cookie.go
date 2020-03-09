package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
	"github.com/zhongdalu/gf/g/os/gtime"
)

func main() {
	s := g.Server()
	s.BindHandler("/cookie", func(r *ghttp.Request) {
		datetime := r.Cookie.Get("datetime")
		r.Cookie.Set("datetime", gtime.Datetime())
		r.Response.Write("datetime:", datetime)
	})
	s.SetPort(8199)
	s.Run()
}
