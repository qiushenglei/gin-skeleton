package student_score_idx

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
)

// StudentScoreIdx es index name
const StudentScoreIdx = "student_score_idx"

// ForeignKey 外键
const ForeignKey = "student_id"

type A interface {
	Gan()
}

var tableMap = map[string]A{
	"StudentScoreUser":  &StudentScoreUser{},
	"StudentScoreClass": &StudentScoreClass{},
	"StudentScoreScore": &StudentScoreScore{},
}

type StudentScoreStruct struct {
	Id         int      `json:"id"`
	Username   string   `json:"username"`
	Label      []string `json:"label"`
	ClassId    int      `json:"class_id"`
	StudentId  string   `json:"student_id"`
	AddTime    string   `json:"add_time"`
	UpdateTime string   `json:"update_time"`
	IsDeleted  int      `json:"is_deleted"`
	ClassInfo  struct {
		ClassId    int    `json:"class_id"`
		ClassName  string `json:"class_name"`
		Grade      int    `json:"grade"`
		AddTime    string `json:"add_time"`
		UpdateTime string `json:"update_time"`
	} `json:"class_info"`
	ScoreInfo []struct {
		Id         int    `json:"id"`
		StudentId  string `json:"student_id"`
		SubjectId  int    `json:"subject_id"`
		Score      int    `json:"score"`
		AddTime    string `json:"add_time"`
		UpdateTime string `json:"update_time"`
	} `json:"score_info"`
}

type StudentScoreSync struct {
}

func NewStudentScoreSync() *StudentScoreSync {
	return &StudentScoreSync{}
}

func (s *StudentScoreSync) UpdateSync(i *dbtoes.Index) error {
	return nil
}

func (s *StudentScoreSync) FindPrimaryTable(i *dbtoes.Index) error {
	// 不是主表并且没有主键，不需要同步到es
	if !i.IsPrimaryTable {
		if _, ok := i.BodyJson[i.ForeignKey]; !ok {
			return errorpkg.NewBizErrx(dbtoes.CodeSyncNoLoop, "更新的表没有主键，不需要同步到es")
		}
	}

	q1 := types.NewQuery()
	q1.Term = map[string]types.TermQuery{
		i.ForeignKey: types.TermQuery{Value: i.BodyJson[i.ForeignKey]},
	}

	request := &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must: []types.Query{
					*q1,
				},
			},
		},
		Sort: make([]types.SortCombinations, 0),
	}
	typeResp, err := i.TypedESConn.Search().Index(StudentScoreIdx).Request(request).Size(5).From(0).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if typeResp.Hits.Total.Value < 1 {
		return errorpkg.NewBizErrx(errorpkg.CodeFalse, "no found primary table")
	}

	if typeResp.Hits.Total.Value > 1 {
		panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "is not unique, fatal"))
	}

	var dataStruct StudentScoreStruct
	err = json.Unmarshal(typeResp.Hits.Hits[0].Source_, &dataStruct)
	if err != nil {
		panic(err)
	}
	i.PrimaryID = typeResp.Hits.Hits[0].Id_
	i.PrimarySource = &dataStruct
	return nil
}

func (s *StudentScoreSync) InsertSync(i *dbtoes.Index) error {

	s.ReflectCallMethod(i, "Insert")
	return nil
}

func (s *StudentScoreSync) GetIdxStructName(i *dbtoes.Index) string {
	table := cases.Title(language.Und).String(i.SyncTable)
	return "StudentScore" + table
}

func (s *StudentScoreSync) ReflectCallMethod(i *dbtoes.Index, MethodName string) {

	// 通过表名回去struct的名字
	structName := s.GetIdxStructName(i)

	// 获取实例
	o, ok := tableMap[structName]
	if ok != true {
		panic("define map false")
	}

	// 通过reflect动态调用表结构体的方法
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Pointer {
		//elem := reflect.Indirect(v)
		//if elem.Kind() == reflect.Struct {
		f := v.MethodByName(MethodName)
		if f.IsValid() == true && f.Kind() == reflect.Func {
			params := []reflect.Value{
				reflect.ValueOf(i),
			}
			f.Call(params)
		}
		//}
	}
}
