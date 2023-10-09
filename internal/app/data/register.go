package data

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	redis "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// MySQL connections
	MySQLSpecClient *gorm.DB

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
	return
}

// RegistMongoDB ...
func RegistMongoDB(username, password, addr, cluster string, mc MongoCli) (err error) {
	clientOptions := options.Client().
		// "mongodb+srv://Mongo:<password>@cluster0.qhi85.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
		ApplyURI(utils.StringConcat("", "mongodb+srv://", username, ":", password, "@", addr, "/", cluster, "?retryWrites=true&w=majority"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mc.Cli, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return
	}
	err = mc.Cli.Ping(context.TODO(), nil)
	mc.Close = func() error { return mc.Cli.Disconnect(context.TODO()) }

	// RET
	return
}

func RegisterMySQL() func() error {
	isRegister := configs.EnvConfig.GetInt("REGISTER_MYSQL")
	if isRegister == 0 {
		return nil
	}

	// 建立 Mysql 连接
	DBConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	username := configs.EnvConfig.GetString("DB_SPEC_RW_USERNAME")
	pw := configs.EnvConfig.GetString("DB_SPEC_RW_PASSWORD")
	host := configs.EnvConfig.GetString("DB_SPEC_RW_HOST")
	port := configs.EnvConfig.GetString("DB_SPEC_RW_PORT")
	db := "testdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, host, port, db)
	MySQLSpecClient, err := gorm.Open(mysql.Open(dsn), DBConfig)
	if err != nil {
		panic(err)
	}
	if configs.EnvConfig.GetString("RUN_MODE") == "DEBUG" {
		MySQLSpecClient = MySQLSpecClient.Debug()
	}

	rawdb, err := MySQLSpecClient.DB()
	if err != nil {
		panic(err)
	}
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
			ResponseHeaderTimeout: 100 * time.Millisecond,
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
