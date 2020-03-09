package demo

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

type Object struct{}

func init() {
	g.Server().BindObject("/object", new(Object))
}

func (o *Object) Index(r *ghttp.Request) {
	r.Response.Write("object index")
}

func (o *Object) Show(r *ghttp.Request) {
	r.Response.Write("object show")
}
