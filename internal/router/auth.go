package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/controller"
)

func GenAuthRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/logout", controller.ToolController{}.ToolAddController)
	}
}

func GenLoginRouter(r *gin.Engine) {
	r.POST("/login", controller.AuthController{}.AuthLoginController)
}
