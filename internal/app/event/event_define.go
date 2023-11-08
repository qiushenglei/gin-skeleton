package event

import "github.com/qiushenglei/gin-skeleton/pkg/eventx"

const (
	PayEvent = "pay_event"
)

var EventConf = map[string]eventx.EventInferface{
	PayEvent: eventx.NewEvent(
		eventx.WithEventName("email"),
		eventx.WithObserver(EmailObserver, SMSObserver),
	),
}

type PayEventParam struct {
	OrderID string
	UserID  string
}
