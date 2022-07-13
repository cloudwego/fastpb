# FastPB

* **Attention**
    - only support proto3 now
    - any is not supported now.

# Install
go >= 1.16
```shell
go install github.com/cloudwego/fastpb/protoc-gen-fastpb
```
go <= 1.15
```shell
go get github.com/cloudwego/fastpb/protoc-gen-fastpb
```

# gen code
```shell
protoc --go_out=. --fastpb_out=. ${your_idl}.proto
```