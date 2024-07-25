package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"web_backend.com/m/v2/internal/app/repository"
	"web_backend.com/m/v2/internal/router"
)

func main() {
	repository.ConnectDB()
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	r := router.Router()
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
