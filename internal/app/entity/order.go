package entity

import "github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/model"

type FindOrderRequest struct {
	AppID    string `json:"app_id" validate:"required,len=5"`
	Page     *int   `json:"page" validate:"required,number,gt=0"`
	PageSize *int   `json:"page_size" validate:"required,number,gt=0"`
}

type FindOrderResponse struct {
	Data  []*model.Order1 `json:"data"`
	Count int             `json:"count"`
}
