package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/controller"
)

// GenToolRouters 后台-工具页面路由
func GenToolRouters(r *gin.Engine) {
	tool := r.Group("/tool")
	{
		tool.GET("/list", controller.ToolController{}.ToolListController)
		tool.POST("/add", controller.ToolController{}.ToolAddController)
	}
}
