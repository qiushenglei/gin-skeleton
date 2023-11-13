package validatorx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	var err error
	if c.Request.Method == "GET" {
		err = c.ShouldBind(body)
	} else {
		err = c.ShouldBindBodyWith(body, binding.JSON)
	}
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

func ValidateRPC(body interface{}) error {
	// validator验证
	v := validator.New()
	err := v.Struct(body)
	if err != nil {
		panic(err)
	}
	return nil
}
