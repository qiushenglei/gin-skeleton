package utils

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"go.uber.org/zap"
)

// Go 封装的go程，不用每次都手写panic
func Go(ctx context.Context, handle func(ctx context.Context), rh func(r interface{})) {
	//p := func() {
	//	if r := recover(); r != nil {
	//		if rh == nil {
	//		}
	//		Go(ctx, func(_ context.Context) {
	//			rh(r)
	//		}, nil)
	//	}
	//}

	go func() {
		defer DefaultDefer(ctx)
		handle(ctx)
	}()
}

// DefaultDefer 接收panic的值类型是error
func DefaultDefer(ctx context.Context) {
	if r := recover(); r != nil {
		if v, ok := r.(error); ok == true {
			logs.Log.Error(ctx, zap.String("recover err", v.Error()))
		}
	}
}

func DefaultRh(r interface{}) {

}
