package Transaction

import (
	"gorm.io/gorm"
)

// 作者模型 Author
type Author struct {
	gorm.Model
	Name string

	//积分
	Points int
}
