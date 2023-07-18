package create

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type NewContent struct {
	gorm.Model

	Subject     string
	Likes       uint
	Views       uint
	PublishTime *time.Time
}

func NewContentDefault() NewContent {
	return NewContent{
		Likes: 114514,
		Views: 114515,
	}
}

func DefaultValueOften() {
	DB.AutoMigrate(&NewContent{})

	data := NewContentDefault()
	data.Subject = "原始内容"
	DB.Create(&data)
	fmt.Println(data.Likes, data.Views)
}

func Default() {
	DB.AutoMigrate(&NewContent{})

	data := NewContent{Subject: "默认值处理", Likes: 0, Views: 0}

	DB.Create(&data)
	fmt.Println(data)

}
func SelectOmit() {
	DB.AutoMigrate(&NewContent{})

	data := NewContent{}
	data.Subject = "插入特定字段"
	data.Likes = 0
	data.Views = 99
	now := time.Now()
	data.PublishTime = &now

	DB.Omit("Subject", "Likes", "UpdatedAt").Create(&data)
}
