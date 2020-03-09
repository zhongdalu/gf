package main

import (
	"fmt"
	"github.com/zhongdalu/gf/g"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	_, err := conn.Do("SUBSCRIBE", "channel")
	if err != nil {
		panic(err)
	}
	for {
		reply, err := conn.ReceiveVar()
		if err != nil {
			panic(err)
		}
		fmt.Println(reply.Strings())
	}
}
