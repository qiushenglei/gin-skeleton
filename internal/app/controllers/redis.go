package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"time"
)

func RedisString(c *gin.Context) {
	key := "stinga"
	val := 3
	data.RedisClient.Get(c, key)
	//data.RedisClient.Set(c, key, val, time.Second*5)
	//data.RedisClient.Incr(c, key)
	//data.RedisClient.IncrBy(c, key, 2)
	//data.RedisClient.MSet(c, map[string]interface{}{"a": 1, "b": 2, "c": 3})
	if ok, err := data.RedisClient.SetNX(c, key, val, time.Second*5).Result(); !ok {
		fmt.Println(err.Error())
	}
}

func RedisHyperLogLog(c *gin.Context) {
	key := "hypeprloglog1"
	i, err := data.RedisClient.PFAdd(c, key, "elem1").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(i)
}

func RedisMutex(c *gin.Context) {
	pool := goredis.NewPool(data.RedisClient)
	rsync := redsync.New(pool)
	m := rsync.NewMutex("mutex", redsync.WithValue("context"), redsync.WithExpiry(time.Second*30))

	m.Lock()
	fmt.Println(111)
}
