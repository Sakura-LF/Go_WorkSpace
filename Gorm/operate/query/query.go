package query

import (
	"database/sql"
	"errors"
	"fmt"
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

	// 不需要迁移
	// 禁用写操作
	SV string `gorm:"-migration;<-:false>"`
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

func QueryOne() {
	c := Content{}
	if err := DB.First(&c, "id > ?", 42).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}

	o := Content{}
	if err := DB.Last(&o, "id > ?", 42).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}

	n := Content{}
	if err := DB.Take(&n, "id > ?", 42).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}

	f := Content{}
	if err := DB.Limit(1).Find(&f, "id > ?", 42).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}

	fs := Content{}
	if err := DB.Find(&fs, "id > ?", 42).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Foud")
		} else {
			log.Fatal(err)
		}
	}
}

func QueryToMap() {
	//查询单条
	One := map[string]any{}
	if err := DB.Model(&Content{}).First(&One, 13).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Found")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(One["id"], One["id"] == 13)
	//value是any类型,使用时需要类型断言
	if One["id"].(uint) == 13 {
		fmt.Println("Test Success")
	}
	fmt.Println(One)

	//查询多条
	Many := []map[string]any{}
	if err := DB.Model(&Content{}).Find(&Many, []uint{13, 14, 15}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record Not Found")
		} else {
			log.Fatal(err)
		}
	}
	for _, value := range Many {
		fmt.Println(value)
	}
}

func QueryPluck() {
	//使用切片存储
	Subjects := []sql.NullString{}
	if err := DB.Model(&Content{}).Pluck("Subject", &Subjects).Error; err != nil {
		log.Fatal(err)
	}
	for _, Subjects := range Subjects {
		//NullString的使用
		//Subject.Valid是布尔类型,表示是否为空字符串
		//true:表示不为空
		//false:表示为空
		if Subjects.Valid {
			fmt.Println(Subjects.String)
		} else {
			fmt.Println("NULL")
		}
	}
}

func QueryPluckEXP() {
	//使用切片存储
	Subjects := []sql.NullString{}
	//Pluck中的column也可以用表达式
	if err := DB.Model(&Content{}).Pluck("concat(subject,'-',likes)", &Subjects).Error; err != nil {
		log.Fatal(err)
	}
	for _, Subjects := range Subjects {
		//NullString的使用
		//Subject.Valid是布尔类型,表示是否为空字符串
		//true:表示不为空
		//false:表示为空
		if Subjects.Valid {
			fmt.Println(Subjects.String)
		} else {
			fmt.Println("NULL")
		}
	}
}

func QuerySelect() {
	c := Content{}
	if err := DB.Select("subject", "likes").First(&c, "14").Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", c)
}

func QuerySelectEXP() {
	c := Content{}
	if err := DB.Select("subject", "likes", "concat(subject,'-',likes) AS sv").First(&c, "14").Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", c)
}

func QueryDistinct() {
	c := Content{}

	if err := DB.Distinct("*").First(&c, 14).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%-v", c)
}
