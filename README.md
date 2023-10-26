# qsl的gin骨架
- 按日切割日志: `zap`、`go-file-rotatelogs` 按天分割文件
- 配置文件: `viper`
- http请求: `imroc/req`
- mysql
- redis
- mongodb
- mysql同步es: 目录pkg/dbtoes/ docker目录下包含了所需要的镜像和配置文件

# 服务列表
- http（会随着http的服务启动rocketmq的生产者）
- crontab
- rocketmq consumer

```bash
# http服务启动 wen子命令行 -e 配置文件 -p http服务端口 -m 运行方式debug gorm会打印sql request会打印请求头和相应头和body
cd cmd/server
go run web -e .env.local -p 10011 -m debug

# http服务启动 wen子命令行 -e 配置文件 -p http服务端口 -m 运行方式debug gorm会打印sql request会打印请求头和相应头和body
cd cmd/server
go run  crontab -e .env.local -p 10011 -m debug

# rocketmq消费者
cd cmd/server
go run rocketmq -e .env.local -p 10011 -m debug

# 生成gorm的DAO和model
go cd cmd/gorm_gene
go run .
```
