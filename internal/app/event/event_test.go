package event

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/eventx"
	"testing"
)

func TestEvent(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	eventx.NewEvent(
		eventx.WithEventName(PayEvent),
		eventx.WithObserver(SMSObserver, EmailObserver),
	).Notice(ctx, PayEventParam{})
}

func TestEvent1(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	param := PayEventParam{
		OrderID: "abcdefg",
		UserID:  "user",
	}
	EventConf[PayEvent].Notice(ctx, param)
}
