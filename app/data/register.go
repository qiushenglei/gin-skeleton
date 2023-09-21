package data

import (
	"context"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/app/configs"
	"github.com/qiushenglei/gin-skeleton/app/global/utils"
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
	// RedisClient2 *redis.Client

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

	isRegister := configs.EnvConfig.GetInt("REGISTER_DATA")
	if isRegister == 0 {
		return
	}

	DBConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 建立 Mysql 连接
	username := configs.EnvConfig.GetString("DB_SPEC_RW_USERNAME")
	pw := configs.EnvConfig.GetString("DB_SPEC_RW_PASSWORD")
	host := configs.EnvConfig.GetString("DB_SPEC_RW_HOST")
	port := configs.EnvConfig.GetString("DB_SPEC_RW_PORT")
	db := "testdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, host, port, db)
	MySQLSpecClient, err = gorm.Open(mysql.Open(dsn), DBConfig)
	if err != nil {
		return
	}
	if configs.EnvConfig.GetString("RUN_MODE") == "DEBUG" {
		MySQLSpecClient = MySQLSpecClient.Debug()
	}
	if rawdb, err := MySQLSpecClient.DB(); err != nil {
		closers = append(closers, rawdb.Close)
	}

	// 建立 Redis 连接
	RedisAddr := configs.EnvConfig.GetString("REDIS_GOODS_RW_HOST") + ":" + configs.EnvConfig.GetString("REDIS_GOODS_RW_PORT") // Host + Port
	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     RedisAddr,                                              // Host + Port
			Password: configs.EnvConfig.GetString("REDIS_GOODS_RW_PASSWORD"), // Password
			DB:       15,                                                     // variables.EnvConfig.GetInt(""), // DB
		},
	)
	if err = RedisClient.Ping(context.Background()).Err(); err != nil {
		return
	}
	closers = append(closers, RedisClient.Close)

	//  建立 Mongo 连接
	//clientOptions := options.Client().
	//	ApplyURI(utils.StringConcat("", "mongodb://", configs.EnvConfig.GetString("MONGO_GOODS_RW_HOST"), ":",
	//		configs.EnvConfig.GetString("MONGO_GOODS_RW_PORT"), "/", configs.EnvConfig.GetString("MONGO_GOODS_DB")),
	//	).
	//	SetAuth(options.Credential{
	//		Username: configs.EnvConfig.GetString("MONGO_GOODS_USERNAME"), Password: configs.EnvConfig.GetString("MONGO_GOODS_PASSWORD"), // Username + Possword
	//	})
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//if MongoClient, err = mongo.Connect(ctx, clientOptions); err != nil {
	//	return
	//}
	//if err = MongoClient.Ping(context.TODO(), nil); err != nil {
	//	return
	//}
	//closers = append(closers, func() error { return MongoClient.Disconnect(context.TODO()) })

	// 建立 Kafka 连接
	// kafkaConf := sarama.NewConfig()
	// kafkaConf.Producer.Return.Successes = true
	// kafkaAddr := []string{
	// configs.EnvConfig.GetString("KAFKA_FOR_FRAMEWORK"), // Addr, in form of `Host:Port``
	// }
	// KafkaClient, err = sarama.NewClient(kafkaAddr, kafkaConf)
	// if err != nil {
	// fmt.Println(err)
	// return
	// }
	// closers = append(closers, KafkaClient.Close)

	// RET
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
