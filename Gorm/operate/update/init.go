package update

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

var logWriter io.Writer
var DB *gorm.DB

func init() {
	// 定义DSN
	// 1. 定义DSN
	const dsn = "root:Sakura@tcp(127.0.0.1:3306)/gormstudy?charset=utf8mb4&parseTime=true&loc=Local"
	//初始化LoggerWriter
	logWriter, _ = os.OpenFile("./sql.log", os.O_CREATE|os.O_APPEND, 0644)
	customLogger := logger.New(log.New(logWriter, "\r\n", log.LstdFlags), logger.Config{
		//慢查询阈值
		SlowThreshold: 200 * time.Millisecond,
		//日志级别
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		//不彩色化
		Colorful: false,
	})

	// 2. 连接服务器(池)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		fmt.Println("连接失败")
		log.Fatalln(err)
	}
	DB = db
}
