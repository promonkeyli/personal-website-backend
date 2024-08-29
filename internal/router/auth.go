package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/controller"
)

func GenAuthRouter(r *gin.RouterGroup) {
	r.POST("/logout", controller.ToolController{}.ToolAddController)
}

func GenLoginRouter(r *gin.RouterGroup) {
	r.POST("/login", controller.AuthController{}.AuthLoginController)
}
