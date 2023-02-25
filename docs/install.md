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


```bash
#  benchmark for docker
docker pull ubuntu:18.04
docker run -it -v ~/code:/workspace -p 
# 安装 anconda环境


# 安装 go 环境

# [挂载USB设备到docker](https://hlyani.github.io/notes/docker/mount_usb_to_docker.html)

```