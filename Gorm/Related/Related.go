package Related

//
//import (
//	"Go_WorkSpace/Gorm/Transaction"
//	"fmt"
//	"log"
//)
//
//func Related() {
//	// 利用migrate构件表
//	//包含一对多,一对一,多对多关系
//	if err := Transaction.DB.AutoMigrate(&Transaction.Author{}); err != nil {
//		log.Fatalln(err)
//	}
//	log.Println("构建成功")
//}
//
//func AssocAppend() {
//	//一对多的关系, Author ,Essay
//	//创建测试数据
//	author := Transaction.Author{Name: "Sakura"}
//	if err := Transaction.DB.Create(&author).Error; err != nil {
//		log.Fatalln(err)
//	}
//	essay := Transaction.Essay{Subject: "第一篇内容"}
//	essay2 := Transaction.Essay{Subject: "第二篇内容"}
//	if err := Transaction.DB.Create([]*Transaction.Essay{&essay, &essay2}).Error; err != nil {
//		log.Fatalln(err)
//	}
//
//	//添加关联
//	if err := Transaction.DB.Model(&author).Association("Essay").Append([]Transaction.Essay{essay, essay2}); err != nil {
//		log.Println(err)
//	}
//
//	//B: 多对多 Essay M:N Tag
//	tag1 := Transaction.Tag{Title: "Golang"}
//	tag2 := Transaction.Tag{Title: "Java"}
//	tag3 := Transaction.Tag{Title: "Python"}
//	//让第一篇文章关联其中连个,第二篇文章关联三个
//	if err := Transaction.DB.Create([]*Transaction.Tag{&tag1, &tag2, &tag3}).Error; err != nil {
//		log.Fatalln(err)
//	}
//	//第一篇文章关联 tag1 tag3
//	//第二篇文章关联 tag1 tag2 tag3
//	if err := Transaction.DB.Model(&essay).Association("Tags").Append([]Transaction.Tag{tag1, tag3}); err != nil {
//		log.Fatalln(err)
//	}
//	if err := Transaction.DB.Model(&essay2).Association("Tags").Append([]Transaction.Tag{tag1, tag2, tag3}); err != nil {
//		log.Fatalln(err)
//	}
//
//	//查看关联表
//	//log.Println()
//}
//
//func BelongTo() {
//	author := Transaction.Author{Name: "LF"}
//	if err := Transaction.DB.Create(&author).Error; err != nil {
//		log.Fatalln(err)
//	}
//
//	essay1 := Transaction.Essay{Subject: "Gorm关联操作"}
//	essay2 := Transaction.Essay{Subject: "Golang并发编程"}
//	if err := Transaction.DB.Create([]*Transaction.Essay{&essay1, &essay2}).Error; err != nil {
//		log.Fatalln(err)
//	}
//
//	//添加关联
//	if err := Transaction.DB.Model(&essay1).Association("Author").Append(&author); err != nil {
//		log.Fatalln(err)
//	}
//	if err := Transaction.DB.Model(&essay2).Association("Author").Append(&author); err != nil {
//		log.Fatalln(err)
//	}
//
//	////最后查看Author的文章
//	//log.Fatalln(author.Essay)
//	log.Println(essay1)
//	log.Println(essay2)
//}
//
//func ReplaceRelated() {
//	// 替换关联
//	author := Transaction.Author{Name: "Sakura"}
//	if err := Transaction.DB.Create(&author).Error; err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(author.ID)
//
//	essay1 := Transaction.Essay{Subject: "第一篇文章"}
//	essay2 := Transaction.Essay{Subject: "第二篇文章"}
//	essay3 := Transaction.Essay{Subject: "第三篇文章"}
//	if err := Transaction.DB.Create([]*Transaction.Essay{&essay1, &essay2, &essay3}).Error; err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(essay1.ID, essay2.ID, essay3.ID)
//
//	//1.首先建立关联
//	if err := Transaction.DB.Model(&author).Association("Essay").Replace([]Transaction.Essay{essay1, essay2, essay3}); err != nil {
//		log.Fatalln(err)
//	}
//	//2.删除关联
//	if err := Transaction.DB.Model(&author).Association("Essay").Clear(); err != nil {
//		log.Fatalln(err)
//	}
//
//}
//
//func RelatedQuery() {
//	essay := Transaction.Essay{}
//	Transaction.DB.First(&essay, 14)
//
//	//查询关联的tags
//	var Tags []Transaction.Tag
//	if err := Transaction.DB.Model(&essay).Association("Tags").Find(&Tags); err != nil {
//		log.Fatalln(err)
//	}
//	//fmt.Println(essay.Tags)
//	fmt.Println(Tags)
//}
//
//func RelatedSave() {
//	var t1 Transaction.Tag
//	Transaction.DB.First(&t1, 4) // id为4的tag
//
//	essay := Transaction.Essay{
//		Subject: "自动存储",
//		Author:  Transaction.Author{Name: "Sakura"},
//		Tags: []Transaction.Tag{
//			t1,
//			{Title: "Go"},
//			{Title: "GORM"},
//		},
//	}
//	// Save可以根据传入的模型是否有主键,来决定insert或者update
//	if err := Transaction.DB.Save(&essay).Error; err != nil {
//		log.Fatalln(err)
//	}
//	log.Println(essay)
//
//}
//
//func Preload() {
//	// A.直接一步查询Author对应的Essays
//	author := Transaction.Author{}
//	if err := Transaction.DB.Preload("Essay").First(&author, 14).Error; err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(author.Essay) //输出流author的两排你文章"Gorm关联操作和"Golang并发编程"
//
//	fmt.Println("------------------------------------------")
//
//	// B.加上条件过滤,过滤id在28的文章
//	if err := Transaction.DB.
//		Preload("Essay", "id IN ?", []uint{28}).
//		First(&author, 14).Error; err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(author.Essay)
//
//	fmt.Println("------------------------------------")
//
//	// C.支持多次调用,同时预加载多个关联
//	e := Transaction.Essay{}
//	if err := Transaction.DB.
//		Preload("Author").
//		Preload("EssayMate").
//		Preload("Tags").
//		First(&e, 14).Error; err != nil {
//		log.Fatalln(err)
//	}
//	log.Println(e)
//}
//
//func LevelPreload() {
//	author := Transaction.Author{}
//	if err := Transaction.DB.
//		//Preload("Essays").
//		// 多级关联
//		Preload("Essay.Tags").
//		First(&author, 7).Error; err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(author.Essay[0].Tags)
//	fmt.Println(author.Essay[1].Tags)
//}
//
//func InsertValue() {
//	author := Transaction.Author{Name: "Sakura"}
//	if err := Transaction.DB.Create(&author).Error; err != nil {
//		log.Fatalln(err)
//	}
//	essay := Transaction.Essay{Subject: "第一篇内容"}
//	essay2 := Transaction.Essay{Subject: "第二篇内容"}
//	if err := Transaction.DB.Create([]*Transaction.Essay{&essay, &essay2}).Error; err != nil {
//		log.Fatalln(err)
//	}
//
//	//添加关联
//	if err := Transaction.DB.Model(&author).Association("Essay").Append([]Transaction.Essay{essay, essay2}); err != nil {
//		log.Println(err)
//	}
//
//	//删除author
//	if err := Transaction.DB.Unscoped().Delete(&author).Error; err != nil {
//		log.Fatalln(err)
//	}
//}
