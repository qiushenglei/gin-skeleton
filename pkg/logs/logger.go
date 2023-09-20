package logs

import (
	"context"
	"fmt"
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

func (l *Logger) A() {

}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	requestId := fmt.Sprintf("RequestID:%s", ctx.Value("RequestID"))
	args = append(args, requestId)
	l.zapLogger.Error(args...)
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	requestId := fmt.Sprintf("RequestID:%s", ctx.Value("RequestID"))
	args = append(args, requestId)
	l.zapLogger.Info(args...)
}

func (l *Logger) Warn(ctx context.Context, args ...interface{}) {
	requestId := fmt.Sprintf("RequestID:%s", ctx.Value("RequestID"))
	args = append(args, requestId)
	l.zapLogger.Warn(args...)
}

func (l *Logger) Test(ctx context.Context, args ...interface{}) {
	requestId := fmt.Sprintf("RequestID:%s", ctx.Value("RequestID"))
	args = append(args, requestId)
	l.zapLogger.Warn(args...)
}
