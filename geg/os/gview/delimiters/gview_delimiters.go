package main

import (
	"fmt"
	"github.com/gogf/gf/g"
)

func main() {
	v := g.View()
	v.SetDelimiters("${", "}")
	b, err := v.Parse("gview_delimiters.tpl", map[string]interface{}{
		"k": "v",
	})
	fmt.Println(err)
	fmt.Println(string(b))
}
