# Fastpb

A faster Protobuf serializer & deserializer.

**Attention**

* only proto3 is supported now.
* known-types(any, api, duration...) is not supported now.

## Install

```shell
go install github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
```

## Usage

Refer to [examples][examples], use in two steps:

1. use fastpb to generate code. ([refer here][gen_sh])
2. use fastpb API to marshal/unmarshal. ([refer here][demo])

### Step 1: Generate Code

Using the command line tool to generate code:

```shell
protoc --go_out=. \
  --fastpb_out=. \
  ${your_idl}.proto

or

protoc --go_opt=paths=source_relative --go_out=. \
  --fastpb_opt=paths=source_relative --fastpb_out=. \
   ${your_idl}.proto
```

### Step 2: Codec Message

Encoding and Decoding must use fastpb's API, shown as [demo][demo]:

```go
package examples

import (
	"github.com/cloudwego/fastpb"
)

// Marshal .
func Marshal(msg fastpb.Writer) []byte {
	// TODO: buffer can be reused.
	buf := make([]byte, msg.Size())

	msg.FastWrite(buf)
	return buf
}

// Unmarshal .
func Unmarshal(buf []byte, msg fastpb.Reader) error {
	_, err := fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), msg)
	return err
}
```

## Performance

```
goos: linux
goarch: amd64
pkg: github.com/cloudwego/fastpb/benchmark
cpu: Intel(R) Xeon(R) Gold 5118 CPU @ 2.30GHz
```

Benchmarks have compared [golang/protobuf][golang] (referred to as `_golang`) and [fastpb][fastpb] here.

### Marshal

<table>
<tr><td>Benchmark_marshal_number_golang-48</td><td> 375.2 ns/op </td><td> ~       </td><td>    96 B/op </td><td> ~        </td><td>   1 allocs/op </td></tr>
<tr><td>Benchmark_marshal_number_fastpb-48</td><td> 145.7 ns/op </td><td> -61.17% </td><td>     0 B/op </td><td> -100.00% </td><td>   0 allocs/op </td></tr><tr><td></td></tr>
<tr><td>Benchmark_marshal_string_golang-48</td><td>  1010 ns/op </td><td> ~       </td><td>  2304 B/op </td><td> ~        </td><td>   1 allocs/op </td></tr>
<tr><td>Benchmark_marshal_string_fastpb-48</td><td> 58.57 ns/op </td><td> -94.20% </td><td>     0 B/op </td><td> -100.00% </td><td>   0 allocs/op </td></tr><tr><td></td></tr>
<tr><td>Benchmark_marshal_list_golang-48  </td><td>  8788 ns/op </td><td> ~       </td><td> 18432 B/op </td><td> ~        </td><td>   1 allocs/op </td></tr>
<tr><td>Benchmark_marshal_list_fastpb-48  </td><td>  3430 ns/op </td><td> -60.97% </td><td>     0 B/op </td><td> -100.00% </td><td>   0 allocs/op </td></tr><tr><td></td></tr>
<tr><td>Benchmark_marshal_map_golang-48   </td><td> 43497 ns/op </td><td> ~       </td><td> 21680 B/op </td><td> ~        </td><td> 393 allocs/op </td></tr>
<tr><td>Benchmark_marshal_map_fastpb-48   </td><td>  5951 ns/op </td><td> -86.32% </td><td>     0 B/op </td><td> -100.00% </td><td>   0 allocs/op </td></tr>
</table>

### Unmarshal

<table>
<tr><td>Benchmark_unmarshal_number_golang-48</td><td> 497.1 ns/op </td><td> ~       </td><td>   144 B/op </td><td>   1 allocs/op </td></tr>
<tr><td>Benchmark_unmarshal_number_fastpb-48</td><td> 431.6 ns/op </td><td> -13.18% </td><td>   144 B/op </td><td>   1 allocs/op </td></tr><tr><td></td></tr>
<tr><td>Benchmark_unmarshal_string_golang-48</td><td> 939.7 ns/op </td><td> ~       </td><td>  2128 B/op </td><td>   3 allocs/op </td></tr>
<tr><td>Benchmark_unmarshal_string_fastpb-48</td><td> 668.4 ns/op </td><td> -28.87% </td><td>  2128 B/op </td><td>   3 allocs/op </td></tr><tr><td></td></tr>
<tr><td>Benchmark_unmarshal_list_golang-48  </td><td> 12527 ns/op </td><td> ~       </td><td> 20296 B/op </td><td>  99 allocs/op </td></tr>
<tr><td>Benchmark_unmarshal_list_fastpb-48  </td><td> 12593 ns/op </td><td> +0.53%  </td><td> 20296 B/op </td><td>  99 allocs/op </td></tr><tr><td></td></tr>
<tr><td>Benchmark_unmarshal_map_golang-48   </td><td> 49868 ns/op </td><td> ~       </td><td> 24226 B/op </td><td> 426 allocs/op </td></tr>
<tr><td>Benchmark_unmarshal_map_fastpb-48   </td><td> 21213 ns/op </td><td> -57.46% </td><td> 21467 B/op </td><td>  61 allocs/op </td></tr>
</table>


[fastpb]: https://github.com/cloudwego/fastpb

[examples]: https://github.com/cloudwego/fastpb/tree/main/examples

[demo]: https://github.com/cloudwego/fastpb/blob/main/examples/demo.go

[gen_sh]: https://github.com/cloudwego/fastpb/blob/main/examples/gen_code.sh

[golang]: https://github.com/golang/protobuf
