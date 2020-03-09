// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Timer Operations

package gtimer_test

import (
	"github.com/gogf/gf/g/container/garray"
	"github.com/gogf/gf/g/os/gtimer"
	"github.com/gogf/gf/g/test/gtest"
	"testing"
	"time"
)

func New() *gtimer.Timer {
	return gtimer.New(10, 10*time.Millisecond)
}

func TestTimer_Add_Close(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		//fmt.Println("start", time.Now())
		timer.Add(200*time.Millisecond, func() {
			//fmt.Println("entry1", time.Now())
			array.Append(1)
		})
		timer.Add(200*time.Millisecond, func() {
			//fmt.Println("entry2", time.Now())
			array.Append(1)
		})
		timer.Add(400*time.Millisecond, func() {
			//fmt.Println("entry3", time.Now())
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 5)
		timer.Close()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_Start_Stop_Close(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.Add(200*time.Millisecond, func() {
			//glog.Println("add...")
			array.Append(1)
		})
		gtest.Assert(array.Len(), 0)
		time.Sleep(300 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
		timer.Stop()
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
		timer.Start()
		time.Sleep(200 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
		timer.Close()
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestTimer_AddSingleton(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.AddSingleton(200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestTimer_AddOnce(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.AddOnce(200*time.Millisecond, func() {
			array.Append(1)
		})
		timer.AddOnce(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
		timer.Close()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_AddTimes(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.AddTimes(200*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestTimer_DelayAdd(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.DelayAdd(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddEntry(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.DelayAddEntry(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		}, false, 100, gtimer.STATUS_READY)
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddSingleton(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.DelayAddSingleton(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 0)

		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddOnce(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.DelayAddOnce(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 0)

		time.Sleep(250 * time.Millisecond)
		gtest.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddTimes(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.DelayAddTimes(200*time.Millisecond, 500*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(200 * time.Millisecond)
		gtest.Assert(array.Len(), 0)

		time.Sleep(600 * time.Millisecond)
		gtest.Assert(array.Len(), 1)

		time.Sleep(600 * time.Millisecond)
		gtest.Assert(array.Len(), 2)

		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestTimer_AddLessThanInterval(t *testing.T) {
	gtest.Case(t, func() {
		timer := gtimer.New(10, 100*time.Millisecond)
		array := garray.New()
		timer.Add(20*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(50 * time.Millisecond)
		gtest.Assert(array.Len(), 0)

		time.Sleep(110 * time.Millisecond)
		gtest.Assert(array.Len(), 1)

		time.Sleep(110 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestTimer_AddLeveledEntry1(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		//glog.Println("start")
		timer.DelayAdd(1000*time.Millisecond, 1001*time.Millisecond, func() {
			//glog.Println("add")
			array.Append(1)
		})
		time.Sleep(1500 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(1300 * time.Millisecond)
		//glog.Println("check")
		gtest.Assert(array.Len(), 1)
	})
}

func TestTimer_Exit(t *testing.T) {
	gtest.Case(t, func() {
		timer := New()
		array := garray.New()
		timer.Add(200*time.Millisecond, func() {
			array.Append(1)
			gtimer.Exit()
		})
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}
