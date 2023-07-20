package create

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (c *NewContent) BeforeCreate(db *gorm.DB) error {
	// 业务
	if c.PublishTime == nil {
		now := time.Now()
		c.PublishTime = &now
	}

	// 配置
	db.Statement.AddClause(clause.OnConflict{UpdateAll: true})

	return nil
}

// 出现一个错误
func (c *NewContent) AfterCreate(db *gorm.DB) error {
	return errors.New("custom error")
}

func Hook() {
	DB.AutoMigrate(&NewContent{})

	content := NewContent{Likes: 10, Subject: "HookTest"}

	DB.Create(&content)
}
