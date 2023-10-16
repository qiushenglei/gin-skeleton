package middleware

import "github.com/gin-gonic/gin"

func AuthRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO::这里做登录验证，自己去实现

		// 业务逻辑执行
		c.Next()
	}
}
