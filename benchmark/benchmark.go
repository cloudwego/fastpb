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
	"math/rand"
)

const (
	MockLen = 1024
	MockCap = 16
)

func randListInt32(list []int32) {
	for i := range list {
		list[i] = rand.Int31()
	}
}

func randListInt64(list []int64) {
	for i := range list {
		list[i] = rand.Int63()
	}
}

func randListUint32(list []uint32) {
	for i := range list {
		list[i] = rand.Uint32()
	}
}

func randListUint64(list []uint64) {
	for i := range list {
		list[i] = rand.Uint64()
	}
}

func randListBool(list []bool) {
	for i := range list {
		list[i] = rand.Intn(2) == 1
	}
}

func randListFloat32(list []float32) {
	for i := range list {
		list[i] = rand.Float32()
	}
}

func randListFloat64(list []float64) {
	for i := range list {
		list[i] = rand.Float64()
	}
}

func randListString(list []string) {
	for i := range list {
		list[i] = string(make([]byte, rand.Intn(MockLen)))
	}
}

func randListBytes(list [][]byte) {
	for i := range list {
		list[i] = make([]byte, rand.Intn(MockLen))
	}
}

func randMapInt32Int64(m map[int32]int64) {
	for i := 0; i < MockCap; i++ {
		m[rand.Int31()] = rand.Int63()
	}
}

func randMapUint32Uint64(m map[uint32]uint64) {
	for i := 0; i < MockCap; i++ {
		m[rand.Uint32()] = rand.Uint64()
	}
}

func randMapStringBytes(m map[string][]byte) {
	for i := 0; i < MockCap; i++ {
		m[string(make([]byte, rand.Intn(MockLen)))] = make([]byte, rand.Intn(MockLen))
	}
}
