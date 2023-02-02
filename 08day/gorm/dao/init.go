package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// gorm 关系映射
// gorm和模型的关系映射
// 为了数据的一致性，gorm默认会在事物里执行写入操作
// 一下设置可以获得性能提升
//
//	db, err := gorm.Open(mysql.Open("gorm.db"), &gorm.Config{
//	  SkipDefaultTransaction: true,
//	})

func Init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connect database failed,err:", err)
		return
	}
	log.Println("connect db success")
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
