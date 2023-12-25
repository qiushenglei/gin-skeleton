package entity

import (
	"github.com/qiushenglei/gin-skeleton/pkg/localtime"
)

type LoginInfo struct {
	UserID    uint                `json:"user_id"`
	LoginTime localtime.LocalTime `json:"login_time"`
}

type LoggedRequest struct {
	AppID string `json:"app_id"`
}

type LoginRequest struct {
	AppID  string `json:"app_id" validate:"required,len=5"`
	UserID uint   `json:"user_id"`
}
