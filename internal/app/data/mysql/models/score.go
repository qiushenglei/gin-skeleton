package models

import (
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"time"
)

func NewScores(scores []*entity.Score, StudentID string) []*model.Score {
	if len(scores) == 0 {
		return nil
	}
	now := time.Now()
	var data []*model.Score
	for _, v := range scores {
		data = append(data, &model.Score{
			ID:         uint32(v.Id),
			StudentID:  StudentID,
			SubjectID:  uint32(v.SubjectId),
			Score:      int32(v.Score),
			AddTime:    now,
			UpdateTime: now,
		})

	}
}
