package controller

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/model"
	"web_backend.com/m/v2/utils"
)

type UserController struct {
}

//	 UserListController
//	@Summary		用户列表
//	@Description	获取所有用户
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response
//	@Failure		400	string		"参数错误！"
//	@Failure		500	string		"服务器错误！"
//	@Router			/users [get]
//	@Security		ApiKeyAuth
func (T UserController) UserListController(c *gin.Context) {
	DB, userList := model.QueryAllUserList()
	if DB.Error != nil {
		HandleError(c, utils.StatusInternalServerError, DB.Error.Error())
	} else {
		HandleOk(c, utils.StatusOK, utils.StatusOK.String(), userList)
	}
}

//	 UserAddController
//	@Summary		新增用户
//	@Description	新增用户
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.User	true	"用户新增"
//	@Success		200		{object}	utils.Response
//	@Failure		400		string		"参数错误！"
//	@Failure		500		string		"服务器错误！"
//	@Router			/users [post]
//	@Security		ApiKeyAuth
func (T UserController) UserAddController(c *gin.Context) {
	HandleOk(c, utils.StatusOK, utils.StatusOK.String(), nil)
}
