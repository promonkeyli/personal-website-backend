package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/repository"
	"web_backend.com/m/v2/internal/router"
)

func init() {
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	repository.ConnectDB()
}

// @title						个人网站后台接口文档
// @version					0.0.1
// @description				Go编写的个人网站后台接口
// @securityDefinitions.apikey	ApiKeyAuth  API的认证方式
// @in							header 发送认证的方式
// @name						Authorization  后端获取认证值得方式
// @host						https:promonkeyli.top:8080
// @BasePath					/api/v1
// @schemes					https
//
// @tag.name					user
// @tag.description			用户相关的操作
//
// @tag.name					tool
// @tag.description			工具相关的操作
func main() {
	r := router.Router()
	gin.DefaultWriter = os.Stdout
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
