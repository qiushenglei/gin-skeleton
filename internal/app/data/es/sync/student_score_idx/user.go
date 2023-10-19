package student_score_idx

import (
	"context"
	"encoding/json"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
)

type StudentScoreUser struct {
	Id         int                  `json:"id,string"`
	Username   string               `json:"username"`
	Label      []string             `json:"label"`
	ClassId    int                  `json:"class_id,string"`
	StudentId  string               `json:"student_id"`
	AddTime    string               `json:"add_time"`
	UpdateTime string               `json:"update_time"`
	IsDeleted  int                  `json:"is_deleted,string"`
	ClassInfo  *StudentScoreClass   `json:"class_info"`
	ScoreInfo  []*StudentScoreScore `json:"score_info"`
}

func (u *StudentScoreUser) Insert(i *dbtoes.Index) error {

	//处理特殊字段
	data := u.handleSpecialFields(i)

	// interface转成struct
	var newData StudentScoreUser

	utils.Interface2Struct(data, &newData)

	// 插入新document
	Resp, err := i.TypedESConn.Index(StudentScoreIdx1).Request(newData).Do(context.Background())
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

	// ES index原数据
	originData, ok := i.PrimarySource.(*StudentScoreUser)
	// 没有查找主表信息的，直接嘎
	if ok == false {
		panic(errorpkg.ErrLogic)
	}

	// 处理msg特殊字段
	data := u.handleSpecialFields(i)

	// interface转成struct
	var newData StudentScoreUser
	utils.Interface2Struct(data, &newData)

	// 将子表内容赋值
	newData.ClassInfo = originData.ClassInfo
	newData.ScoreInfo = originData.ScoreInfo

	Resp, err := i.TypedESConn.Update(StudentScoreIdx1, i.PrimaryID).Doc(newData).Do(context.Background())
	if err != nil {
		// TODO::这里根据业务补救，这不一一写了(直接panic出去，通过rocketmq的重试特性。重新消费)
		panic(err)
	}
	i.SetPrimaryID(Resp.Id_)
	i.SetPrimarySource(newData)
	return nil
}

func (u *StudentScoreUser) handleSpecialFields(i *dbtoes.Index) map[string]interface{} {
	// 获取data数据

	var test []string

	// label字段
	if l, ok := i.BodyFirstData["label"].(string); ok {
		json.Unmarshal([]byte(l), &test)
		i.BodyFirstData["label"] = test
	}

	//其他字段

	return i.BodyFirstData
}
