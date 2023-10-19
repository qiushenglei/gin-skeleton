package models

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/query"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
)

func NewSubject(student *entity.StudentSetDataRequest) (data []*model.Subject) {
	for _, v := range student.ScoreInfo {
		data = append(data, &model.Subject{
			ID:          uint32(v.SubjectInfo.SubjectId),
			SubjectName: v.SubjectInfo.SubjectName,
		})
	}
	return data
}

func GetSubjectBySearchCond(c *gin.Context, cond *entity.SearchCond) *model.Subject {
	if cond.ScoreCond.SubjectName == "" {
		return nil
	}
	res := &model.Subject{}
	if err := query.Q.Subject.WithContext(c).Where(query.Q.Subject.SubjectName.Eq(cond.ScoreCond.SubjectName)).Scan(res); err != nil {
		panic(err)
	}
	return res
}
