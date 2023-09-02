package Transaction

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

func Transaction() {
	//构建表
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}

	//初始化测试数据
	author1 := Author{Name: "Sakura", Points: 2000}
	author2 := Author{Name: "LF", Points: 1000}
	if err := DB.Create([]*Author{&author1, &author2}).Error; err != nil {
		log.Fatalln(err)
	}

	//事务操作
	//开启事务
	tx := DB.Begin()
	// 有时候需要数据库是否支持事务
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	//执行修改
	author1.Points += 2500
	author2.Points -= 2500

	//1.执行错误
	if err := tx.Save(&author1).Error; err != nil {
		tx.Rollback()
	}
	if err := tx.Save(&author2).Error; err != nil {
		tx.Rollback()
	}

	//2.业务逻辑导致的错误:要求author的积分不能为负数
	if author1.Points < 0 || author2.Points < 0 {
		log.Fatalln("事务已经被回滚")
		tx.Rollback()
	}

	//提交事务
	tx.Commit()

	//集中处理错误
	//根据是否有错误决定是否回滚
	//if err1 != nil || err2 != nil {
	//	tx.Rollback()
	//} else {
	//	tx.Commit()
	//}
}

func Callback() {
	//构建表
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}

	//初始化测试数据
	author1 := Author{Name: "Sakura", Points: 2000}
	author2 := Author{Name: "LF", Points: 1000}
	if err := DB.Create([]*Author{&author1, &author2}).Error; err != nil {
		log.Fatalln(err)
	}

	//实现事务
	if err := DB.Transaction(func(tx *gorm.DB) error {
		//执行修改
		author1.Points += 2500
		author2.Points -= 2500

		//执行sql,可能导致的错误
		if err := tx.Save(&author1).Error; err != nil {
			//直接return err
			return err
		}
		if err := tx.Save(&author2).Error; err != nil {
			return err
		}

		//2.业务逻辑导致的错误:要求author的积分不能为负数
		if author1.Points < 0 || author2.Points < 0 {
			return errors.New("author1.Points < 0 || author2.Points < 0")
		}

		// nil 的返回,会导致事务提交
		return nil
	}); err != nil {
		//返回错误，为了后续的业务逻辑处理
		//为了通知我们，事务成功还是失败
		//返回错误，不影响事务的提交和回滚
		log.Fatalln(err)
	}
}

func SavePoint() {
	//构建表
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}

	//初始化测试数据
	author1 := Author{Name: "Sakura", Points: 2000}
	author2 := Author{Name: "LF", Points: 1000}
	if err := DB.Create([]*Author{&author1, &author2}).Error; err != nil {
		log.Fatalln(err)
	}

	//事务操作
	//开启事务
	tx := DB.Begin()
	// 有时候需要数据库是否支持事务
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	//第一次转钱成功
	author1.Points += 200
	if err := tx.Save(&author1).Error; err != nil {
		tx.Rollback()
	}

	author2.Points -= 200
	if err := tx.Save(&author2).Error; err != nil {
		tx.Rollback()
	}

	DB.SavePoint("FirstGet")
	//第二次转钱失败
	author1.Points += 2000
	if err := tx.Save(&author1).Error; err != nil {
		tx.Rollback()
	}
	author2.Points -= 2000
	if err := tx.Save(&author2).Error; err != nil {
		tx.Rollback()
	}

	//2.业务逻辑导致的错误:要求author的积分不能为负数
	if author1.Points < 0 || author2.Points < 0 {
		tx.RollbackTo("FirstGet")
	}

	//提交事务
	tx.Commit()
}
