FROM apache/rocketmq:latest

LABEL authors="49720"

COPY start.sh /home/rocketmq/rocketmq-5.1.3/bin/start.sh

COPY broker.conf /home/rocketmq/rocketmq-5.1.3/conf/broker.conf

WORKDIR /home/rocketmq/rocketmq-5.1.3/bin

CMD ["bash", "start.sh"]