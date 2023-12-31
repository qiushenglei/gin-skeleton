package models

import (
	"encoding/json"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/canal_test/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"time"
)

func NewUser(student *entity.StudentSetDataRequest) *model.User {
	label, _ := json.Marshal(student.Label)
	if student.StudentId == "" {
		student.StudentId = utils.GenerateUniqueNumberBySnowFlake()
	}
	now := time.Now()
	user := &model.User{
		ID:         uint32(student.Id),
		Username:   student.Username,
		Label:      string(label),
		ClassID:    uint32(student.ClassId),
		StudentID:  student.StudentId,
		AddTime:    now,
		UpdateTime: &now,
	}
	return user
}
