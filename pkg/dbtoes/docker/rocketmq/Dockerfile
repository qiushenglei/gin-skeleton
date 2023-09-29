FROM apache/rocketmq:latest

LABEL authors="49720"

#COPY start.sh /home/rocketmq/rocketmq-5.1.3/bin/start.sh

#COPY broker.conf /home/rocketmq/rocketmq-5.1.3/conf/broker.conf

#WORKDIR /home/rocketmq/rocketmq-5.1.3/bin

#ENTRYPOINT ["/bin/bash", "/home/rocketmq/rocketmq-5.1.3/bin/start.sh"， ""]

ENTRYPOINT ["bash", "-c", "/home/rocketmq/rocketmq-5.1.3/bin/start.sh && while true; do sleep 1; done"]