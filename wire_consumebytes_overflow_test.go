// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fastpb

import (
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
)

// A length-delimited field whose declared length is encoded as a 10-byte
// varint near MaxUint64 used to overflow the `int(m) + n` total: int(m) wraps
// negative, the upper-bound check `total > len(b)` passed, and ConsumeBytes then
// sliced `b[n:total]` with start > end, panicking with
// "slice bounds out of range". The reference decoder returns errCodeTruncated.
// This is fuzz-derived (arbitrary wire bytes) and must return a clean error.
func TestConsumeBytesRejectsDeclaredLengthOverflow(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
	}{
		{"max_uint64_length_trailing", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 'X', 'Y'}},
		{"max_uint64_no_trailing", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}},
		{"just_under_max_length", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 'X'}},
		{"legal_2pow63_varint", []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}},
		{"huge_length_zero_top_byte", []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x00, 'a', 'b'}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("ConsumeBytes panicked on %x: %v", tc.in, r)
				}
			}()
			rv, rn := protowire.ConsumeBytes(tc.in)
			fv, fn := ConsumeBytes(tc.in)
			// Error signal must match the reference.
			if (fn < 0) != (rn < 0) {
				t.Fatalf("error-signal mismatch: fastpb n=%d, ref n=%d (in=%x)", fn, rn, tc.in)
			}
			// On success the payload must match the reference byte-for-byte.
			if fn >= 0 {
				if fn != rn || string(fv) != string(rv) {
					t.Fatalf("payload mismatch: fastpb=(%q,n=%d) ref=(%q,n=%d) (in=%x)", fv, fn, rv, rn, tc.in)
				}
			}
		})
	}
}

// A normal small length-delimited value must still decode identically.
func TestConsumeBytesValidUnchanged(t *testing.T) {
	in := []byte{0x03, 'a', 'b', 'c', 't', 'r', 'a', 'i', 'l'}
	v, n := ConsumeBytes(in)
	if n != 4 || string(v) != "abc" {
		t.Fatalf("v=%q n=%d want abc/4", string(v), n)
	}
	rv, rn := protowire.ConsumeBytes(in)
	if n != rn || string(v) != string(rv) {
		t.Fatalf("diverge from ref: fastpb=(%q,%d) ref=(%q,%d)", v, n, rv, rn)
	}
}
