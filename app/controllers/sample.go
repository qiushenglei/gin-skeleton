package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/app/entity"
	"github.com/qiushenglei/gin-skeleton/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/app/services"
)

// SetKeyValue ...
func SetKeyValue(c *gin.Context) {
	// 参数校验
	req := entity.SampleSetRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, nil, err)
		return
	}

	// 生成一个 SampleService 实例
	s := services.NewSampleService(c)
	// 调用 SetKeyValue 方法
	ret, err := s.SetKeyValue(c, &req)

	// 返回
	utils.Response(c, ret, err)
}

// GetKeyValue ...
func GetKeyValue(c *gin.Context) {
	// 参数校验
	req := entity.SampleGetRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, nil, err)
		return
	}

	// 生成一个 SampleService 实例
	s := services.NewSampleService(c)
	// 调用 SetKeyValue 方法
	ret, err := s.GetKeyValue(c, &req)

	// 返回
	utils.Response(c, ret, err)
}
