package config

import (
	"fmt"
	"os"
)

var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbUser = os.Getenv("DB_USER")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbName = os.Getenv("DB_NAME")

var Release_MySql = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

// const Local_Mysql = "root:ly@15984093508@tcp(127.0.0.1:3306)/?charset=utf8mb4"
const Local_Mysql = "root:ly@15984093508@tcp(website-app:3306)/?charset=utf8mb4"
