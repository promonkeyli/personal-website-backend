package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	generateAllRouter(r)
	return r
}

func generateAllRouter(r *gin.Engine) {
	// 不需要权限认证的接口：（swagger，login）
	GenSwaggerRouter(r)
	GenLoginRouter(r)
	// 需要权限认证的接口
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middleware.JWT())
	GenAuthRouter(v1)
	GenToolRouter(v1)
}
