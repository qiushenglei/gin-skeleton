package models

import (
	"context"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/rw_isolate/query"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"hash/crc64"
	"strconv"
)

var OrderTablePower = 2 //2^TablePower = 表数量。用的map的分表和扩容方法

func GetTableName(AppID string) string {
	// 这里按位与
	var num int
	hash := crc64.New(crc64.MakeTable(crc64.ISO))
	if _, err := hash.Write([]byte(AppID)); err != nil {
		panic(errorpkg.ErrGetTableName)
	} else {
		num = int(hash.Sum64()) & (1<<OrderTablePower - 1) // 相当于 hash & 2^orderTablePower，例如 power是2 &运算就是获取低2位，低2位的范围是0-3(00000000, 00000001, 00000010, 00000011)
		num = num + 1                                      // 表命是从1开始的，所以加了1
	}
	return "order" + strconv.Itoa(num)
}

func FindAll(c context.Context, request *entity.FindOrderRequest) ([]*model.Order1, int64, error) {
	if request.AppID == "" {
		panic(errorpkg.ErrParam)
	}
	var res []*model.Order1
	var count int64
	err := query.Q.Transaction(func(tx *query.Query) error {
		order := tx.Order1
		var err error
		res, count, err = order.Table(GetTableName(request.AppID)).WithContext(c).
			Where(order.AppID.Eq(request.AppID)).FindByPage((*request.Page-1)*(*request.PageSize), *request.Page)
		if err != nil {
			panic(err)
		}
		return nil
	})

	return res, count, err
}
