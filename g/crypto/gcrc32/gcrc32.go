// Copyright 2017 gf Author(https://github.com/zhongdalu/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhongdalu/gf.

// Package gcrc32 provides useful API for CRC32 encryption algorithms.
package gcrc32

import (
	"github.com/zhongdalu/gf/g/util/gconv"
	"hash/crc32"
)

// Encrypt encrypts any type of variable using CRC32 algorithms.
// It uses gconv package to convert <v> to its bytes type.
func Encrypt(v interface{}) uint32 {
	return crc32.ChecksumIEEE(gconv.Bytes(v))
}

// Deprecated.
func EncryptString(v string) uint32 {
	return crc32.ChecksumIEEE([]byte(v))
}

// Deprecated.
func EncryptBytes(v []byte) uint32 {
	return crc32.ChecksumIEEE(v)
}
