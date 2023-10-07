package rocketpkg

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type Event struct {
	EventName          EventName
	Topic              string   // 消息主题
	Tags               []string // 消息tags
	NameSpace          string
	ConsumerGroupName  string                                             // 所属消费者组
	ConsumerNum        int                                                // 消费者个数
	ConsumerHandleFunc func(context.Context, *primitive.MessageExt) error // 处理消息方法
}

type EventConf map[EventName]Event

type EventName string

// GetEvent 根据事件名称，从事件配置map中获取配置信息(topic 和 tags等)
func GetEvent(ef EventConf, en EventName) (event Event, err error) {
	var ok bool
	if event, ok = ef[en]; ok != true {
		return Event{}, errors.New("no register event")
	}
	return event, nil
}
