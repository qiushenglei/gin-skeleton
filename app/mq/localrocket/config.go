package localrocket

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg/products"
)

const (
	OrderEvent rocketpkg.EventName = "OrderEvent"
)

var (
	// EventMap 定义事件
	EventMap rocketpkg.EventConf = rocketpkg.EventConf{
		OrderEvent: rocketpkg.Event{
			Topic: "order",
		},
	}

	// Producer rocketmq生产者
	Producer *products.RocketMQP
)

func RegisterRocketMQProducer() {
	Producer = products.NewProducer(EventMap)
	Producer.SendMsg(context.Background(), OrderEvent, []int{1, 2, 3})
}
