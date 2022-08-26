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
)

func init() {
	Init()
}

func Init() {
	_ = goNumber()
	_ = goString()
	_ = goList()
	_ = goMap()

	_ = fastpbNumber()
	_ = fastpbString()
	_ = fastpbList()
	_ = fastpbMap()

	_ = gogoNumber()
	_ = gogoString()
	_ = gogoList()
	_ = gogoMap()

	_ = bytesNumber()
	_ = bytesString()
	_ = bytesList()
	_ = bytesMap()
}

func BenchmarkMain(b *testing.B) {
	Init()

	// marshal ----------------------------------------------------------------------
	b.Run("Benchmark_marshal_number_golang", Benchmark_marshal_number_golang)
	b.Run("Benchmark_marshal_number_fastpb", Benchmark_marshal_number_fastpb)
	b.Run("Benchmark_marshal_number_gogo", Benchmark_marshal_number_gogo)
	println()

	b.Run("Benchmark_marshal_string_golang", Benchmark_marshal_string_golang)
	b.Run("Benchmark_marshal_string_fastpb", Benchmark_marshal_string_fastpb)
	b.Run("Benchmark_marshal_string_gogo", Benchmark_marshal_string_gogo)
	println()

	b.Run("Benchmark_marshal_list_golang", Benchmark_marshal_list_golang)
	b.Run("Benchmark_marshal_list_fastpb", Benchmark_marshal_list_fastpb)
	b.Run("Benchmark_marshal_list_gogo", Benchmark_marshal_list_gogo)
	println()

	b.Run("Benchmark_marshal_map_golang", Benchmark_marshal_map_golang)
	b.Run("Benchmark_marshal_map_fastpb", Benchmark_marshal_map_fastpb)
	b.Run("Benchmark_marshal_map_gogo", Benchmark_marshal_map_gogo)
	println()

	// unmarshal ----------------------------------------------------------------------
	b.Run("Benchmark_unmarshal_number_golang", Benchmark_unmarshal_number_golang)
	b.Run("Benchmark_unmarshal_number_fastpb", Benchmark_unmarshal_number_fastpb)
	b.Run("Benchmark_unmarshal_number_gogo", Benchmark_unmarshal_number_gogo)
	println()

	b.Run("Benchmark_unmarshal_string_golang", Benchmark_unmarshal_string_golang)
	b.Run("Benchmark_unmarshal_string_fastpb", Benchmark_unmarshal_string_fastpb)
	b.Run("Benchmark_unmarshal_string_gogo", Benchmark_unmarshal_string_gogo)
	println()

	b.Run("Benchmark_unmarshal_list_golang", Benchmark_unmarshal_list_golang)
	b.Run("Benchmark_unmarshal_list_fastpb", Benchmark_unmarshal_list_fastpb)
	b.Run("Benchmark_unmarshal_list_gogo", Benchmark_unmarshal_list_gogo)
	println()

	b.Run("Benchmark_unmarshal_map_golang", Benchmark_unmarshal_map_golang)
	b.Run("Benchmark_unmarshal_map_fastpb", Benchmark_unmarshal_map_fastpb)
	b.Run("Benchmark_unmarshal_map_gogo", Benchmark_unmarshal_map_gogo)
	println()
}
