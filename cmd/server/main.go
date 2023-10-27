package main

import "github.com/qiushenglei/gin-skeleton/cmd"

//go:generate protoc --proto_path=../../proto/ --go_out=. --go-grpc_out=.  --go-grpc_opt=paths=source_relative  orderstream.proto

func main() {
	//example.AATest()

	//启动命令行服务
	cmd.CmdExecute()
}
