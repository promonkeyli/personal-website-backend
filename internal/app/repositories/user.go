package repositories

import (
	"gorm.io/gorm"
	"web_backend.com/m/v2/internal/app/models"
	"web_backend.com/m/v2/internal/database"
)

// QuerySingleUser 单个用户查找
func QuerySingleUser(username string) (*gorm.DB, models.User) {
	var user models.User
	return database.GetDB().Where("username = ?", username).First(&user), user
}

// QueryAllUser 查找所有用户
func QueryAllUserList() (*gorm.DB, []models.User) {
	var user []models.User
	return database.GetDB().Find(&user), user
}
