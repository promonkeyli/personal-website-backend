package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web_backend.com/m/v2/config"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	DB, err = gorm.Open(mysql.Open(config.MySql), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("数据库连接成功")
}
