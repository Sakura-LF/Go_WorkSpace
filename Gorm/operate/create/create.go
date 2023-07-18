package create

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type Content struct {
	gorm.Model
	Subject     string
	Likes       uint
	PublishTime *time.Time
}

func CreateBasic() {
	DB.AutoMigrate(&Content{})

	//模型也是记录l
	c1 := Content{}
	c1.Subject = "Gorm的使用"

	//执行创建
	result := DB.Create(&c1)
	//处理错误
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	// 最新的ID和影响的行数
	fmt.Println("ID:", c1.ID, "\n", "影响的行数:", result.RowsAffected)

	// 通过Map指定数据
	values := map[string]any{
		"Subject":     "Map指定值",
		"PublishTime": time.Now(),
	}
	// create
	//需要Model()来缺点哪一个模型对应的表
	result2 := DB.Model(&Content{}).Create(values)

	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println(result2.RowsAffected)

}

func CreateMulti() {
	DB.AutoMigrate(&Content{})

	//定义模型的切片
	modle1 := []Content{
		{Subject: "文章一"},
		{Subject: "文章二"},
		{Subject: "文章三"},
		{Subject: "文章四"},
	}
	// create插入
	result := DB.Create(&modle1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	//打印出影响的行数
	fmt.Println("影响的行数", result.RowsAffected)
	//遍历切片,显示ID
	for _, m := range modle1 {
		fmt.Println("ID:", m.ID)
	}

	//map结构
	model2 := []map[string]any{
		{"Subject": "文章1"},
		{"Subject": "文章2"},
		{"Subject": "文章3"},
		{"Subject": "文章4"},
	}
	result2 := DB.Model(&Content{}).Create(&model2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println("影响的行数", result2.RowsAffected)
}

func CreateBatch() {
	DB.AutoMigrate(&Content{})

	//定义模型的切片
	modle1 := []Content{
		{Subject: "文章一"},
		{Subject: "文章二"},
		{Subject: "文章三"},
		{Subject: "文章四"},
	}
	// create插入
	result := DB.CreateInBatches(&modle1, 2)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	//打印出影响的行数
	fmt.Println("影响的行数", result.RowsAffected)
	//遍历切片,显示ID
	for _, m := range modle1 {
		fmt.Println("ID:", m.ID)
	}

	//map结构
	model2 := []map[string]any{
		{"Subject": "文章1"},
		{"Subject": "文章2"},
		{"Subject": "文章3"},
		{"Subject": "文章4"},
	}
	result2 := DB.Model(&Content{}).CreateInBatches(&model2, 2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println("影响的行数", result2.RowsAffected)
}

func UpSert() {
	DB.AutoMigrate(&Content{})

	// 常规插入,原始数据
	data := Content{Subject: "原始标题", Likes: 10}
	DB.Create(&data)
	fmt.Println(data)

	//// 主键冲突
	//data2 := Content{Subject: "新标题", Likes: 20}
	//data2.ID = data.ID
	//if err := Gorm.DB.Create(&data2).Error; err != nil {
	//	log.Fatal(err)
	//}

	// 冲突后,更新全部字段
	data2 := Content{Subject: "新标题", Likes: 20}
	data2.ID = data.ID
	if err := DB.
		//冲突后只更新部分字段,例:冲突后更新likes
		Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"likes"})}).
		Create(&data2).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(data2)

}
