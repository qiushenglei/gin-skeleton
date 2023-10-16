package validatorx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
)

func Validate(ctx context.Context, body interface{}) error {
	// 判断gin.context
	c, ok := ctx.(*gin.Context)
	if !ok {
		return errorpkg.ErrNoGinContext
	}

	// 绑定struct
	err := c.ShouldBind(body)
	if err != nil {
		return errorpkg.NewBizErrx(errorpkg.CodeBodyBind, err.Error())
	}

	// validator验证
	v := validator.New()
	err = v.Struct(body)
	if err != nil {
		panic(err)
	}
	return nil
}
