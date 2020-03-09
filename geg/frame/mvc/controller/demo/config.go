package demo

import (
	"github.com/zhongdalu/gf/g/frame/gins"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

func init() {
	ghttp.GetServer().BindHandler("/config", func(r *ghttp.Request) {
		r.Response.Write(gins.Config().GetString("database.default.0.host"))
	})
}
