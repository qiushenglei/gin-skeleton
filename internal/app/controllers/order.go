package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/internal/app/services"
	"github.com/qiushenglei/gin-skeleton/pkg/validatorx"
)

func FindAll(c *gin.Context) {
	body := &entity.FindOrderRequest{}
	if err := validatorx.Validate(c, body); err != nil {
		panic(err)
	}

	res, count, err := services.FindAll(c, body)
	utils.Response(c, &entity.FindOrderResponse{
		res,
		int(count),
	}, err)
}

func UpdateOrder(c *gin.Context) {
	body := &entity.UpdateOrderRequest{}
	if err := validatorx.Validate(c, body); err != nil {
		panic(err)
	}

	count, err := services.UpdateOrder(c, body)
	utils.Response(c, &entity.UpdateOrderResponse{
		int(count),
	}, err)
}
