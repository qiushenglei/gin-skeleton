package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
