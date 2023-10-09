package student_score_idx

import (
	"context"
	"fmt"
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

func (s *StudentScoreScore) Gan() {

}
func (u *StudentScoreScore) Insert(i *dbtoes.Index) {
	//
	//dataStruct, ok := i.PrimarySource.(*StudentScoreScore)
	//if ok == false {
	//	panic(11)
	//}
	//dataStruct.Score = "2020-03-10 12:20:31"
	//// plan B
	//q1 := types.NewQuery()
	//q1.Term = map[string]types.TermQuery{
	//	"_id": types.TermQuery{Value: i.PrimaryID},
	//}
	//
	//Resp, err := i.TypedESConn.Update(StudentScoreIdx, i.PrimaryID).Doc(dataStruct).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(Resp.Id_)
	//
	//return nil
}

func (u *StudentScoreScore) Update(i *dbtoes.Index) error {

	dataStruct, ok := i.PrimarySource.(*StudentScoreStruct)
	if ok == false {
		panic(11)
	}
	dataStruct.AddTime = "2020-03-10 12:20:31"

	Resp, err := i.TypedESConn.Update(StudentScoreIdx, i.PrimaryID).Doc(dataStruct).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(Resp.Id_)

	//i.TypedESConn.Index(StudentScoreIdx).Request().Do(context.Background())

	return nil
}
