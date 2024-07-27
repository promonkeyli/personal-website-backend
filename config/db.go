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

const Local_Mysql = "root:ly15984093508@tcp(promonkeyli.top:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
