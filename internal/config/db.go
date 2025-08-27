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
	// 先连接 MySQL 不指定数据库
	dsnRoot := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	dbRoot, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
	if err != nil {
		log.Fatal("无法连接 MySQL，确认用户名密码和端口")
	}

	// 自动创建数据库（如果不存在）
	createSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", DatabaseName)
	if err := dbRoot.Exec(createSQL).Error; err != nil {
		log.Fatal("创建数据库失败:", err)
	}

	// 连接指定数据库
	dsnDB := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DatabaseName)
	db, err := gorm.Open(mysql.Open(dsnDB), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	DB = db
	log.Println("数据库连接成功:", DatabaseName)
}
