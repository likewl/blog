package page

import (
	"blog/dao"
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NotesController(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	user := dao.SelectUser()

	types := dao.SelectName()
	numb := dao.CountType()
	type1, count, err := dao.SelectBlogsByTypeId(16, page)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}
	var blogs []middleware.BlogView
	for _, blog := range *type1 {
		blog.Description = middleware.Markdown(blog.Description)
		blogs = append(blogs, blog)
	}
	tag := dao.RandTag()
	recommend := dao.RandBlog()
	hotBLog := dao.HotBLog()
	c.HTML(http.StatusOK, "note", gin.H{
		"user":      user,
		"count":     count,
		"type":      types,
		"tag":       tag,
		"types":     blogs,
		"numb":      numb,
		"page":      page,
		"recommend": recommend,
		"hotBlog":   hotBLog,
		"active6":   "active",
		"pageTitle": "叨吧叨",
	})
}
