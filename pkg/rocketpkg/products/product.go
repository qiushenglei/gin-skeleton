package products

import (
	"context"
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

// RocketMQProducer rocketmq的生产者
type RocketMQProducer struct {
	rocketmq.Producer                     // rocketmq官方包的生产者
	ef                rocketpkg.EventConf // 私有rocket的事件配置map
}

// NewPkgProducer 创建生产者
func NewPkgProducer(em rocketpkg.EventConf) *RocketMQProducer {
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

	return newRocketMQProducer(RocketProduct, em)
}

// SendMsg 发送消息
func (p *RocketMQProducer) SendMsg(ctx context.Context, en rocketpkg.EventName, data any) error {
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
