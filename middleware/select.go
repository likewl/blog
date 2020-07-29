package middleware

import (
	"blog/model"
)

type BlogType struct {
	Id         int
	Title      string `form:"title" json:"title"`
	TypeId     int    `form:"typeId" json:"typeId"`
	TypeName   string `form:"typeName" json:"typeName"`
	Recommend  string `form:"recommend" json:"recommend"`
	UpdateTime string `form:"update_Time" json:"update_Time"`
}

type CreatBlog struct {
	Blog model.Blog
	Msg  string `form:"msg" json:"msg"`
	Page int    `form:"page" json:"page"`
}

type TagBlog struct {
	TagId  int `form:"tagId" json:"tagId"`
	BlogId int `form:"blogId" json:"blogId"`
}

type BlogView struct {
	Id          int
	Title       string `form:"title" json:"title"`
	TypeId      int    `form:"typeId" json:"typeId"`
	TypeName    string `form:"typeName" json:"typeName"`
	Description string `gorm:"type:text" form:"description" json:"description"`
	Recommend   string `form:"recommend" json:"recommend"`
	Numb        int    `form:"numb" json:"numb"`
	UpdateTime  string `form:"update_Time" json:"update_Time"`
	Image       string `gorm:"type:text" form:"image" json:"image"` //首图
}
type TagAndBlog struct {
	Id          int
	Title       string `form:"title" json:"title"`
	TagId       int    `form:"tagId" json:"tagId"`
	TagName     string `form:"tagName" json:"tagName"`
	Description string `gorm:"type:text" form:"description" json:"description"`
	Recommend   string `form:"recommend" json:"recommend"`
	Numb        int    `form:"numb" json:"numb"`
	UpdateTime  string `form:"update_Time" json:"update_Time"`
	Image       string `gorm:"type:text" form:"image" json:"image"` //首图
	TypeName    string `form:"typeName" json:"typeName"`
}

//年份结构体
