# 简介
此package是一个mysql通过canal同步到es功能的包
一共包括了7个服务，分别是
- canal:伪装成mysql从节点的服务
- canal-admin:canal-server管理服务，用于添加canal-server 和 instance。本身还有一个mysql服务，用来保存canal-server和instance的配置信息
- es:elasticsearch服务
- kibana:暂用作Elastisearch服务的可视乎工具
- mysql:从节点
- rocketmq:canal-server会把binLog的数据封装成一个可读json，发送给rocketmq,go客户端去消费mq从而同步数据
- rocketmq-dashboard:rocketmq的可视化工具

# 启动


## 配置(部分可跳过)

### rocketmq配置
broker.conf指定namserver和broker。具体可看rocketmq目录下的readme

### canal配置
需要配置canal-admin的连接信息

### canal-admin
需要配置mysql的jdbc信息


## 创建docker容器
```shell
cd /path/to/docker-compose-path

docker-compose up
```
### mysql容器内

通过root账号 进mysql服务或者工具，添加canal账号，并给远端操作权限
```bash
create user canal;
GRANT ALL PRIVILEGES ON *.* TO canal@"%" IDENTIFIED BY 'canal' WITH GRANT OPTION;
flush privileges;
```


//进入canal-admin服务，http://127.0.0.1:8089。添加canal-server和instance. canal-server配置时，无法使用域名去配置，目前我都是使用ip，导致每次更新容器时都要重新去更改ip

## go消费mq服务


gin-skeleton骨架封装了mq启动配置，只需要在`internal/app/data/mq/config.go`文件正确配置，就会启动消费者。具体案例是`DBToESEvent`配置项

启动命令行如下：
```bash
go run rocketmq -e .env.local -p 10011 -m debug
```

### 具体案例

`internal/app/data/mq/config.go`文件的`DBToESEvent`配置项的启动

#### mysql表信息
```mysql
CREATE DATABASE `canal_test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */

CREATE TABLE `user` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                        `username` varchar(255) COLLATE utf8mb4_bin NOT NULL,
                        `label` varchar(255) COLLATE utf8mb4_bin NOT NULL,
                        `class_id` int(10) unsigned NOT NULL,
                        `student_id` varchar(10) COLLATE utf8mb4_bin NOT NULL,
                        `add_time` datetime NOT NULL,
                        `update_time` datetime NOT NULL,
                        `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `sid` (`student_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin


CREATE TABLE `score` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `student_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
                         `subject_id` int(10) unsigned NOT NULL,
                         `score` int(11) NOT NULL DEFAULT '0',
                         `add_time` datetime DEFAULT NULL,
                         `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

CREATE TABLE `class` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `class_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
                         `grade` varchar(255) COLLATE utf8mb4_bin NOT NULL,
                         `add_time` datetime NOT NULL,
                         `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

#### elastisearch index的mapping
```shell
PUT /student_score_idx
{
  "mappings": {
    "properties": {
      "id": {
        "type": "long"
      },
      "username": {
        "type": "wildcard"
      },
      "label": {
        "type": "text"
      },
      "class_id": {
        "type": "long"
      },
      "student_id": {
        "type": "text"
      },
      "add_time": {
        "type": "date",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "update_time": {
        "type": "date",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "is_deleted": {
        "type": "integer"
      },
      "class_info": {
        "properties": {
          "class_id": {
            "type": "long"
          },
          "class_name": {
            "type": "text"
          },
          "grade": {
            "type": "integer"
          },
          "add_time": {
            "type": "date",
            "format": "yyyy-MM-dd HH:mm:ss"
          },
          "update_time": {
            "type": "date",
            "format": "yyyy-MM-dd HH:mm:ss"
          }
        }
      },
      "score_info": {
        "type": "nested",
        "properties": {
          "id": {
            "type": "long"
          },
          "student_id": {
            "type": "text"
          },
          "subject_id": {
            "type": "long"
          },
          "score": {
            "type": "float"
          },
          "add_time": {
            "type": "date",
            "format": "yyyy-MM-dd HH:mm:ss"
          },
          "update_time": {
            "type": "date",
            "format": "yyyy-MM-dd HH:mm:ss"
          }
        }
      }
    }
  }
}
```