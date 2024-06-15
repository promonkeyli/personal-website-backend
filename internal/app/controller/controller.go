package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

type JsonStruct struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Message   interface{} `json:"message"`
	Data      interface{} `json:"data"`
	TimeStamp time.Time   `json:"timeStamp"`
}

// HandleOk 成功时调用
func HandleOk(c *gin.Context, code int, success bool, message interface{}, data interface{}, timeStamp time.Time) {
	json := &JsonStruct{Code: code, Success: success, Message: message, Data: data, TimeStamp: timeStamp}
	c.JSON(200, json)
}

// HandleError 失败时调用，失败时Data为{}
func HandleError(c *gin.Context, code int, success bool, message interface{}, timeStamp time.Time) {
	json := &JsonStruct{Code: code, Success: success, Message: message, TimeStamp: timeStamp}
	c.JSON(200, json)
}
