package controller

import (
	"github.com/gin-gonic/gin"
	"web_backend.com/m/v2/internal/app/model"
	"web_backend.com/m/v2/tools"
)

type ToolController struct {
}

//	ToolListController
//
// @Summary		工具列表
// @Description	获取工具
// @Tags			tool
// @Accept			json
// @Produce		json
// @Success		200	{object}	tools.Response
// @Failure		400	string		"参数错误！"
// @Failure		500	string		"服务器错误！"
// @Router			/tools [get]
// @Security		ApiKeyAuth
func (T ToolController) ToolListController(c *gin.Context) {
	toolList := model.QueryToolList()
	HandleOk(c, tools.StatusOK, tools.StatusOK.String(), toolList)
}

//	ToolAddController
//
// @Summary		工具新增
// @Description	添加工具
// @Tags			tool
// @Accept			json
// @Produce		json
// @Param			tool	body		model.Tool	true	"工具新增"
// @Success		200		{object}	tools.Response
// @Failure		400		string		"参数错误！"
// @Failure		500		string		"服务器错误！"
// @Router			/tools [post]
// @Security		ApiKeyAuth
func (T ToolController) ToolAddController(c *gin.Context) {
	// 接受json格式数据，使用map
	var tool model.Tool
	err := c.BindJSON(&tool)
	model.CreateToolItem(tool)
	if err == nil {
		HandleOk(c, tools.StatusOK, tools.StatusOK.String(), tool)
		return
	}
	HandleError(c, tools.StatusInternalServerError, err.Error())
}
