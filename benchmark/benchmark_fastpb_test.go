/*
 * Copyright 2022 CloudWeGo Authors
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

package benchmark

import (
	"testing"

	"github.com/cloudwego/fastpb"

	fastpb_gen "github.com/cloudwego/fastpb/benchmark/fastpb_gen"
)

func Benchmark_marshal_number_fastpb(b *testing.B) {
	encode := fastpbNumber()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encode.FastWrite(buf)
	}
}

func Benchmark_unmarshal_number_fastpb(b *testing.B) {
	buf := bytesNumber()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode fastpb_gen.Number
		_, _ = fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), &decode)
	}
}

func Benchmark_marshal_string_fastpb(b *testing.B) {
	encode := fastpbString()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encode.FastWrite(buf)
	}
}

func Benchmark_unmarshal_string_fastpb(b *testing.B) {
	buf := bytesString()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode fastpb_gen.String
		_, _ = fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), &decode)
	}
}

func Benchmark_marshal_list_fastpb(b *testing.B) {
	encode := fastpbList()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encode.FastWrite(buf)
	}
}

func Benchmark_unmarshal_list_fastpb(b *testing.B) {
	buf := bytesList()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode fastpb_gen.List
		_, _ = fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), &decode)
	}
}

func Benchmark_marshal_map_fastpb(b *testing.B) {
	encode := fastpbMap()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encode.FastWrite(buf)
	}
}

func Benchmark_unmarshal_map_fastpb(b *testing.B) {
	buf := bytesMap()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode fastpb_gen.Map
		_, _ = fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), &decode)
	}
}
