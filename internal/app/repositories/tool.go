package repositories

import "web_backend.com/m/v2/internal/app/models"

func QueryToolList() []models.Tool {
	var tool []models.Tool
	DB.Find(&tool)
	return tool
}

func CreateToolItem(tool models.Tool) {
	DB.Create(&tool)
}
