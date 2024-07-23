package main

import (
	"web_backend.com/m/v2/internal/app/repository"
	"web_backend.com/m/v2/internal/router"
)

func main() {
	repository.ConnectDB()
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
