package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "web_backend.com/m/v2/api/swagger"
	"web_backend.com/m/v2/internal/app/middleware"
)

func main() {
	r := gin.Default() // 初始化gin实例
	r.Use(middleware.CorsMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
