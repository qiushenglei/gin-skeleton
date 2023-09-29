```shell
# docker pull cancal/canal-server:v1.1.4
# docker run address 在
docker run -it --name canal-server --network dbtoes  -e canal.auto.scan=false  -e canal.destinations=test  -e canal.instance.master.address=172.18.0.2:3306   -e canal.instance.dbUsername=canal   -e canal.instance.dbPassword=canal   -e canal.instance.connectionCharset=UTF-8  -e canal.instance.tsdb.enable=true  -e canal.instance.gtidon=false   -p 11112:11112  -p 11111:11111 -p 11110:11110   -d canal/canal-server:v1.1.4

# 启动遇到问题，docker logs container_id 没有日志，一直启动失败。
# 发现是windows下的wsl2在centos6下，有问题。需要在 %userprofile%\.wslconfig添加下面内容 https://github.com/microsoft/WSL/issues/4694#issuecomment-556095344
[wsl2]
kernelCommandLine = vsyscall=emulate

# 拷贝配置文件，并挂载到本地config
docker cp canal-server:/home/admin/canal-server/conf/canal.properties ./config/
docker cp canal-server:/home/admin/canal-server/conf/canal_local.properties ./config/

# 启动挂载目录
docker run -it --name canal-server --network dbtoes -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/canal/config/canal.properties:/home/admin/canal-server/conf/canal.properties -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/canal/config/canal_local.properties:/home/admin/canal-server/conf/canal_local.properties  -p 11112:11112  -p 11111:11111 -p 11110:11110   -d canal/canal-server:v1.1.4


```