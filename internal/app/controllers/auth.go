package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/models"
	"web_backend.com/m/v2/internal/pkg/network"
	"web_backend.com/m/v2/tools"
)

type AuthController struct {
}

// @Summary		用户登录
// @Description	使用用户名密码进行登录
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user	body		models.User	true	"用户名密码登录"
// @Success		200		{object}	network.Response
// @Security		BearerAuth
// @Router			/login [post]
func (T AuthController) AuthLoginController(c *gin.Context) {
	// 请求参数判断
	var bodyUser models.User
	tools.LogInfo("AuthLoginController called")
	if err := c.BindJSON(&bodyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bodyUser)
	// user, err := repositories.QuerySingleUser(bodyUser.UserName)
	// fmt.Println(bodyUser)
	// if err != nil {
	// 	network.HandleError(c, network.StatusInternalServerError, "DB Error")
	// } else {
	// 	p := bodyUser.Password
	// 	if p == user.Password {
	// 		// 签发token
	// 		token, e := tools.GenerateToken(p)
	// 		if e != nil {
	// 			network.HandleError(c, network.StatusInternalServerError, "Token Error")
	// 		} else {
	// 			network.HandleOk(c, network.StatusOK, network.StatusOK.String(), token)
	// 		}

	// 	} else {
	// 		network.HandleError(c, network.StatusBadRequest, "密码错误，请重新输入！")
	// 	}

	// }

}

// @Summary		用户注销
// @Description	用户注销，清除会话、注销令牌
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		200	{object}	network.Response
// @Security		BearerAuth
// @Router			/logout [post]
func (T AuthController) AuthLogOutController(c *gin.Context) {
	//todo 注销令牌，此处后续添加注销逻辑，JWT退出注销给你需要借助外力实现
	network.HandleOk(c, network.StatusOK, "注销成功！", nil)
}
