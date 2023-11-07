package eventx

import "context"

// 通知事件
type handler func(c context.Context, param any) error

type ObserverOption func(*Observer)

// 观察者
type Observer struct {
	Name    string
	Handler handler
}

func NewObserver(...ObserverOption) *Observer {
	return &Observer{}
}

func WithName(name string) ObserverOption {
	return func(o *Observer) {
		o.Name = name
	}
}

func WithHandler(h handler) ObserverOption {
	return func(o *Observer) {
		o.Handler = h
	}
}
