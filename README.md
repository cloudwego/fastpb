# FastPB

A faster Protobuf serializer & deserializer.

**Attention**

* only proto3 is supported now.
* known-types(any, api, duration...) is not supported now.

## Install

go >= 1.16

```shell
go install github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
```

go <= 1.15

```shell
go get github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
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


[examples]: https://github.com/cloudwego/fastpb/tree/main/examples
[demo]: https://github.com/cloudwego/fastpb/blob/main/examples/demo.go
[gen_sh]: https://github.com/cloudwego/fastpb/blob/main/examples/gen_code.sh
