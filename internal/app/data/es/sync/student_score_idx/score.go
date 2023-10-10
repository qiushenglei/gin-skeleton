package student_score_idx

import (
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
)

type StudentScoreScore struct {
	Id         int    `json:"id"`
	StudentId  string `json:"student_id"`
	SubjectId  int    `json:"subject_id"`
	Score      int    `json:"score"`
	AddTime    string `json:"add_time"`
	UpdateTime string `json:"update_time"`
}

func (u *StudentScoreScore) Insert(i *dbtoes.Index) error {
	return nil
}

func (u *StudentScoreScore) Update(i *dbtoes.Index) error {
	return nil
}
