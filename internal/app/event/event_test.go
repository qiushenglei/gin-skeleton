package event

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/eventx"
	"testing"
)

func TestEvent(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	eventx.NewEvent(
		eventx.WithEventName(OrderEvent),
		eventx.WithContext(ctx),
		eventx.WithObserver(SMSObserver, EmailObserver),
	).Notice(ctx, OrderEventParam{})
}
