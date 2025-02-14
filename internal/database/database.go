// internal/database/database.go: 数据库连接
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web_backend.com/m/v2/internal/config"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsn := config.GetDSN()
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
}

// 通过函数安全访问数据库实例
func GetDB() *gorm.DB {
	return db
}
