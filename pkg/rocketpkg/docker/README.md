# `apache/rocketmq` docker run的启动方式
```bash
# 拉去远端镜像
docker pull apache/rocketmq

# 运行容器,有哪些端口需要映射去看docker inspect IMAGE_NAME
docker run -it --network dbtoes -d -p 9876:9876 -p 10909:10909 -p 10911:10911 -p 10912:10912 --name=rocketmq apache/rocketmq bash

# 进入容器内
docker exec -it CONTAINER_NAME bash

# 查看环境变量，查找bin文件
export 

# 启动server
cd /home/rocketmq/rocketmq-5.1.3/bin
nohup /home/rocketmq/rocketmq-5.1.3/bin/mqnamesrv &

# 启动broker
vi ../conf/broker.conf

# 文件尾部追加
namesrvAddr = 127.0.0.1:9876
brokerIP1 = 127.0.0.1

nohup /home/rocketmq/rocketmq-5.1.3/bin/mqbroker -c /home/rocketmq/rocketmq-5.1.3/conf/broker.conf &

# rocketmqctl url:https://rocketmq.apache.org/zh/docs/deploymentOperations/02admintool
# 创建/修改 rocket
./mqadmin updateTopic -c DefaultCluster -n 127.0.0.1:9876 -t order
# 查看topic列表
./mqadmin topicList  -n 127.0.0.1:9876
# 查看topic主题信息
./mqadmin topicStatus  -n 127.0.0.1:9876 -t order
# 查看消费者组整体信息
./mqadmin consumerProgress -n 127.0.0.1:9876 -g OrderPayGroup
# 查看消费者组连接信息
./mqadmin consumerConnection -n 127.0.0.1:9876 -g OrderPayGroup
# 查看消费者组成员
./mqadmin consumerStatus -n 127.0.0.1:9876 -g OrderPayGroup
# 手动添加信息到topic
./mqadmin sendMessage -n 127.0.0.1:9876 -t order -p "这里是手动添加" -c A -c B
```

# dockerfile启动方式
```bash
cd /path/to/dockerfile

# 构建镜像
docker build -t rocketmqdf .

# 启动容器
docker run -it -d -p 9876:9876 -p 10909:10909 -p 10911:10911 -p 10912:10912 --name=qslrocketmq rocketmqdf bash

# 问题，进到容器内ps 未发现启动服务，容器内去执行start.sh文件是可以启动服务的 
# 暂时去容器执行start.sh
/home/rocketmq/rocketmq-5.1.3/bin/start.sh

```


# docker-compose.yaml编排，加上了rocketmq_ui

