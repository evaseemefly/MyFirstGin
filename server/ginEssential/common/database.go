package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	//初始化数据库，并返回 db
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessenital"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect databse,err:" + err.Error())
	}
	// 自动创建数据库
	db.AutoMigrate(&User{})

	return db
}