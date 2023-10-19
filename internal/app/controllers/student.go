package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/internal/app/mq/localrocket"
	"github.com/qiushenglei/gin-skeleton/internal/app/services"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/validatorx"
)

func SetData(c *gin.Context) {
	body := &entity.StudentSetDataRequest{}
	if err := validatorx.Validate(c, body); err != nil {
		utils.Response(c, nil, err)
	}

	userID := services.SetData(c, body)
	utils.Response(c, entity.StudentSetDataResponse{
		Id: int(userID),
	}, nil)
}

func GetData(c *gin.Context) {
	body := &entity.SearchCond{}
	if err := validatorx.Validate(c, body); err != nil {
		utils.Response(c, nil, err)
	}

	services.GetData(c, body)
}

func GetESData(c *gin.Context) {

}

func Delay(c *gin.Context) {
	body := &entity.DelayRequest{}
	if err := validatorx.Validate(c, body); err != nil {
		utils.Response(c, nil, err)
	}
	err := localrocket.Producer.SyncSendMsg(c, localrocket.OrderAutoSuccessEvent, body.Content, &rocketpkg.Delay{Expire: body.Expire})
	utils.Response(c, "不知道", err)
}
