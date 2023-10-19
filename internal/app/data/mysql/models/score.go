package models

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/query"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"time"
)

func NewScores(scores []*entity.Score, subjectData []*model.Subject, StudentID string) []*model.Score {
	if len(scores) == 0 {
		return nil
	}
	subjectMap := make(map[string]*model.Subject, len(subjectData))
	for _, v := range subjectData {
		subjectMap[v.SubjectName] = v
	}

	now := time.Now()
	var data []*model.Score
	for _, v := range scores {
		SubjectId := subjectMap[v.SubjectInfo.SubjectName].ID
		data = append(data, &model.Score{
			ID:         uint32(v.Id),
			StudentID:  StudentID,
			SubjectID:  SubjectId,
			Score:      int32(v.Score),
			AddTime:    now,
			UpdateTime: &now,
		})
	}
	return data
}

func GetScoreBySearchCond(c *gin.Context, cond *entity.SearchCond) ([]*model.Score, error) {
	if cond.ScoreCond.SubjectName == "" || cond.ScoreCond.Score == nil {
		return nil, nil
	}

	subject_info := GetSubjectBySearchCond(c, cond)
	if subject_info.ID == 0 {
		return nil, errorpkg.NewBizErrx(errorpkg.CodeFalse, "subject_name is not define")
	}

	var scores []*model.Score
	if err := query.Q.Score.WithContext(c).Where(query.Q.Score.Score.Eq(int32(*cond.ScoreCond.Score))).
		Where(query.Q.Score.SubjectID.In(subject_info.ID)).Scan(&scores); err != nil {
		panic(err)
	}

	if len(scores) == 0 {
		return scores, errorpkg.NewBizErrx(errorpkg.CodeFalse, "not found score data")
	}
	return scores, nil
}
