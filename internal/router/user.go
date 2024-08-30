package router

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/controller"
)

func GenUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	{
		user.GET("", controller.UserController{}.UserListController)
		user.POST("", controller.UserController{}.UserAddController)
	}
}
