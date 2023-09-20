package middleware

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
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
		c.Next()

		logs.Log.Info(
			context.Background(),
			string(requestBody),
			logWriter.responseBody.String(),
		)
	}
}
