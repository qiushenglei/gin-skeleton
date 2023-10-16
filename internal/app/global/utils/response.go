package utils

import (
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/constants"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response sends out a JSON response.
// If `err` is nil, `data` will be serialized and sent out with constants StatusOk Code and Message;
// if not nil, `err` will be serialized and sent out.
func Response(c *gin.Context, data interface{}, err error) {
	if err == nil {
		c.JSON(
			http.StatusOK,
			entity.DefaultResponse{
				Msg:  constants.StatusOkMessage,
				Data: data,
			},
		)
	} else {
		var code int
		var msg string
		if e, ok := err.(errorpkg.Errx); ok {
			code = e.Code()
			msg = e.Msg()
		} else {
			code = errorpkg.CodeFalse
			msg = err.Error()
		}
		c.JSON(
			http.StatusInternalServerError,
			entity.DefaultResponse{
				Code: code,
				Msg:  msg,
				Data: data,
			},
		)
	}
}

// errZapFields dismantles an gin-awesome-style error into different zap fields.
//func errZapFields(e errors.ErrorIface) []zap.Field {
//	ret := []zap.Field{
//		zap.Int("code", e.GetCode()),
//		zap.String("reason", e.GetReason()),
//		zap.String("message", e.GetMessage()),
//		zap.String("file", e.GetFileLine()),
//		zap.Any("data", e.GetData())}
//	if e.GetStack() != "" {
//		ret = append(ret, zap.String("stack", e.GetStack()))
//	}
//	if e.GetExtraDataMap() != nil {
//		for k, v := range e.GetExtraDataMap() {
//			ret = append(ret, zap.Any(k, v))
//		}
//	}
//	return ret
//}
