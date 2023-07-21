package query

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type ContentStrPK struct {
	ID          string `gorm:"primaryKey"`
	Subject     string
	Likes       uint
	Views       uint
	PublishTime *time.Time
}

type Content struct {
	gorm.Model
	Subject     string
	Likes       uint
	Views       uint
	PublishTime *time.Time
}

func GetPrimaryKey() {
	DB.AutoMigrate(&Content{}, &ContentStrPK{})

	//一:查询单条
	// 1.1主键为数值类型
	c := &Content{}
	if err := DB.First(&c, 10).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}
	// 1.2主键为string类型
	CS := &ContentStrPK{}
	if err := DB.First(&CS, "id= ?", "one").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}
	//二:查询多条
	//2.1主键为数值类型
	var IdSlice []Content
	if err := DB.Find(&IdSlice, []uint{100, 200, 300}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}
	//2.2主键为字符串类型
	var IDSlice2 []ContentStrPK
	if err := DB.Find(&IDSlice2, "id IN ?", []string{"one", "two", "three"}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}
}
