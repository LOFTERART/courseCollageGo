package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

// 初始化数据库连接等mysql
func Initialized() {
	// 申明变量
	var (
		err error
	)
	// 从本地读取环境变量
	godotenv.Load()
	DSN := os.Getenv("MYSQL_DSN")
	// 连接数据库
	DB, err = gorm.Open("mysql", DSN)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
		return
	}
	DB.SingularTable(true)
	DB.AutoMigrate(&Course{}, &UserOrder{}, &User{}, &Address{}, &Book{}, &Admin{}, &Index{}, &Article{}, &Classify{}, &Banner{}, &Studio{})
	// debug 模式开启sql日志
	DB.LogMode(false)
}
