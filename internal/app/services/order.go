package services

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/models"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
)

func FindAll(ctx *gin.Context, request *entity.FindOrderRequest) (*entity.FindOrderResponse, error) {
	res, count, err := models.FindAll(ctx, request)
	if err != nil {
		panic(err)
	}

	return &entity.FindOrderResponse{
		res,
		int(count),
	}, err
}
