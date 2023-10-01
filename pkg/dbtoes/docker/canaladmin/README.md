# 名词

- canal-server：canal单机实例
  - instance：关于mysql同步的配置(db表，mq等配置)
- canal-admin：管理canal-server的UI界面，注册canal-server实例，并创建instance(如果不用canal-admin是需要在canal-server配置文件目录下配置)。canal-admin通过mysql将配置信息保存起来。

# docker 服务启动
```shell
# docker pull cancal/canal-admin:v1.1.4
# docker run 
docker run -it --name canal-admin --network dbtoes  -p 8089:8089 -d canal/canal-admin:v1.1.5

# 拷贝配置文件，并挂载到本地config[目录里已经有配置文件，就不需要再拷贝了]
docker cp canal-admin:/home/admin/canal-admin/conf/application.yml ./config/

# 看canal-admin的mysql数据目录(/etc/my.cnf 配置文件下指定了mysql的目录，正常情况下都mysql配置文件都在这个目录下)
# 挂载到本地的mysqldata，`docker run -v` 挂载的目录如果是空，会先将镜像的文件copy到本地挂载目录。挂载是为了避免每次创建镜像就丢失测试数据（添加到.gitignore）
# 启动挂载目录
# 这个容器里还有个mysql服务，是保存canal-adminUI界面配置信息(注册canal-server和instance配置等)的，但是此容器里改application.xml配置文件无法改动port，因为他是shell脚本写死的，人麻了。想改就只能改shell启动脚本。此处暂不修改
docker run -it --name canal-admin1 --network dbtoes -p 8089:8089 -p 3306:3306 -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/canaladmin/config/application.yml:/home/admin/canal-admin/conf/application.yml -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/canaladmin/mysqldata:/var/lib/mysql  -d canal/canal-admin:v1.1.5

# 登录UI界面，宿主机访问 http://127.0.0.1:8089  admin 123456

# 优先启动canal-server，canal-admin的server列表会自动注册一个进去

mysql -h 127.0.0.1 -u canal1 -pcanal1
mysql -u canal -p
show variables like "%port%";

```

# 遇到问题
## instance一直启动不了，提示` log_file no such file or directory`
canal-server的启动配置，只需要配置canal-admin的连接信息，本仓库的canal.properties文件内容就行
https://github.com/alibaba/canal/issues/2129#issuecomment-527733476

## 远端mysql的 `Table 'canal_test.base table' doesn't exist`
canal-server:1.1.6版本有问题，直接切到1.1.5。作者回复1.1.6的hotfix有热修复。
https://github.com/alibaba/canal/issues/4291#issuecomment-1183186920

## 启动instance 提示 rocketmq connect to 127.0.0.1:10911 failed
启动rocketmq broker时，broker.conf的配置的brokerIP1 = 127.0.0.1，canal-admin连接成了本地的127.0.0.1:10911
把配置项brokerIP1 = CONTAINER_ID 改成容器名，会指向正确的ip地址