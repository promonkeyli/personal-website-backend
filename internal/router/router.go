package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/middleware"
)

func Router() *gin.Engine {
	// gin.default 默认包含了日志记录功能
	r := gin.Default()
	// 设置信任的端口
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil
	}
	// 启用允许跨域中间件
	r.Use(middleware.CorsMiddleware())
	// 路由生成
	generateAllRouter(r)
	return r
}

func generateAllRouter(r *gin.Engine) {
	api := r.Group("/api")

	// v1 版本
	{
		v1 := api.Group("/v1")
		// 不需要授权
		GenSwaggerRouter(r)
		GenLoginRouter(v1)
		// 需要授权，使用自定义JWT中间件
		v1.Use(middleware.JWT())
		GenAuthRouter(v1)
		GenUserRouter(v1)
		GenToolRouter(v1)
	}

}
