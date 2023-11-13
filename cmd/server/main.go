package main

import (
	"github.com/qiushenglei/gin-skeleton/cmd"
)

//go:generate go run ../gorm_gen/.
//go:generate protoc --proto_path=../../proto/ --go_out=../../proto/  --go_opt=paths=source_relative  --go-grpc_out=../../proto/  --go-grpc_opt=paths=source_relative  orderstream.proto

func main() {
	//example.AATest()

	// pprof

	//启动命令行服务
	cmd.CmdExecute()
}
