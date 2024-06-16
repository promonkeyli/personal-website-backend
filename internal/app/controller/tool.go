package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"web_backend.com/m/v2/internal/app/model"
)

type ToolController struct {
}

//	@Summary	我的工具信息列表获取
//	@Produce	json
//	@Success	200	{object}	model.Tool	"请求成功"
//	@Failure	400	{string}	string		"请求错误"
//	@Failure	500	{string}	string		"内部错误"
//	@Router		/tool/list [get]
func (T ToolController) ToolListController(c *gin.Context) {
	toolList := model.QueryToolList()
	HandleOk(c, 200, true, "ok",
		toolList, time.Now())
}

//	@Summary	我的工具项新增
//	@Produce	json
//	@Param		tool	body		model.Tool	true	"Tool object to be added"
//	@Success	200		{object}	model.Tool	"成功"
//	@Failure	400		{object}	string		"请求错误"
//	@Failure	500		{object}	string		"内部错误"
//	@Router		/tool/add [post]
func (T ToolController) ToolAddController(c *gin.Context) {
	// 接受json格式数据，使用map
	var tool model.Tool
	err := c.BindJSON(&tool)
	model.CreateToolItem(tool)
	if err == nil {
		HandleOk(c, 200, true, "ok",
			tool, time.Now())
		return
	}
	HandleError(c, 500, false, err, time.Now())

}
