package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// main 执行前，执行数据库连接，读取数据库环境变量
func init() {
	//mode := os.Getenv("GIN_MODE")
	//if mode == "" {
	//	mode = gin.DebugMode
	//}
	//gin.SetMode(mode)
	//repository.ConnectDB()
}

var (
	DB  *gorm.DB
	err error
)

func main() {
	//r := router.Router()
	//gin.DefaultWriter = os.Stdout
	//err := r.Run(":8081")
	//if err != nil {
	//	return
	//}
	//const Local_Mysql = "root:ly15984093508@tcp(124.222.140.95:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//DB, err = gorm.Open(mysql.Open(Local_Mysql), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("数据库连接失败", err)
	//	return
	//}
	//fmt.Println("数据库连接成功")

	// 连接数据库
	db, err := sql.Open("mysql", "root:ly15984093508@tcp(website-mysql:3306)/blog")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	fmt.Println("Successfully connected to MySQL!")

	r := gin.Default() // 初始化gin实例
	r.Run(":8081")
}
