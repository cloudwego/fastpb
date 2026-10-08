/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fastpb

import (
	"google.golang.org/protobuf/encoding/protowire"
	"testing"
)

// A valid 64-bit varint is at most 10 bytes; the 10th byte can only
// contribute bit 63 (value 0 or 1).  The reference protowire decoder
// rejects anything larger as an overflow (errCodeOverflow == -3).
// Previously fastpb's ConsumeVarint had no length cap and no 10th-byte
// check, so over-long or over-large varints were silently accepted with
// corrupted values (shift-by >= 64 wraps and ORs stray bits).
func TestConsumeVarintRejectsOverflow(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
	}{
		{"10th_byte_value_2", []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x02}},
		{"10th_byte_value_127", []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x7f}},
		{"11th_byte_continuation", []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, rn := protowire.ConsumeVarint(tc.in)
			if rn >= 0 {
				t.Fatalf("reference protowire accepted overflow input (n=%d); test is stale", rn)
			}
			_, n := ConsumeVarint(tc.in)
			if n >= 0 {
				t.Fatalf("ConsumeVarint accepted overflow varint: n=%d, want negative", n)
			}
		})
	}
}

// The maximum valid uint64 varint (10 bytes, 10th byte == 1) must decode
// identically to the reference.
func TestConsumeVarintMaxUint64(t *testing.T) {
	in := []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}
	rv, rn := protowire.ConsumeVarint(in)
	fv, fn := ConsumeVarint(in)
	if fn != rn || fv != rv {
		t.Fatalf("ConsumeVarint = (v=%d,n=%d), ref = (v=%d,n=%d)", fv, fn, rv, rn)
	}
}
