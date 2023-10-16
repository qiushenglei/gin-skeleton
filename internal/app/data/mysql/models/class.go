package models

import (
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"strconv"
	"time"
)

func NewClass(class *entity.ClassInfo) *model.Class {
	if class == nil {
		panic(errorpkg.ErrLogic)
	}
	now := time.Now()
	return &model.Class{
		ID:         uint32(class.ClassId),
		ClassName:  class.ClassName,
		Grade:      strconv.Itoa(class.Grade),
		AddTime:    now,
		UpdateTime: &now,
	}

}
