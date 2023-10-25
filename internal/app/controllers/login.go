package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/constants"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/internal/app/services"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/validatorx"
	"time"
)

func Logged(c *gin.Context) {
	body := &entity.LoggedRequest{}
	validatorx.Validate(c, body)

	var UserInfo entity.LoginInfo
	if v, exists := c.Get("LoginInfo"); !exists {
		panic(errorpkg.ErrNoLogin)
	} else {
		UserInfo = v.(entity.LoginInfo)
	}
	fmt.Println(UserInfo)

	utils.Response(c, time.Time(UserInfo.LoginTime).Format(constants.DateFormat), nil)
}

func Login(c *gin.Context) {
	body := &entity.LoginRequest{}
	validatorx.Validate(c, body)

	loginInfo := services.Login(c, body)
	utils.Response(c, loginInfo, nil)
}
