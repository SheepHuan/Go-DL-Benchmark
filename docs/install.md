```bash


# 初始化go mod
go mod init go-dl-benchmark_v1
go mod tidy

```


```bash
# proto编译
protoc --go_out=. .\remote_terminal.proto
protoc --go-grpc_out=. .\remote_terminal.proto

```


```bash
#  benchmark_v1 for docker
docker pull ubuntu:18.04
docker run -it -v ~/code:/workspace -p 
# 安装 anconda环境


# 安装 go 环境

# [挂载USB设备到docker](https://hlyani.github.io/notes/docker/mount_usb_to_docker.html)

```



### 编译protos
```bash

protoc --go_out=pkg/protos pkg/protos/modelBenchmarkTestArgs.proto
protoc --go_out=pkg/protos pkg/protos/modelBenchmarkTestResult.proto
protoc --go_out=pkg/protos pkg/protos/physicalDeviceDescription.proto

# js 编译
# https://www.jianshu.com/p/00a6a03dff80
protoc --js_out=import_style=commonjs,binary:. modelBenchmarkTestArgs.proto modelBenchmarkTestResult.proto physicalDeviceDescription.proto
```