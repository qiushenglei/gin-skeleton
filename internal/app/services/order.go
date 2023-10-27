package services

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/models"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
)

func FindAll(ctx context.Context, request *entity.FindOrderRequest) ([]*model.Order1, int64, error) {
	res, count, err := models.FindAll(ctx, request)
	if err != nil {
		//panic(err)
	}

	return res, count, err
}
