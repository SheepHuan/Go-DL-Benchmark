```bash


# 初始化go mod
go mod init go-dl-benchmark
go mod tidy

```


```bash
# proto编译
protoc --go_out=. .\remote_terminal.proto
protoc --go-grpc_out=. .\remote_terminal.proto

```