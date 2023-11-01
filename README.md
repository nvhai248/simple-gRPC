# SIMPLE GRPC BY GOLANG
## AUTHOR
Nguyễn Văn Hải - imhai248
nvhai248@gmail.com
## INSTRUCTIONS

Flow the link to read the documentation: https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/

Or you can flow me:

```
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

You need install protoc, follow the url to install it: https://grpc.io/docs/protoc-installation/
After installing protoc, try `protoc --version` to check the latest version

Run this cmd to generating stubs:

Mac or linux:

```
protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   ./proto/demo.proto
```

Win:

```
protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative ./proto/demo.proto
```

Using `go mod tidy` to import all modules and libraries.