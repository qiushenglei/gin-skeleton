package student_score_idx

type StudentScoreClass struct {
	ClassId    int    `json:"class_id"`
	ClassName  string `json:"class_name"`
	Grade      int    `json:"grade"`
	AddTime    string `json:"add_time"`
	UpdateTime string `json:"update_time"`
}

func (c *StudentScoreClass) Gan() {

}

func (u *StudentScoreClass) Insert() {

}

func (u *StudentScoreClass) Update() {

}
