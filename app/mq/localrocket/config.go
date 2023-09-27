package localrocket

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg/consumers"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg/products"
)

const (
	// 订阅事件名
	OrderPayEvent     rocketpkg.EventName = "OrderPayEvent"     // 订单支付事件
	OrderSuccessEvent rocketpkg.EventName = "OrderSuccessEvent" // 订单完成事件

	// 主题topic
	OrderTopic string = "order"

	// 消费者组名
	OrderPayGroup string = "OrderPayGroup"
)

var (
	// EventMap 定义事件
	EventMap = rocketpkg.EventConf{
		OrderPayEvent: rocketpkg.Event{
			EventName:          OrderPayEvent,
			Topic:              OrderTopic,
			Tags:               []string{"A", "B"},
			ConsumerNum:        4,
			ConsumerGroupName:  OrderPayGroup,
			ConsumerHandleFunc: OrderPayEventHandle,
		},
	}

	// Producer rocketmq生产者
	Producer *products.RocketMQProducer
)

func RegisterRocketMQProducer() error {
	Producer = products.NewPkgProducer(EventMap)
	// 测试发送
	err := Producer.SendMsg(context.Background(), OrderPayEvent, []int{1, 2, 3})

	return err
}

func RegisterRocketMQConsumer() {
	err := consumers.NewConsumers(EventMap).Start()
	if err != nil {
		panic(err.Error())
	}
}
