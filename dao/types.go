package dao

import (
	"blog/middleware"
	"blog/model"
	"fmt"
	"html/template"
)

func CreatType(name string) {
	var types model.Type
	types.TypeName = name
	DB.Create(&types)
}
func UpdateType(id int, name string) {
	var type1 model.Type
	DB.Debug().Model(&type1).Where("id = ?", id).Update("type_name", name)

}

func DeleteType(id int) {
	var type1 model.Type
	DB.Delete(&type1, id)
}

func SelectName() *[]model.Type {
	var type1 []model.Type
	DB.Debug().Find(&type1)
	return &type1
}
func CountType() int {
	var count int
	DB.Debug().Table("types").Count(&count)
	return count
}
func Rand() *[]model.Type {
	var rand []model.Type
	DB.Debug().Raw("SELECT * FROM `types` ORDER BY RAND() LIMIT 5").Scan(&rand)
	return &rand
}
func SelectNameAll1(name string) *model.Type {
	var type1 model.Type
	DB.Where("type_name = ?", name).First(&type1)
	return &type1
}
func SelectNameAll(name string, pageIndex int) (*[]middleware.BlogView, int, error) {
	all1 := SelectNameAll1(name)
	id := all1.Id
	fmt.Println(id)
	var Blog []middleware.BlogView
	var count int
	//var Blog model.Blog
	err := DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.description,blogs.recommend,blogs.numb,blogs.update_time,blogs.image").Where("type_id = ?", id).Joins("left join types on types.id = blogs.type_id").Order("blogs.id  DESC").Offset((pageIndex - 1) * 10).Limit(10).Find(&Blog).Error
	count = len(Blog)
	return &Blog, count, err
}
func SelectBlogsByTypeId(typeId int, pageIndex int) (*[]middleware.BlogView, int, error) {
	var Blog []middleware.BlogView
	err := DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.content as description,blogs.recommend,blogs.numb,blogs.update_time,blogs.image").Where("type_id = ?", typeId).Joins("left join types on types.id = blogs.type_id").Order("blogs.id  DESC").Offset((pageIndex - 1) * 10).Limit(10).Find(&Blog).Error
	count := len(Blog)
	return &Blog, count, err
}

func SelectOneType(id int) *model.Type {
	var type1 model.Type
	var blog model.Blog
	DB.Where("id = ?", id).First(&blog)
	DB.Where("id = ?", blog.TypeId).First(&type1)
	return &type1
}

func Safe (string2 string) template.HTML {
	return template.HTML(string2)
}

//根据博文id差分类名
func SelectTypeNameByBlog(id int) *model.Type {
	var type1 model.Type

	DB.Debug().Raw("SELECT t.id,t.type_name FROM `types` t LEFT JOIN `blogs` b ON t.id = b.`type_id` WHERE b.`id`=?", id).Scan(&type1)

	return &type1
}
