#!/bin/bash

# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 获取脚本文件的相对路径
relative_path="$0"

# 获取脚本文件的绝对路径
absolute_path=$(readlink -f "$relative_path")
current_dir=$(dirname "$absolute_path")
echo "cd $current_dir"
cd $current_dir

mkdir -p pb/vpsproxy
protoc --proto_path=src --go_out=pb/vpsproxy --go-grpc_out=pb/vpsproxy --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  vpsproxy.proto
