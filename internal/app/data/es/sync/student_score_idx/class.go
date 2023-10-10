package student_score_idx

import "github.com/qiushenglei/gin-skeleton/pkg/dbtoes"

type StudentScoreClass struct {
	ClassId    int    `json:"class_id,string"`
	ClassName  string `json:"class_name"`
	Grade      int    `json:"grade,string"`
	AddTime    string `json:"add_time"`
	UpdateTime string `json:"update_time"`
}

func (u *StudentScoreClass) Insert(i *dbtoes.Index) error {
	return nil
}

func (u *StudentScoreClass) Update(i *dbtoes.Index) error {
	return nil
}
