package model

type Tag struct {
	Id      int
	TagName string
	Blogs   []Blog `gorm:"many2many:tag_blogs;"`
}
