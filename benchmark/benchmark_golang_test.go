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

	"google.golang.org/protobuf/proto"

	go_gen "github.com/cloudwego/fastpb/benchmark/fastpb_gen"
)

func Benchmark_marshal_number_golang(b *testing.B) {
	encode := goNumber()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(encode)
	}
}

func Benchmark_unmarshal_number_golang(b *testing.B) {
	buf := bytesNumber()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode go_gen.Number
		_ = proto.Unmarshal(buf, &decode)
	}
}

func Benchmark_marshal_string_golang(b *testing.B) {
	encode := goString()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(encode)
	}
}

func Benchmark_unmarshal_string_golang(b *testing.B) {
	buf := bytesString()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode go_gen.String
		_ = proto.Unmarshal(buf, &decode)
	}
}

func Benchmark_marshal_list_golang(b *testing.B) {
	encode := goList()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(encode)
	}
}

func Benchmark_unmarshal_list_golang(b *testing.B) {
	buf := bytesList()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode go_gen.List
		_ = proto.Unmarshal(buf, &decode)
	}
}

func Benchmark_marshal_map_golang(b *testing.B) {
	encode := goMap()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(encode)
	}
}

func Benchmark_unmarshal_map_golang(b *testing.B) {
	buf := bytesMap()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decode go_gen.Map
		_ = proto.Unmarshal(buf, &decode)
	}
}
