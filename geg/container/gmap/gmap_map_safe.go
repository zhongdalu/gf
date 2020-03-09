package main

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/container/gmap"
)

func main() {
	m := gmap.New()
	m.Set("1", "1")

	m1 := m.Map()
	m1["2"] = "2"

	g.Dump(m.Clone())
	g.Dump(m1)
}
