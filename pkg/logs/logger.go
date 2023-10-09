package logs

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Log *Logger

type Logger struct {
	zapLogger *zap.SugaredLogger
}

func NewLogger(zapLogger *zap.SugaredLogger) *Logger {
	l := &Logger{
		zapLogger,
	}
	return l
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	args = l.addRequestID(ctx, args)
	l.zapLogger.Error(args...)
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	args = l.addRequestID(ctx, args)
	l.zapLogger.Info(11)
}

func (l *Logger) Warn(ctx context.Context, args ...interface{}) {
	args = l.addRequestID(ctx, args)
	l.zapLogger.Warn(args...)
}

func retStringRequestID(ctx context.Context) string {
	ginCtx, ok := ctx.(*gin.Context)
	if ok == false {
		return ""
	}
	return ginCtx.GetString("RequestID")
}

func (l *Logger) addRequestID(ctx context.Context, arg []interface{}) []interface{} {
	rid := retStringRequestID(ctx)
	arg = append(arg, zap.String("requestID", rid))
	return arg
}
