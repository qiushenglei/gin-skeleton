package products

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/qiushenglei/gin-skeleton/app/configs"
	"github.com/qiushenglei/gin-skeleton/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"go.uber.org/zap"
)

type RocketMQP struct {
	rocketmq.Producer
	em rocketpkg.EventConf
}

// NewProducer 创建生产者
func NewProducer(em rocketpkg.EventConf) *RocketMQP {
	host := fmt.Sprintf("%s:%s", configs.EnvConfig.GetString("ROCKETMQ_TEST_HOST"), configs.EnvConfig.GetString("ROCKETMQ_TEST_PORT"))
	var err error
	RocketProduct, err := rocketmq.NewProducer(
		producer.WithNsResolver(
			primitive.NewPassthroughResolver(
				[]string{
					host,
				})),
		producer.WithRetry(5),
		producer.WithSendMsgTimeout(10),
	)
	if err != nil {
		panic(err)
	}
	err = RocketProduct.Start()
	if err != nil {
		panic(err)
	}

	return newRocketMQP(RocketProduct, em)
}

// SendMsg 发送消息
func (p *RocketMQP) SendMsg(ctx context.Context, en rocketpkg.EventName, data any) error {
	defer utils.DefaultDefer(ctx)
	// 创建消息
	msg := p.NewMessage(ctx, en, data)

	// 异步发送
	err := p.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, err error) {
		logs.Log.Info(ctx, "async send msg success")
	}, msg)

	if err != nil {
		logs.Log.Info(ctx, zap.String("sendAsync fail", err.Error()))
	}

	return err
}

func (p *RocketMQP) NewMessage(ctx context.Context, en rocketpkg.EventName, data any) *primitive.Message {
	// 获取event的配置信息
	eventInfo, err := p.getEvent(en)
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

func newRocketMQP(p rocketmq.Producer, ef rocketpkg.EventConf) *RocketMQP {
	return &RocketMQP{
		p,
		ef,
	}
}

func (p *RocketMQP) getEvent(e rocketpkg.EventName) (event rocketpkg.Event, err error) {
	var ok bool
	if event, ok = p.em[e]; ok != true {
		return rocketpkg.Event{}, errors.New("no register event")
	}
	return event, nil
}
