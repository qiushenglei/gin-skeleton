# kibana
可视化es工具

```shell

# 启动服务
docker run -d --name kibana --network dbtoes -p 5601:5601 kibana:7.17.13

# 拷贝配置文件
docker cp kibana:/usr/share/kibana/config/kibana.yml ./config

# 修改配置文件连接到本地es的host

# 挂载配置文件启动
docker run -d --name kibana --network dbtoes -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/kibana/config/kibana.yml:/usr/share/kibana/config/kibana.yml -p 5601:5601 kibana:7.17.13
```