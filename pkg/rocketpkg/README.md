```bash
docker pull apache/rocketmq
winpty docker run -it -d -p 9876:9876 apache/rocketmq bash
docker exec -it CONTAINER_NAME bash
# 查看bin文件
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

```

