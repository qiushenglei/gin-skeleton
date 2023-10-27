package grpc

import (
	"github.com/qiushenglei/gin-skeleton/proto"
)

type OrderStream struct {
	proto.UnimplementedOrderStreamServerServer
}

func (o *OrderStream) FindAll(f proto.OrderStreamServer_FindAllServer) error {
	return nil
}
func (o *OrderStream) Insert(i proto.OrderStreamServer_InsertServer) error {
	res, err := i.Recv()
	if err != nil {
		return err
	}
	s := &proto.InsertStreamReply{
		AppId: res.AppId,
	}
	i.Send(s)
	return err
}
