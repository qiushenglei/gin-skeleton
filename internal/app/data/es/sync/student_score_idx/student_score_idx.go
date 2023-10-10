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

type TableSync interface {
	Insert(i *dbtoes.Index) error
	Update(i *dbtoes.Index) error
}

var tableMap = map[string]TableSync{
	"StudentScoreUser":  &StudentScoreUser{},
	"StudentScoreClass": &StudentScoreClass{},
	"StudentScoreScore": &StudentScoreScore{},
}

type StudentScoreSync struct {
}

func NewStudentScoreSync() *StudentScoreSync {
	return &StudentScoreSync{}
}

// FindPrimaryTableByPForeignKey 通过外键查主表信息
func (s *StudentScoreSync) FindPrimaryTableByPForeignKey(i *dbtoes.Index) error {
	// 不是主表并且没有主键，不需要同步到es
	if !i.IsPrimaryTable {
		if _, ok := i.BodyFirstData[i.ForeignKey]; !ok {
			return errorpkg.NewBizErrx(dbtoes.CodeSyncNoLoop, "更新的表没有主键，不需要同步到es")
		}
	}

	q1 := types.NewQuery()
	q1.Term = map[string]types.TermQuery{
		i.ForeignKey: types.TermQuery{Value: i.BodyFirstData[i.ForeignKey]},
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

	var dataStruct StudentScoreUser
	err = json.Unmarshal(typeResp.Hits.Hits[0].Source_, &dataStruct)
	if err != nil {
		panic(err)
	}
	i.PrimaryID = typeResp.Hits.Hits[0].Id_
	i.PrimarySource = &dataStruct
	return nil
}

// InsertSync insert document
func (s *StudentScoreSync) InsertSync(i *dbtoes.Index) error {
	err := s.ReflectCallMethod(i, "Insert")
	return err
}

// UpdateSync update document
func (s *StudentScoreSync) UpdateSync(i *dbtoes.Index) error {
	s.ReflectCallMethod(i, "Update")
	return nil
}

// GetIdxStructName 获取index对应的结构体名称
func (s *StudentScoreSync) GetIdxStructName(i *dbtoes.Index) string {
	table := cases.Title(language.Und).String(i.SyncTable)
	return "StudentScore" + table
}

// ReflectCallMethod 通过反射调用 修改表 的 指定函数名方法
func (s *StudentScoreSync) ReflectCallMethod(i *dbtoes.Index, MethodName string) error {

	// 通过表名回去struct的名字
	structName := s.GetIdxStructName(i)

	// 获取实例
	o, ok := tableMap[structName]
	if ok != true {
		panic("define map false")
	}

	// 通过reflect动态调用表结构体的方法
	//var res []reflect.Value
	v := reflect.ValueOf(o)

	// 因为调用的方法都是指针类型的结构体方法，所以map定义成了&struct，这里的类型就成了Pointer，不知道怎么改
	if v.Kind() == reflect.Pointer {
		f := v.MethodByName(MethodName)
		if f.IsValid() == true && f.Kind() == reflect.Func {
			params := []reflect.Value{
				reflect.ValueOf(i),
			}
			res := f.Call(params)

			// 方法默认都返回error
			if v, ok := res[0].Interface().(error); ok {
				return v
			}
		}
	}
	return nil
}
