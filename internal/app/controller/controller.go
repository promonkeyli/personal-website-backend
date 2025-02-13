package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"web_backend.com/m/v2/tools"
)

// HandleOk 成功时调用
func HandleOk(c *gin.Context, code tools.StatusCode, message string, data interface{}) {
	json := &tools.Response{Code: code, Message: message, Data: data, Time: time.Now().UnixMilli()}
	c.JSON(int(code), json)
}

// HandleError 失败时调用，失败时Data为 nil
func HandleError(c *gin.Context, code tools.StatusCode, message string) {
	json := &tools.Response{Code: code, Message: message, Data: nil, Time: time.Now().UnixMilli()}
	c.JSON(int(code), json)
}
