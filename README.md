# Simple gRPC by Golang

## Author
- Nguyễn Văn Hải - imhai248
- Email: nvhai248@gmail.com

## Instructions

Follow the link to read the documentation: [gRPC Gateway Introduction](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/)

Alternatively, you can use the following commands to install the necessary tools:

```
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Make sure you have `protoc` installed. You can follow the installation instructions here: [Install Protocol Buffers](https://grpc.io/docs/protoc-installation/). After installing `protoc`, verify the installation with `protoc --version`.

To generate `stubs`, you can use the following commands based on your operating system:

1. Mac or linux:

```
protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   ./proto/demo.proto
```

2. Windows:

```
protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative ./proto/demo.proto
```

After setting up the project and installing the necessary dependencies, you can use `go mod tidy` to import all required modules and libraries.