package event

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/eventx"
)

var SMSObserver = eventx.NewObserver(
	eventx.WithName("sms"),
	eventx.WithHandler(SMSHandler),
)

func SMSHandler(ctx context.Context) error {
	// 发送sms
	return nil
}
