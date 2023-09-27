package localrocket

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
)

// OrderPayEventHandle 订单支付事件处理
func OrderPayEventHandle(ctx context.Context, msg *primitive.MessageExt) error {
	fmt.Println(string(msg.Body))
	fmt.Printf("msg: %+v", msg)
	logs.Log.Info(ctx, string(msg.Body))
	return nil
}
