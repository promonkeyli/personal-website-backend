package main

import (
	"github.com/gin-gonic/gin"
	"os"
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

//	@title			个人网站后台接口文档
//	@version		0.0.1
//	@description	使用Golang编写的个人网站后台接口
//	@host			https:promonkeyli.top:8080
//	@BasePath		/api/v1
//	@schemes		https
func main() {
	r := router.Router()
	gin.DefaultWriter = os.Stdout
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
