package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 设置信任的端口
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil
	}
	r.Use(middleware.CorsMiddleware())
	generateAllRouter(r)
	return r
}

func generateAllRouter(r *gin.Engine) {
	api := r.Group("/api")
	// 不需要授权
	GenSwaggerRouter(r)
	GenLoginRouter(api)
	// 需要授权
	api.Use(middleware.JWT())
	GenAuthRouter(api)
	GenToolRouter(api)
}
