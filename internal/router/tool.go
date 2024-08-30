package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/controller"
)

func GenToolRouter(r *gin.RouterGroup) {
	tool := r.Group("/tools")
	{
		tool.GET("", controller.ToolController{}.ToolListController)
		tool.POST("", controller.ToolController{}.ToolAddController)
	}
}
