package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"go.uber.org/zap"
	"io"
)

type LoggingWriter struct {
	gin.ResponseWriter
	responseBody *bytes.Buffer
}

// Write ...
func (w LoggingWriter) Write(p []byte) (int, error) {
	if n, err := w.responseBody.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []byte
		if c.Request.Body != nil {

			// Reader里获取[]byte
			// 这种方式需要知道长度
			//requestBody := make([]byte, 4096)
			//n, err := c.Request.Body.Read(requestBody)

			//从body读出来
			var err error
			requestBody, err = io.ReadAll(c.Request.Body)
			if err != nil {
				panic(fmt.Errorf("logger middleware get request responseBody err"))
			}

			//写回body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
			if err != nil {
				panic(fmt.Errorf("logger middleware write request responseBody err"))
			}
		}

		// gin.ResponseWriter是没有body这个字段的
		// gin框架写进了writer后是没有reader方法读出body内容的
		// 所以在gin框架写入(调用c.ResponseWriter.Write(p))前，把内容存一份到这个body结构体变量里。
		// 这个结构体就完全复制了c.Writer,只是多加了一个Body字段，所以里面的c.Writer的逻辑基本没有变，只是多了上面LoggingWriter.Write方法的写byteBuffer的操作
		logWriter := &LoggingWriter{
			c.Writer,
			bytes.NewBufferString(""),
		}
		c.Writer = logWriter

		// 添加自定义的GlobalRecover
		defer func() {
			if r := recover(); r != nil {
				var err error
				var ok bool
				if err, ok = r.(error); ok {
					utils.Response(c, nil, err)
				}
				if r, ok := r.(string); ok {
					utils.Response(c, nil, errorpkg.NewIOErrx(errorpkg.CodeFalse, r))
				}

				// 异常记录日志
				var rb map[string]interface{}
				var requestBodyStr string
				if len(requestBody) > 0 {
					requestBodyStr = string(utils.JsonBytes2String(requestBody, &rb))
				}

				logs.Log.Error(
					c,
					zap.String("requestBody", requestBodyStr),
					zap.String("responseBody", logWriter.responseBody.String()),
				)

				// 直接跳出中间件的循环了，不然又会走到下一个中间件或者controller
				c.Abort()
			}
		}()

		c.Next()

		// 正常记录日志
		var rb map[string]interface{}
		// requestBody内容里有tab和\r\n,可以用jsondecode和encode过滤掉
		requestBodyBytes := utils.JsonBytes2String(requestBody, &rb)
		logs.Log.Info(
			c,
			zap.Any("requestBody", string(requestBodyBytes)),
			zap.String("responseBody", logWriter.responseBody.String()),
		)
	}
}
