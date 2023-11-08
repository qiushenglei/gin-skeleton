package eventx

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"sync"
)

type EventInferface interface {
	Watch() error
	Notice(ctx context.Context, any2 any) error
}

type EventOptionFunc func(obj *EventOption)

type EventOption struct {
	Name      string
	Observers []*Observer
}

// 事件对象
type Event struct {
	options *EventOption
	EventInferface
}

func NewEvent(op ...EventOptionFunc) EventInferface {
	option := &EventOption{}
	for _, apply := range op {
		apply(option)
	}

	o := &Event{
		options: option,
	}
	return o
}

func WithEventName(name string) EventOptionFunc {
	return func(e *EventOption) {
		e.Name = name
	}
}

func WithObserver(observers ...*Observer) EventOptionFunc {
	return func(e *EventOption) {
		e.Observers = observers
		return
	}
}

func (e *Event) Watch() error {
	return nil
}

// 事件通知观察者
func (e *Event) Notice(ctx context.Context, param any) error {
	wg := sync.WaitGroup{}
	c, _ := context.WithCancel(ctx)
	for _, v := range e.options.Observers {
		wg.Add(1)
		tmp := v
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					logs.Log.Error(c, r)
				}
			}()
			if err := tmp.Handler(c, param); err != nil {
				// TODO::do something
			}
		}()
	}
	wg.Wait()
	return nil
}
