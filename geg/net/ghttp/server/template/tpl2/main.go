package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/frame/gins"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetServerRoot("public")
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(2333)

	s.BindHandler("/", func(r *ghttp.Request) {
		content, _ := gins.View().Parse("test.html", nil)
		r.Response.Write(content)
	})

	s.Run()
}
