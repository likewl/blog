package admin

import (
	"blog/dao"
	"blog/middleware"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//1是成功
//0是无操作
func BlogHandler(c *gin.Context) {
	user := dao.SelectUser()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	all, count := dao.BlogType(page)
	fmt.Println(all)
	fmt.Println(count, "ceshi")
	id, _ := strconv.Atoi(c.Query("id"))
	msg := c.Query("msg")
	typeName := dao.SelectName()
	var research model.Blog
	c.ShouldBind(&research)
	typeId, _ := strconv.Atoi(c.PostForm("typeId"))
	research.TypeId = typeId
	//空字符串是无信息
	//1是删除成功
	//2是修改成功
	//3是创建成功
	if msg == "" {
		msg = ""
	} else if msg == "1" {
		msg = "删除成功"
	} else if msg == "2" {
		msg = "修改成功"
	} else if msg == "3" {
		msg = "创建成功"
	}
	fmt.Println(id)
	//删除标签名称
	if id != 0 {
		dao.DeleteBlog(id)
		dao.DelTagBlog(id)
		pages := c.DefaultQuery("page", "1")
		url := "/admin/blog?msg=1&page=" + pages
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	//搜索
	if research.Title != "" || research.TypeId != 0 || research.Recommend != "" {
		//只选择了标题
		if research.TypeId == 0 && research.Recommend == "" {
			one := dao.ResearchOne(research.Title, page)
			c.HTML(http.StatusOK, "blogs.html", gin.H{
				"msg":      msg,
				"typeName": typeName,
				"page":     page,
				"count":    count,
				"all":      *one,
				"user":     user,
				"choice1":  "active",
			})
			return
			//标题和类型
		} else if research.TypeId != 0 && research.Recommend == "" {
			two := dao.ResearchTwo(research.Title, research.TypeId, page)
			c.HTML(http.StatusOK, "blogs.html", gin.H{
				"msg":      msg,
				"typeName": typeName,
				"page":     page,
				"count":    count,
				"all":      *two,
				"user":     user,
				"choice1":  "active",
			})
			return
			//标题和推荐
		} else if research.TypeId == 0 && research.Recommend == "是" {
			three := dao.ResearchThree(research.Title, page)
			c.HTML(http.StatusOK, "blogs.html", gin.H{
				"msg":      msg,
				"typeName": typeName,
				"page":     page,
				"count":    count,
				"all":      *three,
				"user":     user,
				"choice1":  "active",
			})
			return
			//都选择了
		} else if research.TypeId != 0 && research.Recommend == "是" {
			four := dao.ResearchFour(research.Title, research.TypeId, page)
			c.HTML(http.StatusOK, "blogs.html", gin.H{
				"msg":      msg,
				"typeName": typeName,
				"page":     page,
				"count":    count,
				"all":      *four,
				"user":     user,
				"choice1":  "active",
			})
			return
			//只分类
		} else if research.TypeId != 0 && research.Recommend == "是" {
			four := dao.ResearchFive(research.TypeId, page)
			c.HTML(http.StatusOK, "blogs.html", gin.H{
				"msg":      msg,
				"typeName": typeName,
				"page":     page,
				"count":    count,
				"all":      *four,
				"user":     user,
				"choice1":  "active",
			})
			return
			//只推荐
		} else if research.TypeId != 0 && research.Recommend == "是" {
			four := dao.ResearchSix(page)
			c.HTML(http.StatusOK, "blogs.html", gin.H{
				"msg":      msg,
				"typeName": typeName,
				"page":     page,
				"count":    count,
				"all":      *four,
				"user":     user,
				"choice1":  "active",
			})
			return
		}

	}

	c.HTML(http.StatusOK, "blogs.html", gin.H{
		"msg":      msg,
		"typeName": typeName,
		"page":     page,
		"count":    count,
		"all":      *all,
		"user":     user,
		"choice1":  "active",
	})

}

func BlogEditHandler(c *gin.Context) {
	//修改后
	user := dao.SelectUser()
	var updateBlog model.Blog
	c.ShouldBind(&updateBlog)

	id, _ := strconv.Atoi(c.Query("id"))
	page := c.DefaultQuery("page", "1")
	typeName := dao.SelectName()
	tagName := dao.SelectTag()
	msg := c.PostForm("msg")

	fmt.Println(updateBlog, "ceshi")
	if msg == "" {
		msg = ""
	} else if msg == "1" {
		msg = "修改成功"
	}
	//修改标签名称
	if updateBlog.Title != "" {
		typeId, _ := strconv.Atoi(c.PostForm("typeId"))
		blogid, _ := strconv.Atoi(c.PostForm("id"))
		dao.DelTagBlog(blogid)
		updateBlog.Id = blogid
		updateBlog.TypeId = typeId
		updateBlog.UpdateTime = middleware.GetTime()
		one, _ := dao.SelectOne(blogid)
		updateBlog.CreatTime = one.CreatTime
		updateBlog.Year = one.Year
		updateBlog.Month = one.Month
		updateBlog.Day = one.Day
		updateBlog.Numb = one.Numb
		dao.UpdateBlog(&updateBlog)

		//添加标签

		tags := c.PostFormArray("tags")
		for _, id := range tags {
			atoi, _ := strconv.Atoi(id)
			dao.CreatTagBlog(atoi, updateBlog.Id)
		}
		pages := c.DefaultPostForm("page", "1")
		url := "/admin/blog?msg=2&page=" + pages
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	//选择要修改的标题
	one, _ := dao.SelectOne(id)
	oneTypeName := dao.SelectOneTypeByTypeId(one.TypeId)

	c.HTML(http.StatusOK, "blogs-edit.html", gin.H{
		"page":        page,
		"id":          id,
		"one":         one,
		"typeName":    typeName,
		"oneTypeName": oneTypeName,
		"tagName":     tagName,
		"user":        user,
		"choice1":     "active",
	})
}

//添加博文
func BlogAddHandler(c *gin.Context) {
	user := dao.SelectUser()
	var creatBlog middleware.CreatBlog

	page := c.DefaultQuery("page", "1")
	c.ShouldBind(&creatBlog)

	typeName := dao.SelectName()
	tagName := dao.SelectTag()
	if creatBlog.Msg == "" {
		creatBlog.Msg = ""
	} else if creatBlog.Msg == "1" {
		creatBlog.Msg = "创建成功"
	}
	//添加博文
	if creatBlog.Blog.Title != "" {
		typeId := c.PostForm("typeId")
		typeId1, _ := strconv.Atoi(typeId)
		creatBlog.Blog.TypeId = typeId1
		//时间获取
		creatBlog.Blog.CreatTime = middleware.GetTime()
		creatBlog.Blog.UpdateTime = middleware.GetTime()
		year, month, day := middleware.GetTimes()
		creatBlog.Blog.Year = year
		creatBlog.Blog.Month = month
		creatBlog.Blog.Day = day
		//添加标签
		dao.CreatBlog(&creatBlog.Blog)
		last := dao.Last()
		tags := c.PostFormArray("tags")
		for _, id := range tags {
			atoi, _ := strconv.Atoi(id)
			dao.CreatTagBlog(atoi, last.Id)
		}
		pages := c.DefaultPostForm("page", "1")
		url := "/admin/blog?msg=3&page=" + pages
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	c.HTML(http.StatusOK, "blogs-input.html", gin.H{
		"page":     page,
		"typeName": typeName,
		"tagName":  tagName,
		"user":     user,
		"choice1":  "active",
	})
}
