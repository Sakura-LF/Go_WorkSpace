package update

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

func Sql() {
	type Result struct {
		Id      uint
		Subject string
		Likes   int
	}
	var rs []Result

	// sql语句
	sql := "SELECT `id`,`subject`,`likes` FROM `contents` where `likes` > ? "

	// 执行sql,并扫描结果
	if err := DB.Raw(sql, 1000).Scan(&rs).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rs)
}

func SqlExc() {
	//SQL
	sql := "INSERT INTO contents(id,subject,likes) values (?,?,?)"

	//执行sql语句,Exec()
	if err := DB.Exec(sql, 50, "执行类sql语句", 125).Error; err != nil {
		log.Fatalln(err)
	}
}

func Rows() {
	// sql语句
	sql := "SELECT `id`,`subject`,`likes` FROM `contents` where `likes` > ? "

	// 执行sql,并扫描结果
	rows, err := DB.Raw(sql, 1000).Rows()
	if err != nil {
		log.Fatalln(err)
	}
	//遍历 Rows
	for rows.Next() {
		//扫描列独立的变量
		var id uint
		var subject string
		var likes int
		rows.Scan(&id, &subject, &likes)
		fmt.Println(id, subject, likes)

		fmt.Println("----------------------------")
		//扫描到结构体
		type Result struct {
			Id      uint
			Subject string
			Likes   int
		}
		var rs Result
		if err := DB.ScanRows(rows, &rs); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(rs)
	}
}

func SessionIssue() {
	db := DB.Model(&Content{}).Where("views > ?", 100)
	var cs1 []Content
	db.Where("likes > ?", 99).Find(&cs1)

	// 第二条sql的view条件和之前一样,所以可以复用
	var cs2 []Content
	db.Where("likes > ?", 199).Find(&cs2)
}

func SessionTest() {
	// 需要重复使用的条件
	// 将Session方法前的配置,记录到了当前的会话中
	// 后边再次调用db的方法直到终结方法,会保持会话中的字句选项
	// 执行完终结方法后,再次调用db的方法到终结
	db := DB.Model(&Content{}).Where("views > ?", 100).
		Session(&gorm.Session{})

	var cs1 []Content
	db.Where("likes < ?", 5).Find(&cs1)

	// 第二条sql的view条件和之前一样,所以可以复用
	var cs2 []Content
	db.Where("likes > ?", 100).Find(&cs2)
}

func SessionOption() {
	db := DB.Model(&Content{}).Session(&gorm.Session{
		SkipHooks: true,
	})
	db.Save(&Content{Subject: "no hooks"})
}

func (c *Content) BeforeCreate(db *gorm.DB) error {
	log.Println("content before create hook")
	return nil
}

func DryRun() {
	db := DB.Model(&Content{}).Session(&gorm.Session{
		// DryRun
		DryRun: true,
	})
	//执行之后需要调用Statement属性获取sql语句
	stmt := db.Save(&Content{Subject: "no hooks"}).Statement
	log.Println(stmt.SQL.String()) //sql语句
	log.Println(stmt.Vars)         //参数变量
}

func ContextTOCancel() {
	// 设置一个定时Cancel的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	// 传递Context执行
	var cs []Content
	if err := DB.WithContext(ctx).Limit(10).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cs)
}
