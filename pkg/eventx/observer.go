package eventx

import "context"

// 通知事件
type handler func(c context.Context, param any) error

type ObserverOptionFunc func(*ObserverOption)

type ObserverOption struct {
	Name    string
	Handler handler
}

// 观察者
type Observer struct {
	ObserverOption
}

func NewObserver(options ...ObserverOptionFunc) *Observer {
	option := &ObserverOption{}
	for _, apply := range options {
		apply(option)
	}

	return &Observer{
		ObserverOption: *option,
	}
}

func WithName(name string) ObserverOptionFunc {
	return func(o *ObserverOption) {
		o.Name = name
	}
}

func WithHandler(h handler) ObserverOptionFunc {
	return func(o *ObserverOption) {
		o.Handler = h
	}
}
