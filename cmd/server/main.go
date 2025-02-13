package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/repositories"
	"web_backend.com/m/v2/internal/pkg/middleware"
	"web_backend.com/m/v2/internal/router"
)

func init() {
	mode := os.Getenv("GIN_MODE") // docker build时会自动注入为release生产模式
	if mode == "" {
		mode = gin.DebugMode // 开发以及调试阶段使用：输出详细的调试信息，包括请求的详细日志、路由匹配信息等，有助于开发者在开发和调试阶段快速定位问题。性能较低
	}
	gin.SetMode(mode)
	repositories.ConnectDB()
}

//	@title			个人网站后台接口文档
//	@version		0.0.1
//	@description	Go编写的个人网站后台接口

//	@securityDefinitions.apikey
//	@name			ApiKeyAuth
//	@in				header
//	@description	此API采用API Key进行认证
//	@keyName		Authorization

//	@tag.name			user
//	@tag.description	用户相关的操作

// @tag.name			tool
// @tag.description	工具相关的操作
func main() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	router.Router(r)
	r.Run(":8081")
}
