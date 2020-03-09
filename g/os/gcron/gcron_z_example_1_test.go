// Copyright 2019 gf Author(https://github.com/zhongdalu/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhongdalu/gf.

package gcron_test

import (
	"github.com/zhongdalu/gf/g/os/gcron"
	"github.com/zhongdalu/gf/g/os/glog"
	"time"
)

func ExampleCron_AddSingleton() {
	gcron.AddSingleton("* * * * * *", func() {
		glog.Println("doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
