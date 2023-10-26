package middleware

import (
	"github.com/gin-gonic/gin"
)

func GlobalRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里的代码，我扔到httplogging中间件了，因为接到异常后，还是需要记录http日志。
		// 所以不能直接调用c.Abort().
		//defer func() {
		//	if r := recover(); r != nil {
		//		logs.Log.Error(context.Background())
		//		if r, ok := r.(error); ok {
		//			utils.Response(c, nil, r)
		//		}
		//		if r, ok := r.(string); ok {
		//			utils.Response(c, nil, errorpkg.NewIOErrx(errorpkg.CodeFalse, r))
		//		}
		//		// 这里没法调用c.Abort()，
		//		panic(r)
		//	}
		//}()
		//c.Next()
		//
		//fmt.Println("如果不正常是不会到这里的")
	}
}
