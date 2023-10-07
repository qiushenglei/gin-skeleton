package routes

import (
	"github.com/qiushenglei/gin-skeleton/internal/app/controllers"
	middleware2 "github.com/qiushenglei/gin-skeleton/internal/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes defined to a new gin.Engine, and returns the pointer to the engine
func RegisterRoutes() (r *gin.Engine) {

	// 新建一个 gin Engine
	r = gin.Default()
	// temporary workaround to cease the warning: about trusted proxy https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
	r.SetTrustedProxies(nil)

	// 全局中间件
	//r.Use(cors.CORS())       // 跨域中间件
	//r.Use(logging.Logging()) // 日志中间件
	r.Use(middleware2.BindRequestID()) // 请求ID
	r.Use(middleware2.GlobalRecover()) // 全局recover
	r.Use(middleware2.LogRequest())    // 记录请求信息

	// 探针
	r.GET("/", func(ctx *gin.Context) { ctx.String(http.StatusOK, "Hello, this is Gin-Awesome! ") })
	// 初始接口
	r.POST("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") }) // ping
	//r.POST("/health/check", controller.HealthCheck)                           // 健康检查

	// 演示接口
	sampleAPI := r.Group("/sample/")
	{
		sampleAPI.POST("/set_key_value", controller.SetKeyValue) // 设置 Redis Key
		sampleAPI.POST("/get_key_value", controller.GetKeyValue) // 查询 Redis Key
	}

	// RET
	return
}
