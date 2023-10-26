package models

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/query"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"hash/fnv"
	"strconv"
)

var OrderTablePower = 2 //2^TablePower = 表数量。用的map的分表和扩容方法

func GetTableName(AppID string) string {
	// 这里按位与
	var num int
	hash := fnv.New32()
	if _, err := fnv.New32().Write([]byte(AppID)); err != nil {
		panic(errorpkg.ErrGetTableName)
	} else {
		num = int(hash.Sum32()) & (1<<OrderTablePower - 1) // 相当于 hash & 2^orderTablePower  （math.Pow(2, OrderTablePower)）
	}
	return "order" + strconv.Itoa(num)
}

func FindAll(c context.Context, request *entity.FindOrderRequest) ([]*model.Order1, int64, error) {
	if request.AppID == "" {
		panic(errorpkg.ErrParam)
	}
	order := query.Q.Order1
	res, count, err := order.Table(GetTableName(request.AppID)).WithContext(c).
		Where(order.AppID.Eq(request.AppID)).FindByPage(*request.Page, *request.PageSize)
	if err != nil {
		panic(err)
	}
	return res, count, err
}
