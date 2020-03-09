package stats

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/net/ghttp"
)

var (
	total int
)

func init() {
	g.Server().BindHandler("/stats/total", showTotal)
}

func showTotal(r *ghttp.Request) {
	total++
	r.Response.Write("total:", total)
}
