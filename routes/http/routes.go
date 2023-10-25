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
	r.Use(middleware2.LogRequest())    // 记录请求信息
	r.Use(middleware2.GlobalRecover()) // 全局recover,抛出异常也需要记录请求日志，所以放后面

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

	// student
	studentAPI := r.Group("/student/", middleware2.AuthRequest())
	studentAPI.POST("/set_data", controller.SetData)      // 设置 Redis Key
	studentAPI.POST("/get_data", controller.GetData)      // 查询 Redis Key
	studentAPI.POST("/get_es_data", controller.GetESData) // 查询 Redis Key
	studentAPI.POST("/delay_msg", controller.Delay)       // 查询 Redis Key

	// upload
	UploadAPI := r.Group("/upload")
	UploadAPI.POST("/uploadImg", controller.UploadImg)

	// login test
	r.POST("/login", controller.Login)
	loginAPI := r.Group("/login/", middleware2.AuthRequest())
	{
		loginAPI.POST("/logged", controller.Logged)

	}

	// RET
	return
}
