package student_score_idx

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
)

type StudentScoreClass struct {
	ClassId    int    `json:"class_id,string"`
	ClassName  string `json:"class_name"`
	Grade      int    `json:"grade,string"`
	AddTime    string `json:"add_time"`
	UpdateTime string `json:"update_time"`
}

func (u *StudentScoreClass) Insert(i *dbtoes.Index) error {
	// ES index原数据
	originData, ok := i.PrimarySource.(*StudentScoreUser)
	// 没有查找主表信息的，直接嘎
	if ok == false {
		panic(errorpkg.ErrLogic)
	}

	// 处理msg特殊字段
	data := u.handleSpecialFields(i)

	// interface转成struct
	var newData StudentScoreClass
	utils.Interface2Struct(data, &newData)

	originData.ClassInfo = &newData

	Resp, err := i.TypedESConn.Update(StudentScoreIdx, i.PrimaryID).Doc(originData).Do(context.Background())
	if err != nil {
		// TODO::这里根据业务补救，这不一一写了(直接panic出去，通过rocketmq的重试特性。重新消费)
		panic(err)
	}
	i.SetPrimaryID(Resp.Id_)
	i.SetPrimarySource(originData)
	return nil
}

func (u *StudentScoreClass) Update(i *dbtoes.Index) error {
	// ES index原数据
	originData, ok := i.PrimarySource.(*StudentScoreUser)
	// 没有查找主表信息的，直接嘎
	if ok == false {
		panic(errorpkg.ErrLogic)
	}

	// 处理msg特殊字段
	data := u.handleSpecialFields(i)

	// interface转成struct
	var newData StudentScoreClass
	utils.Interface2Struct(data, &newData)

	originData.ClassInfo = &newData

	Resp, err := i.TypedESConn.Update(StudentScoreIdx, i.PrimaryID).Doc(originData).Do(context.Background())
	if err != nil {
		// TODO::这里根据业务补救，这不一一写了(直接panic出去，通过rocketmq的重试特性。重新消费)
		panic(err)
	}
	i.SetPrimaryID(Resp.Id_)
	i.SetPrimarySource(originData)
	return nil
}

func (u *StudentScoreClass) handleSpecialFields(i *dbtoes.Index) map[string]interface{} {
	// 获取data数据

	return i.BodyFirstData
}
