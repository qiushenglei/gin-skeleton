package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/qiushenglei/gin-skeleton/internal/app/mq/localrocket"
)

// RocketProduct rocketmq消费者
var RocketProduct rocketmq.Producer

func RegisterMQ(serverName string) error {
	isRegister := configs.EnvConfig.GetInt("REGISTER_MQ")
	if isRegister == 0 {
		return nil
	}

	if serverName == "rocketmq" {
		return nil
	}

	err := localrocket.RegisterRocketMQProducer()
	return err
}
