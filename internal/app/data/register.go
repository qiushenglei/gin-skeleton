package data

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/canal_test/query"
	isolateModel "github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/model"
	isolateQuery "github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/query"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// MySQL connections
	MySQLCanalTestClient *gorm.DB
	MySQLIsolateClient   *gorm.DB

	// Redis connections
	RedisClient *redis.Client

	// ElasticSearch connections
	ESClient *elasticsearch.Client

	// ElasticSearch connections
	TypedESClient *elasticsearch.TypedClient

	// MongoDB connections
	//MongoClient MongoCli
	MongoClient *mongo.Client

	// Kafka connections
	KafkaClient sarama.Client
	// KafkaClient2 sarama.Client
)

type MongoCli struct {
	Cli   *mongo.Client
	Close func() error
}

// RegisterData makes new connections to MySQL, Redis, and Kafka, adn returns the closer functions of them
func RegisterData() (closers []func() error, err error) {

	// 建立 mysql 连接
	mysqlClose := RegisterMySQL()
	closers = append(closers, mysqlClose)

	// 建立 Redis 连接
	redisClose := RegisterRedis()
	closers = append(closers, redisClose)

	// 建立ES
	RegisterES()

	//  建立 Mongo 连接
	mongoClose := RegisterMongoDB()
	closers = append(closers, mongoClose)
	return
}

// RegisterMongoDB ...
func RegisterMongoDB() func() error {
	isRegister := configs.EnvConfig.GetInt("REGISTER_MONGO")
	if isRegister == 0 {
		return nil
	}

	//username := configs.EnvConfig.GetString("MONGO_USERNAME")
	//password := configs.EnvConfig.GetString("MONGO_PASSWORD")
	host := configs.EnvConfig.GetString("MONGO_HOST")
	port := configs.EnvConfig.GetString("MONGO_PORT")
	//cluster := "order"

	clientOptions := options.Client().
		// "mongodb+srv://Mongo:<password>@cluster0.qhi85.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
		//ApplyURI(utils.StringConcat("", "mongodb://", username, ":", password, "@", host, ":", port, "/", cluster, "?retryWrites=true&w=majority"))
		ApplyURI(fmt.Sprintf("mongodb://%v:%v/", host, port)).
		SetTimeout(10 * time.Second).
		SetReplicaSet("").
		SetMaxPoolSize(10)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = MongoClient.Ping(ctx, nil)
	close := func() error { return MongoClient.Disconnect(context.TODO()) }

	// RET
	return close
}

func RegisterMySQL() func() error {
	isRegister := configs.EnvConfig.GetInt("REGISTER_MYSQL")
	if isRegister == 0 {
		return nil
	}

	// mysql logger
	mysqlLogger := logs.NewGormLogger(
		logs.Log,
		logger.Config{
			Colorful:             true,
			LogLevel:             logger.Info,
			ParameterizedQueries: false, // false展示value值，true只展示sql ?占位符
		})

	// 建立 cancel_test Mysql 连接
	var err error
	DBConfig := &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),

		Logger: mysqlLogger,
	}
	username := configs.EnvConfig.GetString("DB_SPEC_RW_USERNAME")
	pw := configs.EnvConfig.GetString("DB_SPEC_RW_PASSWORD")
	host := configs.EnvConfig.GetString("DB_SPEC_RW_HOST")
	port := configs.EnvConfig.GetString("DB_SPEC_RW_PORT")
	db := "canal_test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, host, port, db)
	MySQLCanalTestClient, err = gorm.Open(mysql.Open(dsn), DBConfig)
	if err != nil {
		panic(err)
	}
	if configs.EnvConfig.GetString("RUN_MODE") == "DEBUG" {
		MySQLCanalTestClient = MySQLCanalTestClient.Debug()
	}
	MySQLCanalTestClient.Callback().Query().Register("aaa", Ceshi())

	rawdb, err := MySQLCanalTestClient.DB()
	rawdb.SetMaxIdleConns(6)
	rawdb.SetMaxOpenConns(6 * 2)
	if err != nil {
		panic(err)
	}

	// rw_isolate 读写分离库
	db1 := "rw_isolate"
	WriteHost := configs.EnvConfig.GetString("DB_CANAL_WRITE_HOST")
	WritePort := configs.EnvConfig.GetString("DB_CANAL_WRITE_PORT")
	ReadHost1 := configs.EnvConfig.GetString("DB_CANAL_READ1_HOST")
	ReadPort1 := configs.EnvConfig.GetString("DB_CANAL_READ1_PORT")
	ReadHost2 := configs.EnvConfig.GetString("DB_CANAL_READ2_HOST")
	ReadPort2 := configs.EnvConfig.GetString("DB_CANAL_READ2_PORT")
	WriteDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, WriteHost, WritePort, db1)
	Read1DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, ReadHost1, ReadPort1, db1)
	Read2DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, ReadHost2, ReadPort2, db1)
	MySQLIsolateClient, err = gorm.Open(mysql.Open(WriteDSN), DBConfig)
	MySQLIsolateClient.Use(
		dbresolver.Register( //
			dbresolver.Config{
				Sources:  []gorm.Dialector{mysql.Open(WriteDSN)},
				Replicas: []gorm.Dialector{mysql.Open(Read1DSN), mysql.Open(Read2DSN)},
			},
			isolateModel.Order1{}, "order2", //order是分表的 order1 和 order2 可以去从库1，2读取数据
		).Register(
			dbresolver.Config{
				Sources:  []gorm.Dialector{mysql.Open(WriteDSN)},
				Replicas: []gorm.Dialector{mysql.Open(Read2DSN)},
			},
			"order3", "order4", //order是分表的 order3 和 order3 只可以去从库2读取数据
		))

	//MySQLCanalTestClient.AutoMigrate()
	MySQLCanalTestClient.Callback().Query().Register("bbb", Ceshi())
	// 注册gorm generate model
	query.SetDefault(MySQLCanalTestClient)
	isolateQuery.SetDefault(MySQLIsolateClient)
	return rawdb.Close
}

func RegisterRedis() func() error {
	isRegister := configs.EnvConfig.GetInt("REGISTER_REDIS")
	if isRegister == 0 {
		return nil
	}

	RedisAddr := configs.EnvConfig.GetString("REDIS_GOODS_RW_HOST") + ":" + configs.EnvConfig.GetString("REDIS_GOODS_RW_PORT") // Host + Port
	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     RedisAddr,                                              // Host + Port
			Password: configs.EnvConfig.GetString("REDIS_GOODS_RW_PASSWORD"), // Password
			DB:       15,                                                     // variables.EnvConfig.GetInt(""), // DB
		},
	)
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return RedisClient.Close
}

func RegisterES() {
	isRegister := configs.EnvConfig.GetInt("REGISTER_ES")
	if isRegister == 0 {
		return
	}

	ESAddr := "http://" + configs.EnvConfig.GetString("ES_TEST_HOST") + ":" + configs.EnvConfig.GetString("ES_TEST_PORT")
	cfg := elasticsearch.Config{
		Addresses: []string{ESAddr},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 1000 * time.Millisecond,
			//DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
				// ...
			},
		},
	}
	//ESClient = dbtoes.NewESClient(cfg)
	TypedESClient = dbtoes.NewESTypedClient1(cfg)
}

func Ceshi() func(*gorm.DB) {
	return func(db *gorm.DB) {
		fmt.Println("is before_test")
		//db.LogMode(true)
		//db.SetLogger(dbtoes.NewLogger())
		//db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
}
