package query

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func WhereMethod() {
	s := []Content{}

	// inline 条件内联条件
	//if err := DB.Find(&s, "like > ? AND subject like ?", 100, "Gorm%").Error; err != nil {
	//	log.Fatal(err)
	//}

	// Where,通常在多态拼凑条件时使用
	// Where返回DB对象,放到query保存
	query := DB.Where("likes > ?", 100)
	// query再次调用Where,形成复合条件

	Subject := "Sakura"
	// 当前用户输出位Subject不为空时候,才拼凑字符串
	if Subject != "" {
		query.Or("subject like ?", Subject+"%")
	}
	// 用query调用find
	if err := query.Find(&s).Error; err != nil {
		log.Fatal(err)
	}
}

func WhereType() {
	s := []Content{}

	//(1 or 2 and (3 and (4 or 5))
	//(1 or 2)
	//conA := DB.Where("likes > ?", 100).Or("likes = 200", "200")
	//
	////(3 and ( 4 or 5 ))
	//conB := DB.Where("view > ?", 3).Where(DB.Where("views <= ?").Or("Subject like ?", "Sakura"))
	//
	//query := DB.Where(conA).Where(conB)
	//
	//if err := query.Find(&s).Error; err != nil {
	//	log.Fatal(err)
	//}

	// Map 构建条件 and, = in
	//query := DB.Where(map[string]any{
	//	"id":   100,
	//	"view": 200,
	//})
	//if err := query.Find(&s).Error; err != nil {
	//	log.Fatal(err)
	//}

	//struct
	query := DB.Where(Content{
		Views:   200,
		Subject: "Sakura",
	})
	if err := query.Find(&s).Error; err != nil {
		log.Fatal(err)
	}

}

func PlaceHolder() {
	s := []Content{}

	// 匿名
	//query1 := DB.Where("likes = ? AND views IN", 100, 1000)

	// 具名
	//query2 := DB.Where("likes = @like AND subject like @subject", sql.Named("like", 100), sql.Named("subject", "Sakura")) // Grom还支持使用map的形式具名
	query3 := DB.Where("likes = @like AND subject like @subject", map[string]any{
		"subject": "gorm%",
		"like":    100,
	})

	if err := query3.Find(&s).Error; err != nil {
		log.Fatal(err)
	}

}

func OrderBy() {
	var cs []Content

	ids := []uint{2, 3, 1}
	query := DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "field(id, ?)",
			Vars:               []any{ids},
			WithoutParentheses: true,
		},
	})
	if err := query.Find(&cs, ids).Error; err != nil {
		log.Fatalln(err)
	}
	for _, c := range cs {
		fmt.Println(c.ID)
	}
}

// 分页数据结构
type Page struct {
	Page     int
	PageSize int
}

// 默认值
const (
	Default         = 1
	DefaultPageSize = 12
)

// 翻页程序
func PageOperation(pager Page) {
	// 确定 offset 和 pagesize
	page := Default
	if pager.Page != 0 {
		page = pager.Page
	}
	pagesize := DefaultPageSize
	if pager.PageSize != 0 {
		pagesize = pager.PageSize
	}
	//计算offset
	//page ,pagesize,offset
	//1 10 0
	//2 10 0
	offset := pagesize * (page - 1)

	var content []Content
	if err := DB.Offset(offset).Limit(pagesize).Find(&content); err != nil {
		log.Fatalln(err)
	}
}

// 定义一个函数 , 用于得到
func Paginate(pager Page) func(*gorm.DB) *gorm.DB {
	// 确定 offset 和 pagesize
	page := Default
	if pager.Page != 0 {
		page = pager.Page
	}
	pagesize := DefaultPageSize
	if pager.PageSize != 0 {
		pagesize = pager.PageSize
	}

	//计算offset
	//page ,pagesize,offset
	//1 10 0
	//2 10 0
	offset := pagesize * (page - 1)

	return func(db *gorm.DB) *gorm.DB {
		// 使用闭包的变量
		return db.Offset(offset).Limit(pagesize)
	}
}

func PageScope(pager Page) {
	var content []Content
	if err := DB.Scopes(Paginate(pager)).Find(&content); err != nil {
		log.Fatalln(err)
	}
}

func GroupHaving() {
	DB.AutoMigrate(&Content{})

	type Result struct {
		UserID     uint
		TotalLikes int
		TotalViews int
		AvgViews   int
	}

	var rs []Result
	if err := DB.Select("author_id", "SUM(likes) as total_likes", "SUM(views) as total_views", "AVG(views) as avg_views").
		Group("author_id").Having("total_views > ?", 99).
		Find(&rs).Error; err != nil {
		log.Fatalln(err)
	}
	// SELECT `author_id`,SUM(likes) as total_likes,SUM(views) as total_views,AVG(views) as avg_views FROM `msb_content` WHERE `msb_content`.`deleted_at` IS NULL GROUP BY `author_id` HAVING total_views > 99

}

func Locking() {
	var cs []Content
	if err := DB.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&cs, "likes > ?", 10).Error; err != nil {
		log.Fatalln(err)
	}
	//[4.904ms] [rows:19] SELECT * FROM `msb_content` WHERE likes > 10 AND `msb_content`.`deleted_at` IS NULL FOR UPDATE

	if err := DB.Clauses(clause.Locking{Strength: "SHARE"}).Find(&cs, "likes > ?", 10).Error; err != nil {
		log.Fatalln(err)
	}
	// [2.663ms] [rows:19] SELECT * FROM `msb_content` WHERE likes > 10 AND `msb_content`.`deleted_at` IS NULL FOR SHARE
}

// Author模型
type Author struct {
	gorm.Model
	Status int

	Name  string
	Email string
}

func SubQuery() {
	DB.AutoMigrate(&Author{}, &Content{})

	//条件子查询
	//select id form zuthor where status=0
	//authorIDs := DB.Model(&Author{}).Select("id").Where("status=?", 0)
	//var cs []Content
	//if err := DB.
	//	Where("author_id IN (?)", authorIDs).
	//	Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}

	//From型子查询
	type Result struct {
		Subject string
		Likes   int
	}
	var rs []Result
	//select subject, likes from content where publish_time is null
	fromQuery := DB.Model(&Content{}).Select("subject", "likes").Where("publish_time is null")
	if err := DB.Table("(?) as temp", fromQuery).
		Where("likes > ?", 10).
		Find(&rs).Error; err != nil {
		log.Fatalln(err)
	}
}

func (c *Content) AfterFind(db *gorm.DB) (err error) {
	if c.ID == 13 {
		fmt.Println("触发查询钩子函数")
	}
	return nil
}

func QueryHook() {
	var c Content
	if err := DB.First(&c, 13).Error; err != nil {
		logrus.Error(err)
	}
}
