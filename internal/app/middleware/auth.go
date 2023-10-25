package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
)

func AuthRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO::这里做登录验证，自己去实现
		token, err := c.Cookie("token")
		if err != nil {
			panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "get token err"))
		}

		body := make(map[string]interface{})
		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			panic(errorpkg.ErrParam)
		}

		AppID, ok := body["app_id"].(string)
		if !ok || AppID == "" {
			panic(errorpkg.NewBizErrx(errorpkg.CodeParam, "APPID is required"))
		}

		var user entity.LoginInfo
		redisKey := fmt.Sprintf("%s_%s", AppID, token)
		if loginInfo, err := data.RedisClient.WithContext(c).Get(c, redisKey).Bytes(); err == nil && len(loginInfo) > 0 {
			if err := json.Unmarshal(loginInfo, &user); err != nil {
				panic(errorpkg.ErrNoLogin)
			}
			c.Set("LoginInfo", user)
		} else {
			panic(errorpkg.ErrNoLogin)
		}

		// 业务逻辑执行
		c.Next()
	}
}
