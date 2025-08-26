package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DatabaseName string = "LingXiaoYun"

func ConnectDatabase() {
	dsn := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("数据库连接失败")
		panic(err)
	}

	DB = db
	log.Println("数据库连接成功")
}
