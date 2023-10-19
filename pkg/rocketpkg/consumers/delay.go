package consumers

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"strconv"
)

func (c *RocketMQConsumers) redeliver(ctx context.Context, e rocketpkg.Event, msg *primitive.MessageExt) error {
	expire := msg.GetProperty(rocketpkg.Expire)
	num, _ := strconv.Atoi(expire)
	err := c.productMap["default"].SyncSendMsg(ctx, e.EventName, msg.Body, &rocketpkg.Delay{Expire: num})
	return err
}
