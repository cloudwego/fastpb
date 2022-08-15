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

	gogo_gen "github.com/cloudwego/fastpb/benchmark/gogo_gen"
)

func Benchmark_marshal_number_gogo(b *testing.B) {
	encode := gogoNumber()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encode.MarshalTo(buf)
	}
}

func Benchmark_unmarshal_number_gogo(b *testing.B) {
	buf := bytesNumber()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode gogo_gen.Number
		_ = decode.Unmarshal(buf)
	}
}

func Benchmark_marshal_string_gogo(b *testing.B) {
	encode := gogoString()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encode.MarshalTo(buf)
	}
}

func Benchmark_unmarshal_string_gogo(b *testing.B) {
	buf := bytesString()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode gogo_gen.String
		_ = decode.Unmarshal(buf)
	}
}

func Benchmark_marshal_list_gogo(b *testing.B) {
	encode := gogoList()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encode.MarshalTo(buf)
	}
}

func Benchmark_unmarshal_list_gogo(b *testing.B) {
	buf := bytesList()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode gogo_gen.List
		_ = decode.Unmarshal(buf)
	}
}

func Benchmark_marshal_map_gogo(b *testing.B) {
	encode := gogoMap()
	buf := make([]byte, encode.Size())

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encode.MarshalTo(buf)
	}
}

func Benchmark_unmarshal_map_gogo(b *testing.B) {
	buf := bytesMap()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode gogo_gen.Map
		_ = decode.Unmarshal(buf)
	}
}
