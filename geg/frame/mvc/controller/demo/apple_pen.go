package demo

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindHandler("/apple", Apple)
	s.BindHandler("/pen", Pen)
	s.BindHandler("/apple-pen", ApplePen)
}

func Apple(r *ghttp.Request) {
	r.Response.Write("Apple")
}

func Pen(r *ghttp.Request) {
	r.Response.Write("Pen")
}

func ApplePen(r *ghttp.Request) {
	r.Response.Write("Apple-Pen")
}
