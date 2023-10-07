package example

import (
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"io"
)

const StudentScoreIdx = "student_score_idx"

type StudentScoreSync struct {
}

func NewStudentScoreSync() *StudentScoreSync {
	return &StudentScoreSync{}
}

func (s *StudentScoreSync) InsertSync() error {
	return nil
}

func (s *StudentScoreSync) UpdateSync() error {
	return nil
}

func (s *StudentScoreSync) FindPrimaryTable(i *dbtoes.Index) error {
	if i.IsPrimaryTable == false {
		if _, ok := i.BodyJson[i.ForeignKey]; ok == false {
			// errors.New("没有主键，不需要更改")
			return nil
		}
	}

	query := `{
	  query: {
		bool: {
		  must: [
			{term: {
			  student_id: {
				value: 2
			  }
			}}
		  ]
		}
	  }
	}`

	resp, err := i.ES.Search(
		i.ES.Search.WithIndex(StudentScoreIdx),
		i.ES.Search.WithQuery(query),
	)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	return nil
}
