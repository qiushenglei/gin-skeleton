package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
)

func recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Log.Error(context.Background())
			}
		}()
	}
}
