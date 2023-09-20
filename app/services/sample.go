package services

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/app/data"
	"github.com/qiushenglei/gin-skeleton/app/entity"
	"time"

	"github.com/gin-gonic/gin"
)

var (
// _ errors.Error
)

// SampleService ...
type SampleService struct {
	ctx *gin.Context
}

// NewSampleService returns a new SampleServiceService
func NewSampleService(ctx *gin.Context) *SampleService {
	return &SampleService{ctx: ctx}
}

// SetKeyValue ...
func (s *SampleService) SetKeyValue(ctx context.Context, req *entity.SampleSetRequest) (ret interface{}, err error) {
	// 提取参数
	key := req.Key
	value := req.Value

	// 设置 Redis Key-Values pair
	// 暂时不需要用到 context 来携带额外信息，比如超时时间，所以使用 `context.TODO()` 来生成一个空的 context
	// 为了演示目的，设置一个 60 秒的过期时间就行了，不需要永久有效
	data.RedisClient.Set(ctx, key, value, 60*time.Second)

	// RET
	//return entity.EmptyData, err
	return nil, err
}

// GetKeyValue ...
func (s *SampleService) GetKeyValue(ctx context.Context, req *entity.SampleGetRequest) (ret *entity.SomeData, err error) {
	// 提取参数
	key := req.Key

	// 获取 Redis Key-Values pair
	// 暂时不需要用到 context 来携带额外信息，比如超时时间，所以使用 `context.TODO()` 来生成一个空的 context
	data.RedisClient.Get(ctx, key)
	if err != nil {
		//return nil, errors.New(constants.RedisNotFoundCode, "", constants.RedisNotFoundMessage, nil)
		return nil, nil
	}

	// RET
	return &entity.SomeData{Value: nil}, err
}
