# 示例说明

canal伪装成mysql的slave节点，数据变更时生产消息到rocketmq消息中间件。go-client启动消费者同步到es

# mysql scheme and data


# create es index by json
1. 进入kibana服务，创建一条数据
2. 读取index/mapping
3. 字段修改成自己需要的类ing

```shell
# 创建
PUT student_score/_doc/1
{
    "id": 1,
    "username" : "qsl",
    "label": ["1", "2"],
    "class_id": 1,
    "student_id": 1,
    "add_time": "2023-10-01 22:02:43",
    "update_time": "2023-10-01 22:02:43",
    "is_deleted": 0,
    "class_info": {
        "class_id": 1,
        "class_name": "97班",
        "grade": 6,
        "add_time": "2023-10-01 22:02:43",
        "update_time": "2023-10-01 22:02:43"
    },
    "score_info": [
        {
            "id": 1,
            "subject_id": 1,
            "score": 95,
            "add_time": "2023-10-01 22:02:43",
            "update_time": "2023-10-01 22:02:43"
        },
        {
            "id": 1,
            "subject_id": 1,
            "score": 95,
            "add_time": "2023-10-01 22:02:43",
            "update_time": "2023-10-01 22:02:43"
        }
    ]
}
```

## 动态映射
```shell
# 动态映射，获取自动成的mapping
GET student_score/_mapping
{
  "student_score" : {
    "mappings" : {
      "properties" : {
        "add_time" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "class_id" : {
          "type" : "long"
        },
        "class_info" : {
          "properties" : {
            "add_time" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "class_id" : {
              "type" : "long"
            },
            "class_name" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "grade" : {
              "type" : "long"
            },
            "update_time" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "id" : {
          "type" : "long"
        },
        "is_deleted" : {
          "type" : "long"
        },
        "label" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "score_info" : {
          "properties" : {
            "add_time" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "id" : {
              "type" : "long"
            },
            "score" : {
              "type" : "long"
            },
            "subject_id" : {
              "type" : "long"
            },
            "update_time" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "student_id" : {
          "type" : "long"
        },
        "update_time" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "username" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
}


```

## 静态映射
因为动态隐式帮忙定义了很多子字段(一个名字多中类型)，暂不需要，所以案例中给了一些简单的定义
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

# rocketmq获取到的数据结构体
分为N种类型
- insert:
- update：

## update示例具体内容
修改了原有数据的student_id字段
```json
{
    "data": [
        {
            "id": "1",
            "username": "qsl",
            "label": "1,2",
            "class_id": "1",
            "student_id": "2",
            "add_time": "2023-10-01 22:02:43",
            "update_time": "2023-10-01 22:02:43",
            "is_deleted": "0"
        }
    ],
    "database": "canal_test",
    "es": 1696317168000,
    "id": 11,
    "isDdl": false,
    "mysqlType": {
        "id": "int unsigned",
        "username": "varchar(255)",
        "label": "varchar(255)",
        "class_id": "int(10) unsigned",
        "student_id": "int(10) unsigned",
        "add_time": "datetime(0)",
        "update_time": "datetime(0)",
        "is_deleted": "tinyint(4)"
    },
    "old": [
        {
            "student_id": "1"
        }
    ],
    "pkNames": [
        "id"
    ],
    "sql": "",
    "sqlType": {
        "id": 4,
        "username": 12,
        "label": 12,
        "class_id": 4,
        "student_id": 4,
        "add_time": 93,
        "update_time": 93,
        "is_deleted": -6
    },
    "table": "user",
    "ts": 1696317168155,
    "type": "UPDATE"
}
```

```json
{
    "data": [
        {
            "id": "2",
            "username": "ld1",
            "label": "2,3,4",
            "class_id": "1",
            "student_id": "3",
            "add_time": "2023-10-05 13:21:09",
            "update_time": "2023-10-05 13:21:11",
            "is_deleted": "0"
        }
    ],
    "database": "canal_test",
    "es": 1696483376000,
    "id": 5,
    "isDdl": false,
    "mysqlType": {
        "id": "int unsigned",
        "username": "varchar(255)",
        "label": "varchar(255)",
        "class_id": "int(10) unsigned",
        "student_id": "int(10) unsigned",
        "add_time": "datetime(0)",
        "update_time": "datetime(0)",
        "is_deleted": "tinyint(4)"
    },
    "old": [
        {
            "username": "ld",
            "label": "2,3"
        }
    ],
    "pkNames": [
        "id"
    ],
    "sql": "",
    "sqlType": {
        "id": 4,
        "username": 12,
        "label": 12,
        "class_id": 4,
        "student_id": 4,
        "add_time": 93,
        "update_time": 93,
        "is_deleted": -6
    },
    "table": "user",
    "ts": 1696483376712,
    "type": "UPDATE"
}
```

## insert示例具体内容
```json
{
  "data": [
    {
      "id": "1",
      "subject_name": "历史"
    }
  ],
  "database": "canal_test",
  "es": 1696400416000,
  "id": 1,
  "isDdl": false,
  "mysqlType": {
    "id": "int unsigned",
    "subject_name": "varchar(255)"
  },
  "old": null,
  "pkNames": [
    "id"
  ],
  "sql": "",
  "sqlType": {
    "id": 4,
    "subject_name": 12
  },
  "table": "subject",
  "ts": 1696401308805,
  "type": "INSERT"
}
```

```json

{
  "data": [
    {
      "id": "2",
      "username": "ld",
      "label": "2,3",
      "class_id": "1",
      "student_id": "3",
      "add_time": "2023-10-05 13:21:09",
      "update_time": "2023-10-05 13:21:11",
      "is_deleted": "0"
    }
  ],
  "database": "canal_test",
  "es": 1696483273000,
  "id": 4,
  "isDdl": false,
  "mysqlType": {
    "id": "int unsigned",
    "username": "varchar(255)",
    "label": "varchar(255)",
    "class_id": "int(10) unsigned",
    "student_id": "int(10) unsigned",
    "add_time": "datetime(0)",
    "update_time": "datetime(0)",
    "is_deleted": "tinyint(4)"
  },
  "old": null,
  "pkNames": [
    "id"
  ],
  "sql": "",
  "sqlType": {
    "id": 4,
    "username": 12,
    "label": 12,
    "class_id": 4,
    "student_id": 4,
    "add_time": 93,
    "update_time": 93,
    "is_deleted": -6
  },
  "table": "user",
  "ts": 1696483273479,
  "type": "INSERT"
}
```