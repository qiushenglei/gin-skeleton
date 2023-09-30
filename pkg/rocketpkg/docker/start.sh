#!/bin/bash

# 启动server
nohup /home/rocketmq/rocketmq-5.1.3/bin/mqnamesrv &

# 启动broker1
/home/rocketmq/rocketmq-5.1.3/bin/mqbroker -c /home/rocketmq/rocketmq-5.1.3/conf/broker.conf
