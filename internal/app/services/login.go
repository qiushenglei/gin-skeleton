package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/localtime"
	"time"
)

func Login(c *gin.Context, body *entity.LoginRequest) *entity.LoginInfo {
	// 查数据库
	// set redis
	token := utils.GenerateUniqueNumberBySnowFlake()
	RedisKey := fmt.Sprintf("%s_%s", body.AppID, token)
	RedisVal := &entity.LoginInfo{

		UserID:    uint(utils.GenerateUniqueNumberByRand()),
		LoginTime: localtime.LocalTime(time.Now()),
	}
	jv, _ := json.Marshal(RedisVal)
	res, err := data.RedisClient.Set(c, RedisKey, jv, time.Hour*24).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// set cookie
	c.SetCookie("token", token, 2*60*60, "/", configs.EnvConfig.GetString("HTTP_DOMAIN"), false, true)

	return RedisVal
}
