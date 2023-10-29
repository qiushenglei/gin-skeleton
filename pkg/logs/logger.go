package logs

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime"
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

// Printf gorm logger print sql
func (l *Logger) Printf(msg string, data ...interface{}) {
	l.zapLogger.Info(msg)
	l.zapLogger.Info(data...)
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	// 请求id
	args = l.addRequestID(ctx, args)
	// trace信息
	args = l.addTrace(ctx, args)
	l.zapLogger.Error(args...)
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	args = l.addRequestID(ctx, args)
	l.zapLogger.Info(args)
}

func (l *Logger) Warn(ctx context.Context, args ...interface{}) {
	args = l.addRequestID(ctx, args)
	l.zapLogger.Warn(args...)
}

func (l *Logger) addRequestID(ctx context.Context, arg []interface{}) []interface{} {
	rid := retStringRequestID(ctx)
	arg = append(arg, zap.String("requestID", rid))
	return arg
}

func (l *Logger) addTrace(ctx context.Context, arg []interface{}) []interface{} {
	trace := getTrace()
	arg = append(arg, zap.Any("trace", trace))
	return arg
}

type Strace struct {
	File    string
	FunName string
	Line    int
}

// retStringRequestID 请求id中间件设置的request id
func retStringRequestID(ctx context.Context) string {
	ginCtx, ok := ctx.(*gin.Context)
	if ok == false {
		return ""
	}
	return ginCtx.GetString("RequestID")
}

// 获取trace
func getTrace() []Strace {
	pc := make([]uintptr, 10)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	//frames.
	e := true
	var f runtime.Frame
	var trace []Strace
	for e {
		f, e = frames.Next()
		trace = append(trace, Strace{File: f.File, FunName: f.Function, Line: f.Line})
	}

	return trace
}
