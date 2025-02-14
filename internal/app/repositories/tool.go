package repositories

import (
	"web_backend.com/m/v2/internal/app/models"
	"web_backend.com/m/v2/internal/database"
)

func QueryToolList() []models.Tool {
	var tool []models.Tool
	database.GetDB().Find(&tool)
	return tool
}

func CreateToolItem(tool models.Tool) {
	database.GetDB().Create(&tool)
}
