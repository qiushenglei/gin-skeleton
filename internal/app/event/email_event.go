package event

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/eventx"
)

var EmailObserver = eventx.NewObserver(
	eventx.WithName("email"),
	eventx.WithHandler(EmailHandler),
)

func EmailHandler(ctx context.Context, param any) error {
	// 发送email
	return nil
}
