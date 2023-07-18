package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// 1. 定义DSN
	dsn := "root:Sakura@tcp(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=true&loc=Local"

	// 2. 连接服务器(池)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败")
		log.Fatalln(err)
	}
	fmt.Println(db)
}
