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

package fastpb

// Reader is designed for generating code.
type Reader interface {
	FastRead(buf []byte, _type int8, number int32) (n int, err error)
}

// Writer is designed for generating code.
type Writer interface {
	Sizer
	FastWrite(buf []byte) (n int)
}

// Sizer is designed for generating code.
type Sizer interface {
	Size() (n int)
}

// Protocol .
type Protocol interface {
	WriteMessage(buf []byte, number int32, writer Writer) (n int)
	WriteListPacked(buf []byte, number int32, length int, single Marshal) (n int)
	WriteMapEntry(buf []byte, number int32, entry Marshal) (n int)
	WriteBool(buf []byte, number int32, value bool) (n int)
	WriteInt32(buf []byte, number, value int32) (n int)
	WriteInt64(buf []byte, number int32, value int64) (n int)
	WriteUint32(buf []byte, number int32, value uint32) (n int)
	WriteUint64(buf []byte, number int32, value uint64) (n int)
	WriteSint32(buf []byte, number, value int32) (n int)
	WriteSint64(buf []byte, number int32, value int64) (n int)
	WriteFloat(buf []byte, number int32, value float32) (n int)
	WriteDouble(buf []byte, number int32, value float64) (n int)
	WriteFixed32(buf []byte, number int32, value uint32) (n int)
	WriteFixed64(buf []byte, number int32, value uint64) (n int)
	WriteSfixed32(buf []byte, number, value int32) (n int)
	WriteSfixed64(buf []byte, number int32, value int64) (n int)
	WriteString(buf []byte, number int32, value string) (n int)
	WriteBytes(buf []byte, number int32, value []byte) (n int)
	ReadMessage(buf []byte, _type int8, reader Reader) (offset int, err error)
	ReadList(buf []byte, _type int8, single Unmarshal) (n int, err error)
	ReadMapEntry(buf []byte, _type int8, umk, umv Unmarshal) (int, error)
	ReadBool(buf []byte, _type int8) (value bool, n int, err error)
	ReadInt32(buf []byte, _type int8) (value int32, n int, err error)
	ReadInt64(buf []byte, _type int8) (value int64, n int, err error)
	ReadUint32(buf []byte, _type int8) (value uint32, n int, err error)
	ReadUint64(buf []byte, _type int8) (value uint64, n int, err error)
	ReadSint32(buf []byte, _type int8) (value int32, n int, err error)
	ReadSint64(buf []byte, _type int8) (value int64, n int, err error)
	ReadFloat(buf []byte, _type int8) (value float32, n int, err error)
	ReadDouble(buf []byte, _type int8) (value float64, n int, err error)
	ReadFixed32(buf []byte, _type int8) (value uint32, n int, err error)
	ReadFixed64(buf []byte, _type int8) (value uint64, n int, err error)
	ReadSfixed32(buf []byte, _type int8) (value int32, n int, err error)
	ReadSfixed64(buf []byte, _type int8) (value int64, n int, err error)
	ReadString(buf []byte, _type int8) (value string, n int, err error)
	ReadBytes(buf []byte, _type int8) (value []byte, n int, err error)
	SizeBool(number int32, value bool) (n int)
	SizeInt32(number, value int32) (n int)
	SizeInt64(number int32, value int64) (n int)
	SizeUint32(number int32, value uint32) (n int)
	SizeUint64(number int32, value uint64) (n int)
	SizeSint32(number, value int32) (n int)
	SizeSint64(number int32, value int64) (n int)
	SizeFloat(number int32, value float32) (n int)
	SizeDouble(number int32, value float64) (n int)
	SizeFixed32(number int32, value uint32) (n int)
	SizeFixed64(number int32, value uint64) (n int)
	SizeSfixed32(number, value int32) (n int)
	SizeSfixed64(number int32, value int64) (n int)
	SizeString(number int32, value string) (n int)
	SizeBytes(number int32, value []byte) (n int)
	SizeMessage(number int32, sizer Sizer) (n int)
	SizeMapEntry(number int32, entry EntrySize) (n int)
	SizeListPacked(number int32, length int, single EntrySize) (n int)
	Skip(buf []byte, _type int8, number int32) (n int, err error)
}

// Marshal .
type Marshal func(buf []byte, numTagOrKey, numIdxOrVal int32) int

// Unmarshal .
type Unmarshal func(buf []byte, _type int8) (n int, err error)

// EntrySize .
type EntrySize func(numTagOrKey, numIdxOrVal int32) int

// WriteMessage .
func WriteMessage(buf []byte, number int32, writer Writer) (n int) {
	return Impl.WriteMessage(buf, number, writer)
}

// WriteListPacked .
func WriteListPacked(buf []byte, number int32, length int, single Marshal) (n int) {
	return Impl.WriteListPacked(buf, number, length, single)
}

// WriteMapEntry .
func WriteMapEntry(buf []byte, number int32, entry Marshal) (n int) {
	return Impl.WriteMapEntry(buf, number, entry)
}

// WriteBool .
func WriteBool(buf []byte, number int32, value bool) (n int) {
	return Impl.WriteBool(buf, number, value)
}

// WriteInt32 .
func WriteInt32(buf []byte, number, value int32) (n int) {
	return Impl.WriteInt32(buf, number, value)
}

// WriteInt64 .
func WriteInt64(buf []byte, number int32, value int64) (n int) {
	return Impl.WriteInt64(buf, number, value)
}

// WriteUint32 .
func WriteUint32(buf []byte, number int32, value uint32) (n int) {
	return Impl.WriteUint32(buf, number, value)
}

// WriteUint64 .
func WriteUint64(buf []byte, number int32, value uint64) (n int) {
	return Impl.WriteUint64(buf, number, value)
}

// WriteSint32 .
func WriteSint32(buf []byte, number, value int32) (n int) {
	return Impl.WriteSint32(buf, number, value)
}

// WriteSint64 .
func WriteSint64(buf []byte, number int32, value int64) (n int) {
	return Impl.WriteSint64(buf, number, value)
}

// WriteFloat .
func WriteFloat(buf []byte, number int32, value float32) (n int) {
	return Impl.WriteFloat(buf, number, value)
}

// WriteDouble .
func WriteDouble(buf []byte, number int32, value float64) (n int) {
	return Impl.WriteDouble(buf, number, value)
}

// WriteFixed32 .
func WriteFixed32(buf []byte, number int32, value uint32) (n int) {
	return Impl.WriteFixed32(buf, number, value)
}

// WriteFixed64 .
func WriteFixed64(buf []byte, number int32, value uint64) (n int) {
	return Impl.WriteFixed64(buf, number, value)
}

// WriteSfixed32 .
func WriteSfixed32(buf []byte, number, value int32) (n int) {
	return Impl.WriteSfixed32(buf, number, value)
}

// WriteSfixed64 .
func WriteSfixed64(buf []byte, number int32, value int64) (n int) {
	return Impl.WriteSfixed64(buf, number, value)
}

// WriteString .
func WriteString(buf []byte, number int32, value string) (n int) {
	return Impl.WriteString(buf, number, value)
}

// WriteBytes .
func WriteBytes(buf []byte, number int32, value []byte) (n int) {
	return Impl.WriteBytes(buf, number, value)
}

// ReadMessage .
func ReadMessage(buf []byte, _type int8, reader Reader) (offset int, err error) {
	return Impl.ReadMessage(buf, _type, reader)
}

// ReadList .
func ReadList(buf []byte, _type int8, single Unmarshal) (n int, err error) {
	return Impl.ReadList(buf, _type, single)
}

// ReadMapEntry .
func ReadMapEntry(buf []byte, _type int8, umk, umv Unmarshal) (int, error) {
	return Impl.ReadMapEntry(buf, _type, umk, umv)
}

// ReadBool .
func ReadBool(buf []byte, _type int8) (value bool, n int, err error) {
	return Impl.ReadBool(buf, _type)
}

// ReadInt32 .
func ReadInt32(buf []byte, _type int8) (value int32, n int, err error) {
	return Impl.ReadInt32(buf, _type)
}

// ReadInt64 .
func ReadInt64(buf []byte, _type int8) (value int64, n int, err error) {
	return Impl.ReadInt64(buf, _type)
}

// ReadUint32 .
func ReadUint32(buf []byte, _type int8) (value uint32, n int, err error) {
	return Impl.ReadUint32(buf, _type)
}

// ReadUint64 .
func ReadUint64(buf []byte, _type int8) (value uint64, n int, err error) {
	return Impl.ReadUint64(buf, _type)
}

// ReadSint32 .
func ReadSint32(buf []byte, _type int8) (value int32, n int, err error) {
	return Impl.ReadSint32(buf, _type)
}

// ReadSint64 .
func ReadSint64(buf []byte, _type int8) (value int64, n int, err error) {
	return Impl.ReadSint64(buf, _type)
}

// ReadFloat .
func ReadFloat(buf []byte, _type int8) (value float32, n int, err error) {
	return Impl.ReadFloat(buf, _type)
}

// ReadDouble .
func ReadDouble(buf []byte, _type int8) (value float64, n int, err error) {
	return Impl.ReadDouble(buf, _type)
}

// ReadFixed32 .
func ReadFixed32(buf []byte, _type int8) (value uint32, n int, err error) {
	return Impl.ReadFixed32(buf, _type)
}

// ReadFixed64 .
func ReadFixed64(buf []byte, _type int8) (value uint64, n int, err error) {
	return Impl.ReadFixed64(buf, _type)
}

// ReadSfixed32 .
func ReadSfixed32(buf []byte, _type int8) (value int32, n int, err error) {
	return Impl.ReadSfixed32(buf, _type)
}

// ReadSfixed64 .
func ReadSfixed64(buf []byte, _type int8) (value int64, n int, err error) {
	return Impl.ReadSfixed64(buf, _type)
}

// ReadString .
func ReadString(buf []byte, _type int8) (value string, n int, err error) {
	return Impl.ReadString(buf, _type)
}

// ReadBytes .
func ReadBytes(buf []byte, _type int8) (value []byte, n int, err error) {
	return Impl.ReadBytes(buf, _type)
}

// SizeBool .
func SizeBool(number int32, value bool) (n int) {
	return Impl.SizeBool(number, value)
}

// SizeInt32 .
func SizeInt32(number, value int32) (n int) {
	return Impl.SizeInt32(number, value)
}

// SizeInt64 .
func SizeInt64(number int32, value int64) (n int) {
	return Impl.SizeInt64(number, value)
}

// SizeUint32 .
func SizeUint32(number int32, value uint32) (n int) {
	return Impl.SizeUint32(number, value)
}

// SizeUint64 .
func SizeUint64(number int32, value uint64) (n int) {
	return Impl.SizeUint64(number, value)
}

// SizeSint32 .
func SizeSint32(number, value int32) (n int) {
	return Impl.SizeSint32(number, value)
}

// SizeSint64 .
func SizeSint64(number int32, value int64) (n int) {
	return Impl.SizeSint64(number, value)
}

// SizeFloat .
func SizeFloat(number int32, value float32) (n int) {
	return Impl.SizeFloat(number, value)
}

// SizeDouble .
func SizeDouble(number int32, value float64) (n int) {
	return Impl.SizeDouble(number, value)
}

// SizeFixed32 .
func SizeFixed32(number int32, value uint32) (n int) {
	return Impl.SizeFixed32(number, value)
}

// SizeFixed64 .
func SizeFixed64(number int32, value uint64) (n int) {
	return Impl.SizeFixed64(number, value)
}

// SizeSfixed32 .
func SizeSfixed32(number, value int32) (n int) {
	return Impl.SizeSfixed32(number, value)
}

// SizeSfixed64 .
func SizeSfixed64(number int32, value int64) (n int) {
	return Impl.SizeSfixed64(number, value)
}

// SizeString .
func SizeString(number int32, value string) (n int) {
	return Impl.SizeString(number, value)
}

// SizeBytes .
func SizeBytes(number int32, value []byte) (n int) {
	return Impl.SizeBytes(number, value)
}

// SizeMessage .
func SizeMessage(number int32, sizer Sizer) (n int) {
	return Impl.SizeMessage(number, sizer)
}

// SizeMapEntry .
func SizeMapEntry(number int32, entry EntrySize) (n int) {
	return Impl.SizeMapEntry(number, entry)
}

// SizeListPacked .
func SizeListPacked(number int32, length int, single EntrySize) (n int) {
	return Impl.SizeListPacked(number, length, single)
}

// Skip .
func Skip(buf []byte, _type int8, number int32) (n int, err error) {
	return Impl.Skip(buf, _type, number)
}
