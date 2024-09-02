package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web_backend.com/m/v2/config"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	var dsn string
	if gin.Mode() == gin.DebugMode {
		dsn = config.Local_Mysql
	} else {
		dsn = config.Release_MySql
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
}
