package middleware

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/anguloc/zet/pkg/safe"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/sentinelx"
)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := sentinel.InitWithConfigFile(safe.Path("/internal/app/sentinel/sentinel.yaml")); err != nil {
			panic(err)
		}
		e, err := sentinel.Entry(sentinelx.GlobalRateLimiter, sentinel.WithResourceType(base.ResTypeCommon))
		defer e.Exit()
		if err != nil {
			// 限流
			c.AbortWithStatusJSON(429, gin.H{
				"code": 429,
				"msg":  "请求过于频繁",
			})
			return
		}
		e.Exit()
	}
}
