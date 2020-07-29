package dao

import (
	"blog/middleware"
	"blog/model"
)

func CreatTag(name string) {
	var tags model.Tag
	tags.TagName = name
	DB.Create(&tags)
}
func UpdateTag(id int, name string) {
	var tag model.Tag
	DB.Debug().Model(&tag).Where("id = ?", id).Update("tag_name", name)

}

func DeleteTag(id int) {
	var tag model.Tag
	DB.Delete(&tag, id)
}

func SelectTag() *[]model.Tag {
	var tag []model.Tag
	DB.Debug().Find(&tag)
	return &tag
}
func CountTag() int {
	var count int
	DB.Debug().Table("tags").Count(&count)
	return count
}

//分页查询
func TagPage(pageIndex int) (*[]model.Tag, int) {
	var tag []model.Tag
	var count int
	DB.Offset((pageIndex - 1) * 10).Limit(10).Find(&tag)
	count = len(tag)

	return &tag, count
}

//删除标签和博文关联表 以博文为主键
func DelTagBlog(id int) {
	var TagBlog middleware.TagBlog
	DB.Debug().Where("blog_id=?", id).Delete(&TagBlog)
}

//删除标签和博文关联表 以标签为主键
func DelBlogTag(id int) {
	var TagBlog middleware.TagBlog
	DB.Debug().Where("tag_id=?", id).Delete(&TagBlog)
}
func RandTag() *[]model.Tag {
	var rand []model.Tag
	DB.Debug().Raw("SELECT * FROM `tags` ORDER BY RAND() LIMIT 10").Scan(&rand)
	return &rand
}

//查询每个标签的博文数量
func SelectBlogTag(id int) int {
	var count int
	DB.Debug().Table("tag_blogs").Where("tag_id=?", id).Count(&count)
	return count
}

//查询每个标签de id
func SelectTagId(name string) *model.Tag {
	var id model.Tag
	DB.Debug().Table("tags").Where("tag_name=?", name).First(&id)
	return &id
}
func SelectTagNameAll(name string, pageIndex int) (*[]middleware.TagAndBlog, int) {
	all1 := SelectTagId(name)
	id := all1.Id
	var Blog []middleware.TagAndBlog
	//var Blog model.Blog
	DB.Debug().Raw("SELECT b.id,b.title,b.type_id,t.tag_name,b.description,b.recommend,b.numb,b.update_time,b.image,a.`type_name`  FROM `blogs` b LEFT JOIN `tag_blogs` tb ON b.id = tb.`blog_id` LEFT JOIN tags t ON tb.`tag_id`=t.`id` LEFT JOIN `types` a ON b.type_id = a.`id` WHERE t.`id`=?", id).Order("b.id  DESC").Offset((pageIndex - 1) * 10).Limit(10).Scan(&Blog)
	count := len(Blog)
	return &Blog, count
}

func SelectTagNameBlog(id int) *[]model.Tag {
	var Blog []model.Tag
	//var Blog model.Blog
	DB.Debug().Raw("SELECT t.id,t.tag_name FROM `tags` t LEFT JOIN `tag_blogs` tb ON t.id = tb.`tag_id` LEFT JOIN `blogs` b ON tb.`blog_id`=b.`id`  WHERE b.`id`=?", id).Scan(&Blog)
	return &Blog
}
