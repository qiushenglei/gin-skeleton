package essearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/es/sync/student_score_idx"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"go.uber.org/zap"
)

const StudentScoreIdx = "student_score_idx"

type StudentScoreIdxSearch struct {
}

func (s *StudentScoreIdxSearch) Search(ctx context.Context, cond *entity.SearchCond) {

	request := &search.Request{
		From: &cond.Page,
		Size: &cond.PageSize,
	}

	filterQuery := s.buildFilterQuery(ctx, cond)
	var shouldQuery []types.Query
	var mustNoQuery []types.Query

	request.Query = &types.Query{
		Bool: &types.BoolQuery{
			Filter:  filterQuery,
			Should:  shouldQuery,
			MustNot: mustNoQuery,
		},
	}

	resp, err := data.TypedESClient.Search().Index(student_score_idx.StudentScoreIdx1).Request(request).Do(ctx)
	if err != nil {
		panic(err)
	}

	s.handleData(ctx, resp)
	return
}

func (s *StudentScoreIdxSearch) buildFilterQuery(ctx context.Context, cond *entity.SearchCond) []types.Query {
	var filterQuery []types.Query

	if query := s.searchStudentId(cond); query != nil {
		filterQuery = append(filterQuery, *query)
	}

	if query := s.searchClassId(cond); query != nil {
		filterQuery = append(filterQuery, *query)
	}
	if query := s.searchClassName(cond); query != nil {
		filterQuery = append(filterQuery, *query)
	}
	if query := s.searchUsername(cond); query != nil {
		filterQuery = append(filterQuery, *query)
	}
	if query := s.searchLabel(cond); query != nil {
		filterQuery = append(filterQuery, *query)
	}
	if query := s.searchScoreCond(cond); query != nil {
		for _, v := range query {
			filterQuery = append(filterQuery, *v)
		}
	}

	return filterQuery
}

// searchClassId 学号查询 keyword类型，term精准查询，不应该支持模糊和分词，所以用term
func (s *StudentScoreIdxSearch) searchStudentId(cond *entity.SearchCond) *types.Query {
	if cond.StudentId == "" {
		return nil
	}
	query := types.Query{
		Term: map[string]types.TermQuery{
			"class_id": {
				Value: cond.StudentId,
			},
		},
	}
	return &query
}

// searchClassId long类型，term精准查询
func (s *StudentScoreIdxSearch) searchClassId(cond *entity.SearchCond) *types.Query {
	if cond.ClassId == 0 {
		return nil
	}
	query := types.Query{
		Term: map[string]types.TermQuery{
			"class_info.class_id": {
				Value: cond.ClassId,
			},
		},
	}
	return &query
}

// searchClassName text类型，match可分词模糊查询
func (s *StudentScoreIdxSearch) searchClassName(cond *entity.SearchCond) *types.Query {
	if cond.ClassName == "" {
		return nil
	}
	query := types.Query{
		Match: map[string]types.MatchQuery{
			"class_info.class_name": {
				Query: cond.ClassName,
			},
		},
	}
	return &query
}

// searchUsername wildcard类型(是keywords的一种)，可分词模糊查询，这里用了query_string查询方法，是match查询的高级使用
func (s *StudentScoreIdxSearch) searchUsername(cond *entity.SearchCond) *types.Query {
	if cond.Username == "" {
		return nil
	}
	query := types.Query{
		QueryString: &types.QueryStringQuery{
			Query: fmt.Sprintf("*%s*", cond.Username),
		},
	}
	return &query
}

// searchLabel text类型，但是是数组形式，terms查询方法
func (s *StudentScoreIdxSearch) searchLabel(cond *entity.SearchCond) *types.Query {
	if len(cond.Label) == 0 {
		return nil
	}
	query := types.Query{
		Terms: &types.TermsQuery{
			TermsQuery: map[string]types.TermsQueryField{
				"label": cond.Label,
			},
		},
	}
	return &query
}

// searchScoreCond 两个字段都是
func (s *StudentScoreIdxSearch) searchScoreCond(cond *entity.SearchCond) []*types.Query {

	if cond.ScoreCond.SubjectName == "" || cond.ScoreCond.Score == nil {
		return nil
	}
	query := make([]*types.Query, 2)
	query[] = &types.Query{
		Term: map[string]types.TermQuery{
			"score_info.subject_info.subject_name.keyword": types.TermQuery{ //subject_name是text类型，因为这里需要精准查询，所以es的index/mapping中给找个字段加了keyword类型的子字段并取名名为keyword
				Value: cond.ScoreCond.SubjectName,
			},
		},
	}
	query[] = &types.Query{
		Term: map[string]types.TermQuery{
			"score_info.score": types.TermQuery{ //score分数，也需要精准匹配，map定义的是long
				Value: cond.ScoreCond.Score,
			},
		},
	}

	return query
}

func (s *StudentScoreIdxSearch) handleData(ctx context.Context, response *search.Response) {
	// 这里做字段的处理,我这里默认写成记录到日志里
	data, _ := json.Marshal(response)
	logs.Log.Info(ctx, zap.ByteString("response", data))
	return
}
