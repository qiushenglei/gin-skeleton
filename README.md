# qsl的gin骨架
- 按日切割日志: `zap`、`go-file-rotatelogs`
- 配置文件: `viper`
- http请求: `imroc/req`
- mysql
- redis
- mongodb
- rocketmq: 添加本地调试的docker镜像 path:pkg/rocketpkg/docker

# 服务列表
- http
- crontab

```bash
# http服务启动 wen子命令行 -e 配置文件 -p http服务端口 -m 运行方式debug gorm会打印sql request会打印请求头和相应头和body
go run . web -e .env.local -p 10011 -m debug


# http服务启动 wen子命令行 -e 配置文件 -p http服务端口 -m 运行方式debug gorm会打印sql request会打印请求头和相应头和body
go run . crontab -e .env.local -p 10011 -m debug
```
