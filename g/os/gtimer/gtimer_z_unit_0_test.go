// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package functions

package gtimer_test

import (
	"github.com/gogf/gf/g/container/garray"
	"github.com/gogf/gf/g/os/gtimer"
	"github.com/gogf/gf/g/test/gtest"
	"testing"
	"time"
)

func TestSetTimeout(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.SetTimeout(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestSetInterval(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.SetInterval(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(1100 * time.Millisecond)
		gtest.Assert(array.Len(), 5)
	})
}

func TestAddEntry(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.AddEntry(200*time.Millisecond, func() {
			array.Append(1)
		}, false, 2, gtimer.STATUS_READY)
		time.Sleep(1100 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestAddSingleton(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.AddSingleton(200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(1100 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestAddTimes(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.AddTimes(200*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestDelayAdd(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.DelayAdd(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(200 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestDelayAddEntry(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.DelayAddEntry(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		}, false, 2, gtimer.STATUS_READY)
		time.Sleep(300 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}

func TestDelayAddSingleton(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.DelayAddSingleton(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(300 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestDelayAddOnce(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.DelayAddOnce(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 1)
	})
}

func TestDelayAddTimes(t *testing.T) {
	gtest.Case(t, func() {
		array := garray.New()
		gtimer.DelayAddTimes(200*time.Millisecond, 200*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		gtest.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		gtest.Assert(array.Len(), 2)
	})
}
