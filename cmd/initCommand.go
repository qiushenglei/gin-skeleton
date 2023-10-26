package cmd

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/qiushenglei/gin-skeleton/internal/app/crontabs"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/constants"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	localgrpc "github.com/qiushenglei/gin-skeleton/internal/app/grpc"
	"github.com/qiushenglei/gin-skeleton/internal/app/mq/localrocket"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"github.com/qiushenglei/gin-skeleton/proto"
	routes "github.com/qiushenglei/gin-skeleton/routes/http"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	//命令行参数
	RootCmd = &cobra.Command{}

	// go build -gcflags '-m -l' main.go

	// 因为这里是子命令行web，所以build之后，还需要./exefile run -e .env.local -p 10011\
	ServerCmd = &cobra.Command{
		Use:              "web",
		Short:            "启动 web 服务",
		Long:             "启动 web 服务",
		Run:              RunHttpServer,
		PersistentPreRun: HttpServerPersistentPreRun,
	}

	// 因为这里是子命令行crontab，所以build之后，还需要./exefile crontab -e .env.local -p 10011
	CrontabCmd = &cobra.Command{
		Use:   "crontab",
		Short: "启动 定时任务 服务",
		Long:  "启动 定时任务 服务",
		Run:   RunCrontab,
	}

	// 因为这里是子命令行rocketmq，所以build之后，还需要./exefile crontab -e .env.local -p 10011
	RocketMQCmd = &cobra.Command{
		Use:   "rocketmq",
		Short: "启动 rocketmq消费者 服务",
		Long:  "启动 rocketmq消费者 服务",
		Run:   RunRocketMQ,
	}

	// 因为这里是子命令行rpc，所以build之后，还需要./exefile rpc -e .env.local -p 10011
	RPCCmd = &cobra.Command{
		Use:   "rpc",
		Short: "启动 rpc 服务",
		Long:  "启动 rpc 服务",
		Run:   RunRPC,
	}
)

func CmdExecute() {
	// 接收参数初始化变量
	RootCmd.PersistentFlags().StringVarP(&configs.BasePath, "base_path", "b", "./", "项目根目录")
	RootCmd.PersistentFlags().StringVarP(&configs.EnvFile, "env_file", "e", ".env", ".env文件名称")
	RootCmd.PersistentFlags().StringVarP(&configs.HttpPort, "http_port", "p", "10011", "http端口")
	RootCmd.PersistentFlags().StringVarP(&configs.AppRunMode, "mode", "m", constants.ReleaseMode, "运行模式")

	// RPC服务参数
	RPCCmd.PersistentFlags().StringVarP(&configs.RpcPort, "rpc_port", "r", "10012", "rpc服务端口")

	RootCmd.AddCommand(ServerCmd, CrontabCmd, RocketMQCmd, RPCCmd)
	if err := RootCmd.Execute(); err != nil {
		return
	}
}

func HttpServerPersistentPreRun(cmd *cobra.Command, args []string) {
	//boot
	cmd.Flags().GetString("http_port")

}

func RunHttpServer(cmd *cobra.Command, args []string) {
	// 注册除了路由以外的所有东西
	closers := RegistAll(cmd.Use)

	// 设置gin运行模式
	gin.SetMode(configs.AppRunMode)

	// 注册路由
	r := routes.RegisterRoutes()

	// 注册 HTTP 响应服务
	srv := &http.Server{Addr: utils.StringConcat("", ":", configs.HttpPort), Handler: r}
	// 协程走起

	// 启动 HTTP 响应服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.ErrorLogger.Info("Graceful shutdown error", zap.Error(err))
		}
	}()

	ListenSignal()

	// 结束 HTTP 响应
	if srv != nil {
		if err := srv.Shutdown(context.Background()); err != nil {
		}
	}
	GracefulShutdown(closers)
}

func RunCrontab(cmd *cobra.Command, args []string) {
	// 注册除了路由以外的所有东西
	closers := RegistAll(cmd.Use)

	// crontab
	c := cron.New()

	for _, eachJob := range crontabs.JobList {
		_, err := c.AddFunc(eachJob.Schedule, eachJob.Fn)
		if err != nil {
			logs.Log.Error(context.Background(), err.Error())
			panic(err)
		}
	}

	c.Start()

	ListenSignal()

	// 结束 HTTP 响应
	if c != nil {
		if err := c.Stop(); err != nil {
		}
	}
	GracefulShutdown(closers)

}

func RunRocketMQ(cmd *cobra.Command, args []string) {
	// 注册除了路由以外的所有东西
	closers := RegistAll(cmd.Use)

	localrocket.RegisterRocketMQConsumer()

	ListenSignal()

	GracefulShutdown(closers)

}

func RunRPC(cmd *cobra.Command, args []string) {
	// 注册除了路由以外的所有东西
	closers := RegistAll(cmd.Use)

	lis, err := net.Listen("tcp", "127.0.0.1:"+configs.RpcPort)
	if err != nil {
		return
	}
	opt := []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Second * 10),
		grpc.ChainUnaryInterceptor(GRPCInterceptor),
	}

	GrpcServer := grpc.NewServer(opt...)
	proto.RegisterOrderServerServer(GrpcServer, &localgrpc.OrderServer{})

	//为了优雅的关闭，到了子协程里开启grpc服务。Server会for自旋，accept socket等待连接
	go func() {
		if err = GrpcServer.Serve(lis); err != nil {
			return
		}
	}()

	closers = append(closers, func() error {
		GrpcServer.Stop()
		return nil
	})

	ListenSignal()
	GracefulShutdown(closers)
}

func GRPCInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// before
	//fmt.Println(info.FullMethod)
	// 这个就是中间件的next
	resp, err = handler(ctx, req)
	// after
	return resp, err
}
