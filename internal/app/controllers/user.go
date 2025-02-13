package controllers

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/models"
	"web_backend.com/m/v2/tools"
)

type UserController struct {
}

//	UserListController
//
// @Summary		用户列表
// @Description	获取所有用户
// @Tags			user
// @Accept			json
// @Produce		json
// @Success		200	{object}	tools.Response
// @Failure		400	string		"参数错误！"
// @Failure		500	string		"服务器错误！"
// @Router			/users [get]
// @Security		ApiKeyAuth
func (T UserController) UserListController(c *gin.Context) {
	DB, userList := models.QueryAllUserList()
	if DB.Error != nil {
		HandleError(c, tools.StatusInternalServerError, DB.Error.Error())
	} else {
		HandleOk(c, tools.StatusOK, tools.StatusOK.String(), userList)
	}
}

//	UserAddController
//
// @Summary		新增用户
// @Description	新增用户
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			user	body		models.User	true	"用户新增"
// @Success		200		{object}	tools.Response
// @Failure		400		string		"参数错误！"
// @Failure		500		string		"服务器错误！"
// @Router			/users [post]
// @Security		ApiKeyAuth
func (T UserController) UserAddController(c *gin.Context) {
	HandleOk(c, tools.StatusOK, tools.StatusOK.String(), nil)
}
