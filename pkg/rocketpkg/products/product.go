package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/qiushenglei/gin-skeleton/app/configs"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"go.uber.org/zap"
)

type RocketMQP struct {
	rocketmq.Producer
	em rocketpkg.EventConf
}

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

func (p *RocketMQP) SendMsg(ctx context.Context, en rocketpkg.EventName, data any) error {

	topic := p.getEvent(en).Topic
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	msg := primitive.NewMessage(topic, b)

	err = p.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, err error) {
		logs.Log.Info(ctx, "async send msg success")
	}, msg)

	if err != nil {
		logs.Log.Info(ctx, zap.String("sendAsync fail", err.Error()))
	}

	return err
}

func newRocketMQP(p rocketmq.Producer, ef rocketpkg.EventConf) *RocketMQP {
	return &RocketMQP{
		p,
		ef,
	}
}

func (p *RocketMQP) getEvent(e rocketpkg.EventName) (event rocketpkg.Event) {
	var ok bool
	if event, ok = p.em[e]; ok != true {
		panic("undefine")
	}
	return event
}
