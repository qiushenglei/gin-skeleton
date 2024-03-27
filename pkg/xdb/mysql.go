package xdb

import (
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RDB struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DBName       string `json:"dbname"`
	IdleConn     int    `json:"idle_conn"`
	OpenConn     int    `json:"open_conn"`
	LogWriteType int    `json:"log_write_type"`
}

const (
	ConsoleLogWrite int = iota
	FileLogWrite
)

func NewRDB() *RDB {
	return &RDB{}
}

func (r *RDB) RegisterMySQL() (*gorm.DB, error) {

	var log logger.Interface
	switch r.LogWriteType {
	case FileLogWrite:
		log = logs.NewGormLogger(
			logs.Log,
			logger.Config{
				Colorful:             true,
				LogLevel:             logger.Info,
				ParameterizedQueries: false, // false展示value值，true只展示sql ?占位符
			})
	case ConsoleLogWrite:
	default:
		//log = logger.Default
		log = logger.Default.LogMode(logger.Info)
	}

	// 建立 cancel_test Mysql 连接
	var err error
	DBConfig := &gorm.Config{

		Logger: log,
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", r.Username, r.Password, r.Host, r.Port, r.DBName)
	cli, err := gorm.Open(mysql.Open(dsn), DBConfig)
	if err != nil {
		return cli, err
	}

	rawdb, err := cli.DB()
	rawdb.SetMaxIdleConns(6)
	rawdb.SetMaxOpenConns(6 * 2)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
