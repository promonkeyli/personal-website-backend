package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_backend.com/m/v2/internal/app/model"
	"web_backend.com/m/v2/utils"
)

type AuthController struct {
}

// @Summary		用户登录
// @Description	使用用户名密码进行登录
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			user	body		model.User	true	"用户名密码登录"
// @Success		200		{object}	utils.Response
// @Failure		400		{object}	string	"请求错误！"
// @Failure		500		string		"服务器错误！"
// @Router			/login [post]
func (T AuthController) AuthLoginController(c *gin.Context) {
	var bodyUser model.User
	if err := c.BindJSON(&bodyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB, user := model.QuerySingleUser(bodyUser.UserName)
	fmt.Println(bodyUser)
	if DB.Error != nil {
		HandleError(c, utils.StatusInternalServerError, "DB Error")
	} else {
		p := bodyUser.Password
		if p == user.Password {
			// 签发token
			token, e := utils.GenerateToken(p)
			if e != nil {
				HandleError(c, utils.StatusInternalServerError, "Token Error")
			} else {
				HandleOk(c, utils.StatusOK, utils.StatusOK.String(), token)
			}

		} else {
			HandleError(c, utils.StatusBadRequest, "密码错误，请重新输入！")
		}

	}

}

// @Summary		用户注销
// @Description	用户注销，清除会话、注销令牌
// @Tags			用户
// @Accept			json
// @Produce		json
// @Success		200	{object}	utils.Response
// @Failure		400	string		"参数错误！"
// @Failure		500	string		"服务器错误！"
// @Router			/logout [post]
func (T AuthController) AuthLogOutController(c *gin.Context) {
	//todo 注销令牌，此处后续添加注销逻辑，JWT退出注销给你需要借助外力实现
	HandleOk(c, utils.StatusOK, "注销成功！", nil)
}
