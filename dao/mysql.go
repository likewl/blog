package dao

import (
	"blog/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := "like:Qq123_456@(localhost)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Print("数据库连接失败")
		return
	}
	//自动迁移
	DB.AutoMigrate(&model.Blog{})
	DB.AutoMigrate(&model.Tag{})
	DB.AutoMigrate(&model.Comment{})
	DB.AutoMigrate(&model.Type{})
	DB.AutoMigrate(&model.User{})
	////创建用户
	//user := model.User{
	//	BlogName: "like‘s blog",
	//	Name: "like",
	//	UserName: "admin",
	//	Password: "admin",
	//	UserEmail: "",
	//	UserType: 1,
	//	UserImg: "/static/images/header.jpg",
	//	UserCreate: middleware.GetTime(),
	//	UserUpdate: middleware.GetTime(),
	//}
	//DB.Create(&user)

	//创建表

	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
