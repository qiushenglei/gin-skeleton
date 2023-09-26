package localrocket

import (
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg/products"
)

const (
	OrderEvent rocketpkg.EventName = "OrderEvent"
)

var (
	// EventMap 定义事件
	EventMap = rocketpkg.EventConf{
		OrderEvent: rocketpkg.Event{
			Topic: "order",
			Tags:  []string{"A", "B"},
		},
	}

	// Producer rocketmq生产者
	Producer *products.RocketMQP
)

func RegisterRocketMQProducer() {
	Producer = products.NewProducer(EventMap)
	// 测试发送
	//Producer.SendMsg(context.Background(), OrderEvent, []int{1, 2, 3})
}
