package middleware

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func BindRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成unique，用雪花
		//这里先写成随机值
		rid := rand.Int31()
		c.Set("RequestID", fmt.Sprintf("%d", rid))
		c.Next()
		logging.Info("request over")
	}
}
