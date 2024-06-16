package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/middleware"
)

func Router() *gin.Engine {
	r := gin.Default() // 初始化gin实例
	r.Use(middleware.CorsMiddleware())
	GenSwaggerRouters(r)
	GenToolRouters(r)
	return r // 返回指针，供外部调用
}
