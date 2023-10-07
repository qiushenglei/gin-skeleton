# 名词
- NameServer 充当了注册中心的角色，主要用于管理和协调 Broker 的信息。它维护了一个存储了所有 Broker 信息的路由表，包括了 Topic、Queue 等信息。客户端在发送消息之前会先从 NameServer 获取到 Broker 的路由信息，然后才能发送消息给对应的 Broker。

- Broker 是消息存储和处理的服务器节点。它接收来自生产者的消息并进行持久化存储，同时也提供了消费者订阅和消费消息的功能。Broker 之间可以形成集群来提供高可用性和负载均衡的支持。在 RocketMQ 中，Broker 包含了主节点（Master）和从节点（Slave），其中主节点负责消息的写入和读取，而从节点用于备份主节点的数据，以实现高可用性和数据冗余。
> 要正确使用 RocketMQ，你需要启动 NameServer 和 Broker。首先启动 NameServer 节点，然后启动 Broker 节点并将 Broker 注册到 NameServer 上。生产者和消费者在发送和接收消息时，会与 NameServer 进行交互以获取 Broker 的路由信息，从而进行消息的发送和消费。

# docker启动容器
## docker run
```bash
# 拉去远端镜像
docker pull apache/rocketmq:

# 运行容器,有哪些端口需要映射去看docker inspect IMAGE_NAME
docker run -it --network dbtoes -d -p 9876:9876 -p 10909:10909 -p 10911:10911 -p 10912:10912 --name=qslrocketmq apache/rocketmq bash

# 进入容器内
docker exec -it CONTAINER_NAME bash

# 查看环境变量，查找bin文件
export 

# 启动server
cd /home/rocketmq/rocketmq-5.1.3/bin
nohup /home/rocketmq/rocketmq-5.1.3/bin/mqnamesrv &

# broker配置文件拷贝
cd path/to/local_config_path
docker cp qslrocketmq:/home/rocketmq/rocketmq-5.1.3/conf/broker.conf ./

# 启动broker
vi ../conf/broker.conf

# 文件尾部追加，配置成container_id，不然canal-admin启动instance会连接失败
namesrvAddr = CONTAINER_ID:9876 
brokerIP1 = CONTAINER_ID

nohup /home/rocketmq/rocketmq-5.1.3/bin/mqbroker -c /home/rocketmq/rocketmq-5.1.3/conf/broker.conf &

```


##  dockerfile启动方式
```bash
cd /path/to/dockerfile

# 构建镜像
docker build -t rocketmqdf .

# 启动容器
docker run -it --network dbtoes -d -p 9876:9876 -p 10909:10909 -p 10911:10911 -p 10912:10912 --name=qslrocketmq rocketmqdf bash
docker run -it --network dbtoes -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/rocketmq/broker.conf:/home/rocketmq/rocketmq-5.1.3/conf/broker.conf -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/rocketmq/store:/home/rocketmq/store -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/rocketmq/start.sh:/home/rocketmq/rocketmq-5.1.3/bin/start.sh  -d -p 9876:9876 -p 10909:10909 -p 10911:10911 -p 10912:10912 --name=qslrocketmq rocketmqdf bash


# 问题，进到容器内ps 未发现启动服务，容器内去执行start.sh文件是可以启动服务的 
# 暂时去容器执行start.sh
/home/rocketmq/rocketmq-5.1.3/bin/start.sh

```


## docker-compose.yaml编排，加上了rocketmq_ui

# rocketmq ctl的使用
```shell
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
```