package eventx

import (
	"context"
	"sync"
)

type Event interface {
	Watch() error
	Notice() error
}

type EventOption func(obj *EventObj)

// 事件对象
type EventObj struct {
	Name      string
	Observers []*Observer
	ctx       context.Context
	Event
}

func NewEvent(...EventOption) *EventObj {

	return &EventObj{}
}

func WithContext(ctx context.Context) EventOption {
	return func(e *EventObj) {
		e.ctx = ctx
	}
}

func WithEventName(name string) EventOption {
	return func(e *EventObj) {
		e.Name = name
	}
}

func WithObserver(observers ...*Observer) EventOption {
	return func(e *EventObj) {
		e.Observers = observers
		return
	}
}

func (e *EventObj) Watch() error {
	return nil
}

// 事件通知观察者
func (e *EventObj) Notice() error {
	wg := sync.WaitGroup{}
	c, _ := context.WithCancel(e.ctx)
	for _, v := range e.Observers {
		wg.Add(1)
		go func() {
			wg.Done()
			v.Handler(c)
		}()
	}
	return nil
}
