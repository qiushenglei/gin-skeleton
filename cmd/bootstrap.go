package cmd

import (
	"fmt"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/mq"
	"github.com/qiushenglei/gin-skeleton/internal/app/sentinelx"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

// RegistAll regists everything (Configs, Loggers, Data Connections, etc.) but the routes
func RegistAll(serverName string) (closers []func() error) {

	// 注册 配置文件
	err := registerConfig()
	if err != nil {
		panic("Failed to regist config: " + err.Error())
	}

	// 注册 Logger
	syncers, err := logs.RegisterLogger()
	if err != nil {
		panic("Failed to regist logger: " + err.Error())
	}

	// 注册限流sentinel
	sentinelx.RegisterSentinelRule()

	// 注册 数据层的连接
	dataClosers, err := data.RegisterData()
	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		closeAllClosers(ctx, cancel, closers)
		panic("Failed to regist data connections: " + err.Error())
	}
	closers = append(closers, dataClosers...)

	// 注册 MQ product
	err = mq.RegisterMQ(serverName)
	if err != nil {
		panic("Failed to regist rocketMQ Producer: " + err.Error())
	}

	// 最后再显式地注册日志 syncers 到 closers 中。这样，日志文件最后关闭
	closers = append(closers, syncers)

	// RET
	return
}

func registerConfig() error {

	// 读取 `.env` 文件的信息，存入配置文件的全局变量
	if ParseEnv() != nil {
		return nil
	}
	// RET
	return nil
}

// ListenSignal 监听信号
func ListenSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("挂起服务启动协程") // graceful shutdown
	<-quit
	fmt.Println("\nShutdown all server ...")
}

// GracefulShutdown frees all resources prior to exiting
func GracefulShutdown(closers []func() error) {

	// 设置五秒超时的 graceful shutdown
	defer fmt.Println("Server has been gracefully shutdown ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// 释放所有资源
	closeAllClosers(ctx, cancel, closers)
}

// closeAllClosers simply closes all the closer functions, in other words, frees all resources
func closeAllClosers(ctx context.Context, cancel context.CancelFunc, closers []func() error) {
	defer cancel()

	// 释放所有占有的资源
	for _, closer := range closers {
		if err := closer(); err != nil {
			logs.Log.Error(ctx, err)
		}
	}
}
