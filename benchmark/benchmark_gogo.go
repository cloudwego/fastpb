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
	gogo_gen "github.com/cloudwego/fastpb/benchmark/gogo_gen"
)

var (
	gogo_number *gogo_gen.Number
	gogo_string *gogo_gen.String
	gogo_list   *gogo_gen.List
	gogo_map    *gogo_gen.Map
)

func gogoNumber() *gogo_gen.Number {
	if gogo_number != nil {
		return gogo_number
	}
	gogo_number = &gogo_gen.Number{
		Field1:  _goNumber.Field1,
		Field2:  _goNumber.Field2,
		Field3:  _goNumber.Field3,
		Field4:  _goNumber.Field4,
		Field5:  _goNumber.Field5,
		Field6:  _goNumber.Field6,
		Field7:  _goNumber.Field7,
		Field8:  _goNumber.Field8,
		Field9:  _goNumber.Field9,
		Field10: _goNumber.Field10,
		Field11: _goNumber.Field11,
		Field12: _goNumber.Field12,
		Field13: _goNumber.Field13,
	}
	return gogo_number
}

func gogoString() *gogo_gen.String {
	if gogo_string != nil {
		return gogo_string
	}
	gogo_string = &gogo_gen.String{
		Field1: _goString.Field1,
		Field2: _goString.Field2,
	}
	return gogo_string
}

func gogoList() *gogo_gen.List {
	if gogo_list != nil {
		return gogo_list
	}
	gogo_list = &gogo_gen.List{
		Field1:  _goList.Field1,
		Field2:  _goList.Field2,
		Field3:  _goList.Field3,
		Field4:  _goList.Field4,
		Field5:  _goList.Field5,
		Field6:  _goList.Field6,
		Field7:  _goList.Field7,
		Field8:  _goList.Field8,
		Field9:  _goList.Field9,
		Field10: _goList.Field10,
		Field11: _goList.Field11,
		Field12: _goList.Field12,
		Field13: _goList.Field13,
		Field14: _goList.Field14,
		Field15: _goList.Field15,
	}
	return gogo_list
}

func gogoMap() *gogo_gen.Map {
	if gogo_map != nil {
		return gogo_map
	}
	gogo_map = &gogo_gen.Map{
		Field1: _goMap.Field1,
		Field2: _goMap.Field2,
		Field3: _goMap.Field3,
		Field4: _goMap.Field4,
		Field5: _goMap.Field5,
		Field6: _goMap.Field6,
	}
	return gogo_map
}
