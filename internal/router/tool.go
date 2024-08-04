package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/controller"
)

func GenToolRouter(r *gin.RouterGroup) {
	tool := r.Group("/tool")
	{
		tool.GET("/list", controller.ToolController{}.ToolListController)
		tool.POST("/add", controller.ToolController{}.ToolAddController)
	}
}
