```shell
docker pull apacherocketmq/rocketmq-dashboard:latest


# 在同一个网络下，通过container_id来连接另外一个容器下的rocketmq-server
docker run -d --network dbtoes  --name rocketmq-dashboard -e "JAVA_OPTS=-Drocketmq.namesrv.addr=qslrocketmq:9876" -p 8080:8080 -t apacherocketmq/rocketmq-dashboard:latest

```