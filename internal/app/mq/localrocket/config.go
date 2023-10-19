package localrocket

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg/consumers"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg/products"
)

const (
	// 订阅事件名
	DBToESEvent           rocketpkg.EventName = "DBToESEvent"
	OrderPayEvent         rocketpkg.EventName = "OrderPayEvent"     // 订单支付事件
	OrderAutoSuccessEvent rocketpkg.EventName = "OrderSuccessEvent" // 订单自动完成

	// 主题topic
	OrderTopic      string = "order"
	DBToESTopic     string = "dbtoes1"
	OrderDelayTopic string = "order_delay"

	// 消费者组名
	OrderPayGroup         string = "OrderPayGroup"
	CanalSyncESGroup      string = "CanalSyncESGroup1"
	OrderAutoSuccessGroup string = "OrderAutoSuccessGroup"

	// topic的enamespace，不同namespace下可以同名topic
	SyncDataNameSpace string = "dbtoes1"
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
		DBToESEvent: rocketpkg.Event{
			EventName:          DBToESEvent,
			Topic:              DBToESTopic,
			Tags:               []string{},
			ConsumerNameSpace:  SyncDataNameSpace,
			ConsumerNum:        4,
			ConsumerGroupName:  CanalSyncESGroup,
			ConsumerHandleFunc: CanalSyncESHandle,
		},
		OrderAutoSuccessEvent: rocketpkg.Event{
			EventName:          OrderAutoSuccessEvent,
			Topic:              OrderDelayTopic,
			Tags:               []string{"AutoSuccess"},
			ProducerGroup:      string(OrderAutoSuccessEvent),
			ConsumerNum:        4,
			ConsumerGroupName:  OrderAutoSuccessGroup,
			ConsumerHandleFunc: DelayHandle,
		},
	}

	// Producer rocketmq生产者
	Producer *products.RocketMQProducer
)

func RegisterRocketMQProducer() error {
	Producer = products.NewPkgProducer(EventMap, "productAction")
	// 测试普通消息发送
	var err error
	for i := 0; i < 1; i++ {
		err = Producer.SyncSendMsg(context.Background(), OrderPayEvent, "这里是测试", nil)
		err = Producer.SyncSendMsg(context.Background(), OrderAutoSuccessEvent, "这里是延时测试", &rocketpkg.Delay{Expire: 60})
	}
	return err
}

func RegisterRocketMQConsumer() {
	err := consumers.NewConsumers(EventMap).Start()
	if err != nil {
		panic(err.Error())
	}
}
