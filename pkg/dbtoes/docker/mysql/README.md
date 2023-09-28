```shell
# run ，遇到的问题 docker logs CONTAINER_ID 查看启动日志
docker run -it --name mysql --network dbtoes --name="mysql" -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/data:/var/lib/mysql -v /f/go_code/gin-skeleton/pkg/dbtoes/docker/mysql/config:/etc/mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=cancal -e MYSQL_PASSWORD=cancal -p 3306:3306 -d mysql:5.7


# 拷贝日志到本地目录
 docker cp mysql:/etc/mysql ./config
 
# 测试连接
mysql -h 127.0.0.1 -u root -p my-secret-pw
# 创建mysql-canal账号
create user canal;
GRANT ALL PRIVILEGES ON *.* TO canal@"%" IDENTIFIED BY 'canal' WITH GRANT OPTION;
flush privileges; 
 
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


docker run --name sicanal -e canal.auto.scan=false  -e canal.destinations=test  -e canal.instance.master.address=172.18.0.2:3306   -e canal.instance.dbUsername=canal   -e canal.instance.dbPassword=canal   -e canal.instance.connectionCharset=UTF-8  -e canal.instance.tsdb.enable=true  -e canal.instance.gtidon=false   -p 11112:11112  -p 11111:11111 -p 11110:11110   -d canal/canal-server:v1.1.4

,
"registry-mirrors": [
"https://registry.hub.docker.com",
"http://hub-mirror.c.163.com",
"https://docker.mirrors.ustc.edu.cn",
"https://registry.docker-cn.com"
]