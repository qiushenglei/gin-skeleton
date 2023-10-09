package student_score_idx

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
)

type StudentScoreUser struct {
	Id         int                  `json:"id"`
	Username   string               `json:"username"`
	Label      []string             `json:"label"`
	ClassId    int                  `json:"class_id"`
	StudentId  string               `json:"student_id"`
	AddTime    string               `json:"add_time"`
	UpdateTime string               `json:"update_time"`
	IsDeleted  int                  `json:"is_deleted"`
	ClassInfo  *StudentScoreClass   `json:"class_info"`
	ScoreInfo  []*StudentScoreScore `json:"score_info"`
}

func (u *StudentScoreUser) Gan() {

}

func (u *StudentScoreUser) Insert(i *dbtoes.Index) error {
	// 获取data数据
	data := i.BodyData[0]

	// interface转成struct
	var newData StudentScoreUser
	utils.Interface2Struct(data, &newData)

	// 插入新document
	Resp, err := i.TypedESConn.Index(StudentScoreIdx).Request(newData).Do(context.Background())
	if err != nil {
		// TODO::这里根据业务补救，这不一一写了(直接panic出去，通过rocketmq的重试特性。重新消费)
		panic(err)
	}

	// 更新i的内容
	i.SetPrimaryID(Resp.Id_)
	i.SetPrimarySource(newData)

	return nil
}

func (u *StudentScoreUser) Update(i *dbtoes.Index) error {
	originData, ok := i.PrimarySource.(*StudentScoreUser)
	if ok == false {
		panic(11)
	}

	// 获取data数据
	data := i.BodyData[0]

	// interface转成struct
	var newData StudentScoreUser
	utils.Interface2Struct(data, &newData)

	// 将子表内容赋值
	newData.ClassInfo = originData.ClassInfo
	newData.ScoreInfo = originData.ScoreInfo

	Resp, err := i.TypedESConn.Update(StudentScoreIdx, i.PrimaryID).Doc(newData).Do(context.Background())
	if err != nil {
		// TODO::这里根据业务补救，这不一一写了(直接panic出去，通过rocketmq的重试特性。重新消费)
		panic(err)
	}
	i.SetPrimaryID(Resp.Id_)
	i.SetPrimarySource(newData)
	return nil
}
