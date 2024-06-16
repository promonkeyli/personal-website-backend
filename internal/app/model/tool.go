package model

import (
	"web_backend.com/m/v2/internal/app/repository"
)

// Tool represents a tool entity
// @Description Tool represents a tool entity
type Tool struct {
	ID       uint   `json:"id,omitempty" gorm:"column:id" swaggerignore:"true"`
	Category string `json:"category" gorm:"column:category"`
	Name     string `json:"name" gorm:"column:name"`
	Url      string `json:"url" gorm:"column:url"`
}

// TableName 自定义表名
func (Tool) TableName() string {
	return "tool"
}

func QueryToolList() []Tool {
	var tool []Tool
	repository.DB.Find(&tool)
	return tool
}

func CreateToolItem(tool Tool) {
	repository.DB.Create(&tool)
}
