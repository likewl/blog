package page

import (
	"blog/dao"
	"blog/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func BlogHandler(c *gin.Context) {
	name := c.Param("name")
	id, _ := strconv.Atoi(c.Param("id"))

	blog, err := dao.SelectOne(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}
	oneType := dao.SelectOneType(blog.TypeId)
	comments := dao.SelectCommentByBlog(id)
	fmt.Println(err, "错误111111111")
	tagName := dao.SelectTagNameBlog(id)
	blog.Numb = blog.Numb + 1
	dao.UpdateBlog(blog)
	title := blog.Title
	content1 := blog.Content
	//content := middleware.MarkdownToHTML(content1)
	content := middleware.Markdown(content1)
	user := dao.SelectUser()
	host := c.Request.Host
	hotBLog := dao.HotBLog()
	c.HTML(http.StatusOK, "blog.html", gin.H{
		"user":      user,
		"blog":      blog,
		"pageTitle": title,
		"content":   template.HTML(content),
		"comments":  comments,
		"type":      name,
		"typeName":  oneType.TypeName,
		"Type":      oneType,
		"tagName":   tagName,
		"host":      host,
		"hotBlog":   hotBLog,
	})
}
