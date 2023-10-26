package grpc

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/constants"
	"github.com/qiushenglei/gin-skeleton/internal/app/services"
	"github.com/qiushenglei/gin-skeleton/pkg/validatorx"
	"github.com/qiushenglei/gin-skeleton/proto"
)

type OrderServer struct {
	proto.UnimplementedOrderServerServer
}

func (o *OrderServer) FindAll(c context.Context, r *proto.FindAllRequest) (*proto.FindAllReply, error) {
	if err := validatorx.ValidateRPC(r); err != nil {
		return nil, err
	}
	page := int(r.Page)
	pageSize := int(r.GetPageSize())
	request := &entity.FindOrderRequest{
		AppID:    r.AppId,
		Page:     &page,
		PageSize: &pageSize,
	}

	data, count, err := services.FindAll(c, request)

	var list []*proto.OrderData
	for _, v := range data {
		list = append(list, &proto.OrderData{
			Id:         v.ID,
			OrderId:    v.OrderID,
			AppId:      v.AppID,
			AddTime:    v.AddTime.Format(constants.DateFormat),
			UpdateTime: v.AddTime.Format(constants.DateFormat),
		})
	}

	res := &proto.FindAllReply{
		List:  list,
		Count: int32(count),
	}
	return res, err
}
