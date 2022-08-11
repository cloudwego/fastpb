// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. ("CloudWeGo Modifications").
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

package fastpb

import (
	"errors"

	"google.golang.org/protobuf/encoding/protowire"
)

// Generic field names and numbers for synthetic map entry messages.
const (
	MapEntry_Key_FieldNumber   = 1  // protoreflect.FieldNumber
	MapEntry_Value_FieldNumber = 2  // protoreflect.FieldNumber
	SkipTagNumber              = -1 // protowire.Number
	SkipTypeCheck              = -1 // protowire.Type
)

// errUnknown is used internally to indicate fields which should be added
// to the unknown field set of a message. It is never returned from an exported
// function.
var errUnknown = errors.New("BUG: internal error (unknown)")

var errDecode = errors.New("cannot parse invalid wire-format data")

var errInvalidUTF8 = errors.New("field contains invalid UTF-8")

// ConsumeTag parses b as a varint-encoded tag, reporting its length.
// This returns a negative length upon an error (see ParseError).
func ConsumeTag(b []byte) (protowire.Number, protowire.Type, int) {
	v, n := ConsumeVarint(b)
	if n < 0 {
		return 0, 0, n // forward error code
	}
	num, typ := protowire.DecodeTag(v)
	if num < protowire.MinValidNumber {
		return 0, 0, -2
	}
	return num, typ, n
}

// AppendTag encodes num and typ as a varint-encoded tag and appends it to b.
func AppendTag(b []byte, num protowire.Number, typ protowire.Type) int {
	return AppendVarint(b, protowire.EncodeTag(num, typ))
}

// ConsumeVarint parses b as a varint-encoded uint64, reporting its length.
// This returns a negative length upon an error (see ParseError).
func ConsumeVarint(b []byte) (v uint64, n int) {
	for i := 0; i < len(b); i++ {
		v |= uint64(b[i]&0x7F) << (i * 7)
		if b[i] < 0x80 {
			return v, i + 1
		}
	}
	return 0, -1
}

// AppendVarint appends v to b as a varint-encoded uint64.
func AppendVarint(buf []byte, v uint64) int {
	switch {
	case v < 1<<7:
		buf[0] = byte(v)
		return 1
	case v < 1<<14:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte(v >> 7)
		return 2
	case v < 1<<21:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte(v >> 14)
		return 3
	case v < 1<<28:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte(v >> 21)
		return 4
	case v < 1<<35:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte((v>>21)&0x7f | 0x80)
		buf[4] = byte(v >> 28)
		return 5
	case v < 1<<42:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte((v>>21)&0x7f | 0x80)
		buf[4] = byte((v>>28)&0x7f | 0x80)
		buf[5] = byte(v >> 35)
		return 6
	case v < 1<<49:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte((v>>21)&0x7f | 0x80)
		buf[4] = byte((v>>28)&0x7f | 0x80)
		buf[5] = byte((v>>35)&0x7f | 0x80)
		buf[6] = byte(v >> 42)
		return 7
	case v < 1<<56:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte((v>>21)&0x7f | 0x80)
		buf[4] = byte((v>>28)&0x7f | 0x80)
		buf[5] = byte((v>>35)&0x7f | 0x80)
		buf[6] = byte((v>>42)&0x7f | 0x80)
		buf[7] = byte(v >> 49)
		return 8
	case v < 1<<63:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte((v>>21)&0x7f | 0x80)
		buf[4] = byte((v>>28)&0x7f | 0x80)
		buf[5] = byte((v>>35)&0x7f | 0x80)
		buf[6] = byte((v>>42)&0x7f | 0x80)
		buf[7] = byte((v>>49)&0x7f | 0x80)
		buf[8] = byte(v >> 56)
		return 9
	default:
		buf[0] = byte((v>>0)&0x7f | 0x80)
		buf[1] = byte((v>>7)&0x7f | 0x80)
		buf[2] = byte((v>>14)&0x7f | 0x80)
		buf[3] = byte((v>>21)&0x7f | 0x80)
		buf[4] = byte((v>>28)&0x7f | 0x80)
		buf[5] = byte((v>>35)&0x7f | 0x80)
		buf[6] = byte((v>>42)&0x7f | 0x80)
		buf[7] = byte((v>>49)&0x7f | 0x80)
		buf[8] = byte((v>>56)&0x7f | 0x80)
		buf[9] = 1
		return 10
	}
}

// AppendFixed32 appends v to b as a little-endian uint32.
func AppendFixed32(b []byte, v uint32) int {
	b[0] = byte(v >> 0)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return 4
}

// AppendFixed64 appends v to b as a little-endian uint64.
func AppendFixed64(b []byte, v uint64) int {
	b[0] = byte(v >> 0)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	return 8
}

// ConsumeBytes parses b as a length-prefixed bytes value, reporting its length.
// This returns a negative length upon an error (see ParseError).
func ConsumeBytes(b []byte) (v []byte, total int) {
	m, n := ConsumeVarint(b)
	total = int(m) + n
	if n < 0 || total > len(b) {
		return nil, -1 // forward error code
	}
	return b[n:total], total
}

// AppendBytes appends v to b as a length-prefixed bytes value.
func AppendBytes(b, v []byte) (n int) {
	n += AppendVarint(b, uint64(len(v)))
	n += copy(b[n:], v)
	return n
}

// AppendString appends v to b as a length-prefixed bytes value.
func AppendString(b []byte, v string) (n int) {
	n += AppendVarint(b, uint64(len(v)))
	n += copy(b[n:], v)
	return n
}

// EnforceUTF8 only support proto3 now.
func EnforceUTF8() bool {
	return true
}
