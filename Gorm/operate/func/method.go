package _func

import (
	"Go_WorkSpace/Gorm/operate/create"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model

	Username string
	Name     string
	Email    string
	Birthday *time.Time
}

func OperatorType() {
	create.DB.AutoMigrate(&User{})

	var users []User

	//一步操作
	err := create.DB.Where("birthday IS NOT NULL").
		Where("email like ?", "@163.com%").
		Order("name DESC").
		Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

	//分步操作:返回的都是存储好的DB对象,所以就可以存储起来
	//query := DB.Where("birthday IS NOT NULL")
	//query = query.Where("email like ?", "@163.com%")
	//query = query.Order("name DESC")
	//query.Find(&users)
}
