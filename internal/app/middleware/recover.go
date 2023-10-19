package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
)

func GlobalRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Log.Error(context.Background())
				if r, ok := r.(error); ok {
					utils.Response(c, nil, r)
				}
				if r, ok := r.(string); ok {
					utils.Response(c, nil, errorpkg.NewIOErrx(errorpkg.CodeFalse, r))
				}
			}
		}()
		c.Next()
	}
}
