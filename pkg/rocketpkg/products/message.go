package products

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
)

// newRocketMQProducer 创建私有包rocketmq的生产者
func newRocketMQProducer(p rocketmq.Producer, ef rocketpkg.EventConf) *RocketMQProducer {
	return &RocketMQProducer{
		p,
		ef,
	}
}

// getEvent 根据事件名称，从事件配置map中获取配置信息(topic 和 tags等)
func (p *RocketMQProducer) getEvent(e rocketpkg.EventName) (event rocketpkg.Event, err error) {
	var ok bool
	if event, ok = p.ef[e]; ok != true {
		return rocketpkg.Event{}, errors.New("no register event")
	}
	return event, nil
}

// NewMessage 从事件配置map中获取配置，创建rocket消息并指定topic 和 tags
func (p *RocketMQProducer) NewMessage(ctx context.Context, en rocketpkg.EventName, data any) *primitive.Message {
	// 获取event的配置信息
	eventInfo, err := rocketpkg.GetEvent(p.ef, en)
	if err != nil {
		panic(err)
	}

	// encode结构体信息
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// 创建Msg
	msg := primitive.NewMessage(eventInfo.Topic, b)

	// 添加tag
	for _, v := range eventInfo.Tags {
		msg.WithTag(v)
	}

	return msg
}
