package demo

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

type Order struct{}

func init() {
	g.Server().BindObject("/{.struct}-{.method}", new(Order))
}

func (o *Order) List(r *ghttp.Request) {
	r.Response.Write("List")
}
