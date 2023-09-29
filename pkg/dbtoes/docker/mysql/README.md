```shell
# 启动容器run 
# 挂载配置文件 和 本地mysql数据，容器内3306端口映射到宿主机的3307端口
# 遇到的问题 docker logs CONTAINER_ID 查看启动日志
docker run -it --name mysql --network dbtoes --name="mysql" -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/data:/var/lib/mysql -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/config/my.cnf:/etc/my.cnf -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=cancal -e MYSQL_PASSWORD=cancal -p 3307:3306  -d  mysql:5.7

# 拷贝日志到本地目录
docker cp mysql:/etc/my.cnf ./config/
docker cp mysql:usr/local/bin/docker-entrypoint.sh ./


# 测试连接
mysql -h 127.0.0.1 -u root -p my-secret-pw

# 创建mysql-canal账号
create user canal;
GRANT ALL PRIVILEGES ON *.* TO canal@"%" IDENTIFIED BY 'canal' WITH GRANT OPTION;
flush privileges; 

# 测试连接
mysql -h 127.0.0.1 -u canal -p canal


# 修改my.cnf
[mysqld]
log-bin=mysql-bin # 开启 binlog
binlog-format=ROW # 选择 ROW 模式
server_id=1 # 配置 MySQL replaction 需要定义，不要和 canal 的 slaveId 重复


 
# 想把my.cnf挂载到本地目录，mysqld启动时会判定my.cnf的权限，不允许777，所以只能到容器里又修改配置文件权限，重新启动服务。下面是尝试的2种方法都不可能
# 1.docker run里改权限，不知道为什么会报no such file or directory
# 2.想在shell文件里加chmod 644 /etc/my.cnf 也又问题
docker run -it --name mysql --network dbtoes --name="mysql" -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/data:/var/lib/mysql -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/config/my.cnf:/etc/my.cnf  -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=cancal -e MYSQL_PASSWORD=cancal -p 3307:3306  --entrypoint "bash -c 'chmod 644 /etc/my.cnf && /entrypoint.sh mysqld'" -d  mysql:5.7
docker run -it --name mysql --network dbtoes --name="mysql" -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/data:/var/lib/mysql -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/config/my.cnf:/etc/my.cnf -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/docker-entrypoint.sh:/usr/local/bin/docker-entrypoint.sh -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=cancal -e MYSQL_PASSWORD=cancal -p 3307:3306 -d  mysql:5.7

 # 1.容器内修改权限
 # 2.重启容器
 # 3.进入mysql客户端检测是否配置生效
chmod 664 /etc/my.cnf
docker restart mysql
mysql -h 127.0.0.1 -u canal -p canal
show variables like "%server_id%";
show variables like 'log_bin';
show variables like 'binlog_format';

 
 # port 
3306

# 
"Env": [
    "MYSQL_ROOT_PASSWORD=my-secret-pw",
    "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
    "GOSU_VERSION=1.16",
    "MYSQL_MAJOR=5.7",
    "MYSQL_VERSION=5.7.43-1.el7",
    "MYSQL_SHELL_VERSION=8.0.34-1.el7"
],
"Cmd": [
    "mysqld"
],
"Image": "mysql:5.7",
"Volumes": {
    "/var/lib/mysql": {}
},
"WorkingDir": "",
"Entrypoint": [
    "docker-entrypoint.sh"
],
```



"registry-mirrors": [
"https://registry.hub.docker.com",
"http://hub-mirror.c.163.com",
"https://docker.mirrors.ustc.edu.cn",
"https://registry.docker-cn.com"
]