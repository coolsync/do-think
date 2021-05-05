# RPC



## Install

1.  protobuf



## 

## Example CPU

pcbook/proto/processer_message.proto

```protobuf
syntax = "proto3";

message CPU {
    string brand = 1;
    string name = 2;
    uint32 number_cores = 3;
    uint32 number_threads = 4;
    double min_ghz = 5;
    double max_ghz = 6;
}
```



## gen go file



下面已失效：

```bash
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```



报错：

```bash
please specify it with the full Go package path as
a future release of protoc-gen-go will require this be specified.
See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```



```protobuf
protoc --go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
	proto/processer_message.proto
```


正确做法：

--proto_path=PATH

```bash
protoc --go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
	proto/*.proto
```

```bash
option go_package = "./pb;pcbook";	// 设置输出目录 和 当前的包名

protoc --go_out=. proto/*.proto

or:
protoc --proto_path=proto --go_out=. proto/*.proto

or:
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

```bash
protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
```



脚本执行：

Makefile

1. gen
2. clean
3. run

```makefile
gen:
	protoc --go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

clean:
	rm pb/proto/*.go

run:
	go run main.go
```















