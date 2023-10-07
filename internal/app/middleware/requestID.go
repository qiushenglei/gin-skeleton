package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func BindRequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 生成unique，用雪花
		//这里先写成随机值
		rid := rand.Int31()
		context.Set("RequestID", fmt.Sprintf("%d", rid))
		context.Next()
	}
}
