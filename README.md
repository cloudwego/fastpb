# FastPB

A faster Protobuf serializer & deserializer.

**Attention**

* Only proto3 is supported now
* Any is not supported now.

## Install

go >= 1.16

```shell
go install github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
```

go <= 1.15

```shell
go get github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
```

## Generate Code

```shell
protoc --go_out=. --fastpb_out=. ${your_idl}.proto
```
