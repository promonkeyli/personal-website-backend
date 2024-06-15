package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"web_backend.com/m/v2/internal/app/model"
)

type ToolController struct {
}

func (T ToolController) ToolListController(c *gin.Context) {
	toolList := model.QueryToolList()
	HandleOk(c, 200, true, "ok",
		toolList, time.Now())
}

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
