package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/internal/app/services"
	"github.com/qiushenglei/gin-skeleton/pkg/validatorx"
)

func SetData(c *gin.Context) {
	body := &entity.StudentSetData{}
	if err := validatorx.Validate(c, body); err != nil {
		utils.Response(c, nil, err)
	}

	services.SetData(c, body)

}

func GetData(ctx *gin.Context) {

}

func GetESData(ctx *gin.Context) {

}
