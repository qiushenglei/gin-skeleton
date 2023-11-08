package event

import (
	"context"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/eventx"
)

var EmailObserver = eventx.NewObserver(
	eventx.WithName("email"),
	eventx.WithHandler(EmailHandler),
)

func EmailHandler(ctx context.Context, param any) error {
	// 发送email
	v, ok := param.(PayEventParam)
	if !ok {
		return errorpkg.ErrParam
	}

	fmt.Println(v.UserID)
	return nil
}
