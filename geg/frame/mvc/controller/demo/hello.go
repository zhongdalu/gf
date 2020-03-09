package demo

import "github.com/zhongdalu/gf/g/net/ghttp"

func init() {
	ghttp.GetServer().BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello World!")
	})
}
