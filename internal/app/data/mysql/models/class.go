package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/canal_test/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/canal_test/query"

	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"strconv"
	"time"
)

func NewClass(class *entity.ClassInfo) *model.Class {
	if class == nil {
		panic(errorpkg.ErrLogic)
	}
	now := time.Now()
	return &model.Class{
		ID:         uint32(class.ClassId),
		ClassName:  class.ClassName,
		Grade:      strconv.Itoa(class.Grade),
		AddTime:    now,
		UpdateTime: &now,
	}
}

func GetClassBySearchCond(c *gin.Context, cond *entity.SearchCond) (*model.Class, error) {
	if cond.ClassId == 0 && cond.ClassName == "" {
		// 不需要查询
		return nil, nil
	}

	var res *model.Class
	q := query.Q.Class.WithContext(c).Select()
	if cond.ClassId != 0 {
		q.Where(query.Q.Class.ID.Eq(uint32(cond.ClassId)))
	}
	if cond.ClassName != "" {
		q.Where(query.Q.Class.ClassName.Like(fmt.Sprintf("%s%%", cond.ClassName)))
	}
	if err := q.Scan(&res); err != nil {
		panic(err)
	}

	if res.ID == 0 {
		// 没查到数据
		return res, errorpkg.NewBizErrx(errorpkg.CodeFalse, "subject_name is not define")
	}
	return res, nil
}
