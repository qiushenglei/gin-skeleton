package example

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"io"
)

const StudentScoreIdx = "student_score_idx"

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
	if i.IsPrimaryTable == false {
		if _, ok := i.BodyJson[i.ForeignKey]; ok == false {
			// errors.New("没有主键，不需要更改")
			return nil
		}
	}

	// planA ESConn client
	query :=
		`{
		  "query": {
			"bool": {
			  "must": [
				{"term": {
				  "student_id": {
					"value": "2"
				  }
				}}
			  ]
			}
		  }
		}`

	resp, err := i.ESConn.Search(
		i.ESConn.Search.WithIndex(StudentScoreIdx),
		i.ESConn.Search.WithBody(bytes.NewBufferString(query)),
	)
	_, err = io.ReadAll(resp.Body)

	// plan B
	q1 := types.NewQuery()
	q1.Term = map[string]types.TermQuery{
		"student_id": types.TermQuery{Value: 2},
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
		return errors.New("no find primary table, loop")
	}

	//if typeResp.Hits.Total.Value > 1 {
	//	panic(errors.New("is not unique, fatal"))
	//}

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

	dataStruct, ok := i.PrimarySource.(*StudentScoreStruct)
	if ok == false {
		panic(11)
	}
	dataStruct.AddTime = "2020-03-10 12:20:31"
	// plan B
	q1 := types.NewQuery()
	q1.Term = map[string]types.TermQuery{
		"_id": types.TermQuery{Value: i.PrimaryID},
	}

	Resp, err := i.TypedESConn.Update(StudentScoreIdx, i.PrimaryID).Doc(dataStruct).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(Resp.Id_)

	//i.TypedESConn.Index(StudentScoreIdx).Request().Do(context.Background())

	return nil
}
