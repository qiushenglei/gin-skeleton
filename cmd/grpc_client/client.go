package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var action string

func main() {
	flag.StringVar(&action, "mod", "stream", "create or turnon")
	flag.Parse()
	conn, err := grpc.Dial("127.0.0.1:10012", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	if action == "stream" {
		client := proto.NewOrderStreamServerClient(conn)
		r := &proto.InsertStreamRequest{
			AppId: "abdsfg",
		}

		insertClient, err := client.Insert(context.Background())
		if err != nil {
			fmt.Println("create insert client fail")
			return
		}
		if err := insertClient.Send(r); err != nil {
			fmt.Println("send fail")
			return
		}

		for {
			res, err := insertClient.Recv()
			if err != nil {
				fmt.Println("recv error")
			}
			fmt.Println(res)
			time.Sleep(1 * time.Second)
		}
		return
	} else {

		client := proto.NewOrderServerClient(conn)
		request := &proto.FindAllRequest{
			AppId:    "bcdefa",
			Page:     1,
			PageSize: 10,
		}
		res, err := client.FindAll(context.Background(), request)
		if err != nil {
			//panic(err)
		}
		fmt.Println(res)
	}
}
