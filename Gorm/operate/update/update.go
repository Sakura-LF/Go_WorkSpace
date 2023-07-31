package update

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

//type UpContent struct {
//	gorm.Model
//	Subject     string
//	Likes       uint
//	Views       uint
//	PublishTime *time.Time
//}

type Content struct {
	gorm.Model
	Subject     string
	Likes       uint
	Views       uint
	PublishTime *time.Time
}

func PkUpdate() {
	c := Content{}

	//无主键ID:插入
	if err := DB.Save(&c).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)

	//有主键ID:更新
	if err := DB.Save(&c).Error; err != nil {
		log.Fatalln(err)
	}
}

func UpdateWhere() {
	// 更新的字段值数据
	values := map[string]any{
		"subject": "where Update Row",
		"likes":   114514,
	}

	//执行带有条件的更新
	result := DB.Model(&Content{}).Where("likes > ?", 100).Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	// 获取更新结果,更新的记录数量,(受影响的记录数)
	log.Println("updated rows num:", result.RowsAffected)
}

func UpdateNoWhere() {
	// 更新的字段值数据
	values := map[string]any{
		"subject": "where Update Row",
		"likes":   114514,
	}

	//执行带有条件的更新
	result := DB.Model(&Content{}).
		//Where("likes > ?", 100).
		Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// 获取更新结果,更新的记录数量,(受影响的记录数)
	log.Println("updated rows num:", result.RowsAffected)
}

func UpdateExpr() {
	// 更新的字段值数据
	values := map[string]any{
		"subject": "where Update Row",
		"likes":   gorm.Expr("likes + ?", 10),
	}

	//执行带有条件的更新
	result := DB.Model(&Content{}).
		Where("likes > ?", 100).
		Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// 获取更新结果,更新的记录数量,(受影响的记录数)
	log.Println("updated rows num:", result.RowsAffected)
}

func Deletewhere() {
	//删除
	//不需要model,因为delete()要传入一个模型
	result := DB.Delete(&Content{}, "likes < ?", 100)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	//条件删除
	result2 := DB.Where("likes < ?", 100).Delete(&Content{})
	if result2.Error != nil {
		log.Fatalln(result2.Error)
	}
}
