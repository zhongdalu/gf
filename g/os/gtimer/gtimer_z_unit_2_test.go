// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Entry Operations

package gtimer_test

import (
	"github.com/gogf/gf/g/container/garray"
	"github.com/gogf/gf/g/os/gtimer"
	"github.com/gogf/gf/g/test/gtest"
	"testing"
	"time"
)

func TestEntry_Start_Stop_Close(t *testing.T) {
	timer := New()
	array := garray.New()
	entry := timer.Add(200*time.Millisecond, func() {
		array.Append(1)
	})
	time.Sleep(250 * time.Millisecond)
	gtest.Assert(array.Len(), 1)
	entry.Stop()
	time.Sleep(250 * time.Millisecond)
	gtest.Assert(array.Len(), 1)
	entry.Start()
	time.Sleep(250 * time.Millisecond)
	gtest.Assert(array.Len(), 2)
	entry.Close()
	time.Sleep(250 * time.Millisecond)
	gtest.Assert(array.Len(), 2)

	gtest.Assert(entry.Status(), gtimer.STATUS_CLOSED)
}

func TestEntry_Singleton(t *testing.T) {
	timer := New()
	array := garray.New()
	entry := timer.Add(200*time.Millisecond, func() {
		array.Append(1)
		time.Sleep(10 * time.Second)
	})
	gtest.Assert(entry.IsSingleton(), false)
	entry.SetSingleton(true)
	gtest.Assert(entry.IsSingleton(), true)
	time.Sleep(250 * time.Millisecond)
	gtest.Assert(array.Len(), 1)

	time.Sleep(250 * time.Millisecond)
	gtest.Assert(array.Len(), 1)
}

func TestEntry_SetTimes(t *testing.T) {
	timer := New()
	array := garray.New()
	entry := timer.Add(200*time.Millisecond, func() {
		array.Append(1)
	})
	entry.SetTimes(2)
	time.Sleep(1200 * time.Millisecond)
	gtest.Assert(array.Len(), 2)
}

func TestEntry_Run(t *testing.T) {
	timer := New()
	array := garray.New()
	entry := timer.Add(1000*time.Millisecond, func() {
		array.Append(1)
	})
	entry.Run()
	gtest.Assert(array.Len(), 1)
}
