package dao

import (
	"blog/middleware"
	"blog/model"
)

func CreatBlog(blog *model.Blog) {
	DB.Create(blog)
}

//保存
func UpdateBlog(update *model.Blog) {
	DB.Debug().Save(update)

}
func SelectOne(id int) (*model.Blog, error) {
	var blog model.Blog
	err := DB.Where("id = ?", id).First(&blog).Error
	return &blog, err
}

func DeleteBlog(id int) {
	var tag model.Blog
	DB.Delete(&tag, id)
}
func Last() *model.Blog {
	//var tag *model.Blog
	m := new(model.Blog)
	DB.Last(m)
	return m
}

func BlogType(pageIndex int) (*[]middleware.BlogType, int) {
	var BlogType []middleware.BlogType
	var count int
	//var Blog model.Blog
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Joins("left join types on types.id = blogs.type_id").Order("blogs.id DESC").Offset((pageIndex - 1) * 10).Limit(10).Find(&BlogType)
	count = len(BlogType)
	return &BlogType, count
}

func CreatTagBlog(tag int, blog int) {
	var TagBlog middleware.TagBlog
	TagBlog.TagId = tag
	TagBlog.BlogId = blog
	DB.Debug().Create(&TagBlog)
}

//模糊查找
//只标题
func ResearchOne(title string, pageIndex int) *[]middleware.BlogType {
	var findOne []middleware.BlogType

	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Where("title like ?", "%"+title+"%").Joins("left join types on types.id = blogs.type_id").Order("blogs.id").Offset((pageIndex - 1) * 10).Limit(10).Find(&findOne)
	return &findOne
}

//标题和分类
func ResearchTwo(title string, typeId int, pageIndex int) *[]middleware.BlogType {
	var findTwo []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Where("type_id = ? AND title like ? ", typeId, "%"+title+"%").Joins("left join types on types.id = blogs.type_id").Order("blogs.id").Offset((pageIndex - 1) * 10).Limit(10).Find(&findTwo)

	return &findTwo
}

//推荐
func ResearchThree(title string, pageIndex int) *[]middleware.BlogType {
	var findThree []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Where("title like ? AND recommend = '是'", "%"+title+"%").Joins("left join types on types.id = blogs.type_id").Order("blogs.id").Offset((pageIndex - 1) * 10).Limit(10).Find(&findThree)
	return &findThree
}

//都勾选
func ResearchFour(title string, typeId int, pageIndex int) *[]middleware.BlogType {
	var findFour []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Where("title like ? AND type_id = ? AND recommend = '是' ", "%"+title+"%", typeId).Joins("left join types on types.id = blogs.type_id").Order("blogs.id").Offset((pageIndex - 1) * 10).Limit(10).Find(&findFour)
	return &findFour
}

//只分类
func ResearchFive(typeId int, pageIndex int) *[]middleware.BlogType {
	var findTwo []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Where("type_id = ?", typeId).Joins("left join types on types.id = blogs.type_id").Order("blogs.id").Offset((pageIndex - 1) * 10).Limit(10).Find(&findTwo)

	return &findTwo
}

//只推荐
func ResearchSix(pageIndex int) *[]middleware.BlogType {
	var findTwo []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.recommend,blogs.update_time").Where("recommend = '是'").Joins("left join types on types.id = blogs.type_id").Order("blogs.id").Offset((pageIndex - 1) * 10).Limit(10).Find(&findTwo)

	return &findTwo
}

func BlogView(pageIndex int) (*[]middleware.BlogView, int) {
	var Blog []middleware.BlogView
	var count int
	//var Blog model.Blog
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.description,blogs.recommend,blogs.numb,blogs.update_time,blogs.image").Joins("left join types on types.id = blogs.type_id").Order("blogs.id  DESC").Offset((pageIndex - 1) * 10).Limit(10).Find(&Blog)
	count = len(Blog)
	return &Blog, count
}
func Count() int {
	var count int
	DB.Debug().Table("blogs").Count(&count)
	return count
}

//随机博文
func RandBlog() *[]middleware.BlogType {
	var find []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name").Joins("left join types on types.id = blogs.type_id").Order("RAND()").Limit(5).Find(&find)
	return &find
}

//查询最新推荐博文
func Recommend() *[]middleware.BlogType {
	var find []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name").Where("recommend = '是'").Joins("left join types on types.id = blogs.type_id").Order("blogs.id DESC").Limit(5).Find(&find)
	return &find
}

//查询最热博文
func HotBLog() *[]middleware.BlogType {
	var find []middleware.BlogType
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name").Joins("left join types on types.id = blogs.type_id").Order("blogs.numb DESC").Limit(5).Find(&find)
	return &find
}

//查询所有博客
func SelectBlogAll() *[]model.Blog {
	var blogs []model.Blog
	DB.Find(&blogs)
	return &blogs
}

//按照时间归档

//123
//按照时间归档年份
//1
func ArchYear() *[]int {
	list := make([]int, 12)
	var blog []model.Blog
	DB.Debug().Raw("SELECT * FROM `blogs` GROUP BY YEAR DESC").Scan(&blog)
	for i := 0; i < len(blog); i++ {
		y := blog[i].Year
		list[i] = y
	}
	return &list
}

//按照时间归档月份
//2
func ArchMonth(year int) *[]int {
	list := make([]int, 10)
	var blog []model.Blog
	DB.Debug().Raw("SELECT * FROM `blogs` where year = ? GROUP BY month DESC", year).Scan(&blog)
	for i := 0; i < len(blog); i++ {
		m := blog[i].Month
		list[i] = m
	}
	return &list
}

//查询归档博客
//3
func BlogYear(year int, month int) *[]model.Blog {
	var blog []model.Blog
	DB.Debug().Raw("SELECT * FROM `blogs` where year = ? and month = ?", year, month).Order("`id` DESC").Find(&blog)
	return &blog
}
func BlogResearch(title string, pageIndex int) (*[]middleware.BlogView, int) {
	var Blog []middleware.BlogView
	var count int
	//var Blog model.Blog
	DB.Debug().Table("blogs").Select("blogs.id,blogs.title,types.type_name,blogs.description,blogs.recommend,blogs.numb,blogs.update_time,blogs.image").Where("title like ?", "%"+title+"%").Joins("left join types on types.id = blogs.type_id").Order("blogs.id  DESC").Offset((pageIndex - 1) * 10).Limit(10).Find(&Blog)
	count = len(Blog)
	return &Blog, count
}
