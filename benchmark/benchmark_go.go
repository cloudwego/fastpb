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

	"google.golang.org/protobuf/proto"

	go_gen "github.com/cloudwego/fastpb/benchmark/fastpb_gen"
)

var (
	_goNumber *go_gen.Number
	_goString *go_gen.String
	_goList   *go_gen.List
	_goMap    *go_gen.Map
)

func goNumber() *go_gen.Number {
	if _goNumber != nil {
		return _goNumber
	}
	_goNumber = &go_gen.Number{
		Field1:  rand.Int31(),
		Field2:  rand.Int63(),
		Field3:  rand.Uint32(),
		Field4:  rand.Uint64(),
		Field5:  rand.Int31(),
		Field6:  rand.Int63(),
		Field7:  true,
		Field8:  rand.Float32(),
		Field9:  rand.Float64(),
		Field10: rand.Uint32(),
		Field11: rand.Uint64(),
		Field12: rand.Int31(),
		Field13: rand.Int63(),
	}
	return _goNumber
}

func goString() *go_gen.String {
	if _goString != nil {
		return _goString
	}
	_goString = &go_gen.String{
		Field1: string(make([]byte, MockLen)),
		Field2: make([]byte, MockLen),
	}
	return _goString
}

func goList() *go_gen.List {
	if _goList != nil {
		return _goList
	}
	_goList = &go_gen.List{
		Field1:  make([]int32, MockCap),
		Field2:  make([]int64, MockCap),
		Field3:  make([]uint32, MockCap),
		Field4:  make([]uint64, MockCap),
		Field5:  make([]int32, MockCap),
		Field6:  make([]int64, MockCap),
		Field7:  make([]bool, MockCap),
		Field8:  make([]float32, MockCap),
		Field9:  make([]float64, MockCap),
		Field10: make([]uint32, MockCap),
		Field11: make([]uint64, MockCap),
		Field12: make([]int32, MockCap),
		Field13: make([]int64, MockCap),
		Field14: make([]string, MockCap),
		Field15: make([][]byte, MockCap),
	}
	randListInt32(_goList.Field1)
	randListInt64(_goList.Field2)
	randListUint32(_goList.Field3)
	randListUint64(_goList.Field4)
	randListInt32(_goList.Field5)
	randListInt64(_goList.Field6)
	randListBool(_goList.Field7)
	randListFloat32(_goList.Field8)
	randListFloat64(_goList.Field9)
	randListUint32(_goList.Field10)
	randListUint64(_goList.Field11)
	randListInt32(_goList.Field12)
	randListInt64(_goList.Field13)
	randListString(_goList.Field14)
	randListBytes(_goList.Field15)
	return _goList
}

func goMap() *go_gen.Map {
	if _goMap != nil {
		return _goMap
	}
	_goMap = &go_gen.Map{
		Field1: make(map[int32]int64, MockCap),
		Field2: make(map[uint32]uint64, MockCap),
		Field3: make(map[int32]int64, MockCap),
		Field4: make(map[uint32]uint64, MockCap),
		Field5: make(map[int32]int64, MockCap),
		Field6: make(map[string][]byte, MockCap),
	}
	randMapInt32Int64(_goMap.Field1)
	randMapUint32Uint64(_goMap.Field2)
	randMapInt32Int64(_goMap.Field3)
	randMapUint32Uint64(_goMap.Field4)
	randMapInt32Int64(_goMap.Field5)
	randMapStringBytes(_goMap.Field6)
	return _goMap
}

var (
	_bytesNumber []byte
	_bytesString []byte
	_bytesList   []byte
	_bytesMap    []byte
)

func bytesNumber() []byte {
	if _bytesNumber != nil {
		return _bytesNumber
	}
	_bytesNumber, _ = proto.Marshal(goNumber())
	return _bytesNumber
}

func bytesString() []byte {
	if _bytesString != nil {
		return _bytesString
	}
	_bytesString, _ = proto.Marshal(goString())
	return _bytesString
}

func bytesList() []byte {
	if _bytesList != nil {
		return _bytesList
	}
	_bytesList, _ = proto.Marshal(goList())
	return _bytesList
}

func bytesMap() []byte {
	if _bytesMap != nil {
		return _bytesMap
	}
	_bytesMap, _ = proto.Marshal(goMap())
	return _bytesMap
}
