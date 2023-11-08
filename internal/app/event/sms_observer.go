package event

import (
	"context"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/eventx"
)

var SMSObserver = eventx.NewObserver(
	eventx.WithName("sms"),
	eventx.WithHandler(SMSHandler),
)

func SMSHandler(ctx context.Context, param any) error {
	// 发送sms
	v, ok := param.(PayEventParam)
	if !ok {
		return errorpkg.ErrParam
	}

	fmt.Println(v.OrderID)
	return nil
}
