// Copyright 2017 gf Author(https://github.com/zhongdalu/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhongdalu/gf.

package gparser_test

import (
	"github.com/zhongdalu/gf/g"
	"github.com/zhongdalu/gf/g/encoding/gparser"
	"github.com/zhongdalu/gf/g/test/gtest"
	"testing"
)

func Test_New(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		v := j.Value().(g.Map)
		gtest.Assert(v["n"], 123456789)
	})
}

func Test_NewUnsafe(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.NewUnsafe(data)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a.1"), 2)
	})
}

func Test_Encode(t *testing.T) {
	value := g.Slice{1, 2, 3}
	gtest.Case(t, func() {
		b, err := gparser.VarToJson(value)
		gtest.Assert(err, nil)
		gtest.Assert(b, []byte(`[1,2,3]`))
	})
}

func Test_Decode(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a.1"), 2)
	})
}

func Test_SplitChar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		j.SetSplitChar(byte('#'))
		gtest.AssertNE(j, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m#k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a#1"), 2)
	})
}

func Test_ViolenceCheck(t *testing.T) {
	data := []byte(`{"m":{"a":[1,2,3], "v1.v2":"4"}}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.Assert(j.Get("m.a.2"), 3)
		gtest.Assert(j.Get("m.v1.v2"), nil)
		j.SetViolenceCheck(true)
		gtest.Assert(j.Get("m.v1.v2"), 4)
	})
}

func Test_GetToVar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		var m map[string]string
		var n int
		var a []int
		j := gparser.New(data)
		gtest.AssertNE(j, nil)

		j.GetToVar("n", &n)
		j.GetToVar("m", &m)
		j.GetToVar("a", &a)

		gtest.Assert(n, "123456789")
		gtest.Assert(m, g.Map{"k": "v"})
		gtest.Assert(a, g.Slice{1, 2, 3})
	})
}

func Test_GetMap(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.Assert(j.GetMap("n"), nil)
		gtest.Assert(j.GetMap("m"), g.Map{"k": "v"})
		gtest.Assert(j.GetMap("a"), nil)
	})
}

func Test_GetArray(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.Assert(j.GetArray("n"), g.Array{123456789})
		gtest.Assert(j.GetArray("m"), g.Array{g.Map{"k": "v"}})
		gtest.Assert(j.GetArray("a"), g.Array{1, 2, 3})
	})
}

func Test_GetString(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.AssertEQ(j.GetString("n"), "123456789")
		gtest.AssertEQ(j.GetString("m"), `{"k":"v"}`)
		gtest.AssertEQ(j.GetString("a"), `[1,2,3]`)
		gtest.AssertEQ(j.GetString("i"), "")
	})
}

func Test_GetStrings(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.AssertEQ(j.GetStrings("n"), g.SliceStr{"123456789"})
		gtest.AssertEQ(j.GetStrings("m"), g.SliceStr{`{"k":"v"}`})
		gtest.AssertEQ(j.GetStrings("a"), g.SliceStr{"1", "2", "3"})
		gtest.AssertEQ(j.GetStrings("i"), nil)
	})
}

func Test_GetInterfaces(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.Case(t, func() {
		j := gparser.New(data)
		gtest.AssertNE(j, nil)
		gtest.AssertEQ(j.GetInterfaces("n"), g.Array{123456789})
		gtest.AssertEQ(j.GetInterfaces("m"), g.Array{g.Map{"k": "v"}})
		gtest.AssertEQ(j.GetInterfaces("a"), g.Array{1, 2, 3})
	})
}

func Test_Len(t *testing.T) {
	gtest.Case(t, func() {
		p := gparser.New(nil)
		p.Append("a", 1)
		p.Append("a", 2)
		gtest.Assert(p.Len("a"), 2)
	})
	gtest.Case(t, func() {
		p := gparser.New(nil)
		p.Append("a.b", 1)
		p.Append("a.c", 2)
		gtest.Assert(p.Len("a"), 2)
	})
	gtest.Case(t, func() {
		p := gparser.New(nil)
		p.Set("a", 1)
		gtest.Assert(p.Len("a"), -1)
	})
}

func Test_Append(t *testing.T) {
	gtest.Case(t, func() {
		p := gparser.New(nil)
		p.Append("a", 1)
		p.Append("a", 2)
		gtest.Assert(p.Get("a"), g.Slice{1, 2})
	})
	gtest.Case(t, func() {
		p := gparser.New(nil)
		p.Append("a.b", 1)
		p.Append("a.c", 2)
		gtest.Assert(p.Get("a"), g.Map{
			"b": g.Slice{1},
			"c": g.Slice{2},
		})
	})
	gtest.Case(t, func() {
		p := gparser.New(nil)
		p.Set("a", 1)
		err := p.Append("a", 2)
		gtest.AssertNE(err, nil)
		gtest.Assert(p.Get("a"), 1)
	})
}

func Test_Convert(t *testing.T) {
	gtest.Case(t, func() {
		p := gparser.New(`{"name":"gf","bool":true,"int":1,"float":1,"ints":[1,2],"floats":[1,2],"time":"2019-06-12","person": {"name": "gf"}}`)
		gtest.Assert(p.GetVar("name").String(), "gf")
		gtest.Assert(p.GetString("name"), "gf")
		gtest.Assert(p.GetBool("bool"), true)
		gtest.Assert(p.GetInt("int"), 1)
		gtest.Assert(p.GetInt8("int"), 1)
		gtest.Assert(p.GetInt16("int"), 1)
		gtest.Assert(p.GetInt32("int"), 1)
		gtest.Assert(p.GetInt64("int"), 1)
		gtest.Assert(p.GetUint("int"), 1)
		gtest.Assert(p.GetUint8("int"), 1)
		gtest.Assert(p.GetUint16("int"), 1)
		gtest.Assert(p.GetUint32("int"), 1)
		gtest.Assert(p.GetUint64("int"), 1)
		gtest.Assert(p.GetInts("ints")[0], 1)
		gtest.Assert(p.GetFloat32("float"), 1)
		gtest.Assert(p.GetFloat64("float"), 1)
		gtest.Assert(p.GetFloats("floats")[0], 1)
		gtest.Assert(p.GetTime("time").Format("2006-01-02"), "2019-06-12")
		gtest.Assert(p.GetGTime("time").Format("Y-m-d"), "2019-06-12")
		gtest.Assert(p.GetDuration("time").String(), "0s")
		name := struct {
			Name string
		}{}
		err := p.GetToStruct("person", &name)
		gtest.Assert(err, nil)
		gtest.Assert(name.Name, "gf")
		gtest.Assert(p.ToMap()["name"], "gf")
		err = p.ToStruct(&name)
		gtest.Assert(err, nil)
		gtest.Assert(name.Name, "gf")
		p.Dump()

		p = gparser.New(`[0,1,2]`)
		gtest.Assert(p.ToArray()[0], 0)

		err = gparser.VarToStruct(`{"name":"gf"}`, &name)
		gtest.Assert(err, nil)
		gtest.Assert(name.Name, "gf")
	})
}

func Test_Convert2(t *testing.T) {
	gtest.Case(t, func() {
		xmlArr := []byte{60, 114, 111, 111, 116, 47, 62}
		p := gparser.New(`<root></root>`)
		arr, err := p.ToXml("root")
		gtest.Assert(err, nil)
		gtest.Assert(arr, xmlArr)
		arr, err = gparser.VarToXml(`<root></root>`, "root")
		gtest.Assert(err, nil)
		gtest.Assert(arr, xmlArr)

		arr, err = p.ToXmlIndent("root")
		gtest.Assert(err, nil)
		gtest.Assert(arr, xmlArr)
		arr, err = gparser.VarToXmlIndent(`<root></root>`, "root")
		gtest.Assert(err, nil)
		gtest.Assert(arr, xmlArr)

		p = gparser.New(`{"name":"gf"}`)
		str, err := p.ToJsonString()
		gtest.Assert(err, nil)
		gtest.Assert(str, `{"name":"gf"}`)
		str, err = gparser.VarToJsonString(`{"name":"gf"}`)
		gtest.Assert(err, nil)
		gtest.Assert(str, `{"name":"gf"}`)

		jsonIndentArr := []byte{123, 10, 9, 34, 110, 97, 109, 101, 34, 58, 32, 34, 103, 102, 34, 10, 125}
		arr, err = p.ToJsonIndent()
		gtest.Assert(err, nil)
		gtest.Assert(arr, jsonIndentArr)
		arr, err = gparser.VarToJsonIndent(`{"name":"gf"}`)
		gtest.Assert(err, nil)
		gtest.Assert(arr, jsonIndentArr)

		str, err = p.ToJsonIndentString()
		gtest.Assert(err, nil)
		gtest.Assert(str, "{\n\t\"name\": \"gf\"\n}")
		str, err = gparser.VarToJsonIndentString(`{"name":"gf"}`)
		gtest.Assert(err, nil)
		gtest.Assert(str, "{\n\t\"name\": \"gf\"\n}")

		yamlArr := []byte{124, 50, 10, 10, 32, 32, 110, 97, 109, 101, 58, 103, 102, 10}
		p = gparser.New(`
name:gf
`)
		arr, err = p.ToYaml()
		gtest.Assert(err, nil)
		gtest.Assert(arr, yamlArr)
		arr, err = gparser.VarToYaml(`
name:gf
`)
		gtest.Assert(err, nil)
		gtest.Assert(arr, yamlArr)

		tomlArr := []byte{110, 97, 109, 101, 32, 61, 32, 34, 103, 102, 34, 10}
		p = gparser.New(`
name= "gf"
`)
		arr, err = p.ToToml()
		gtest.Assert(err, nil)
		gtest.Assert(arr, tomlArr)
		arr, err = gparser.VarToToml(`
name= "gf"
`)
		gtest.Assert(err, nil)
		gtest.Assert(arr, tomlArr)
	})
}
