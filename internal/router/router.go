package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default() // 初始化gin实例
	GenToolRouters(r)
	return r // 返回指针，供外部调用
}
