package event

import "github.com/qiushenglei/gin-skeleton/pkg/eventx"

const (
	OrderEvent = "order_event"
)

type OrderEventParam struct {
}

// 不能像启动mq服务这样配置，这样直接定义EventObj不是线程安全的，这里需要每次触发事件后都重新注册监听者
var Conf = map[string]eventx.EventObj{
	OrderEvent: eventx.EventObj{
		Name: OrderEvent,
		Observers: []*eventx.Observer{
			{
				Name: "email",
				//Handler: ExampleHandler,
			},
			{
				Name: "sms",
				//Handler: ExampleHandler,
			},
		},
	},
}
