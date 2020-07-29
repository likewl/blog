package page

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TagsHandler(c *gin.Context) {
	tags := dao.SelectTag()
	user := dao.SelectUser()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	blog, count := dao.BlogView(page)
	numb := dao.CountTag()
	c.HTML(http.StatusOK, "tags.html", gin.H{
		"user":      user,
		"blog":      blog,
		"count":     count,
		"numb":      numb,
		"tag":       tags,
		"page":      page,
		"active3":   "active",
		"pageTitle": "标签页面",
	})
}

func TagsNameHandler(c *gin.Context) {
	name := c.Param("name")
	tags := dao.SelectTag()
	user := dao.SelectUser()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	blog, count := dao.SelectTagNameAll(name, page)
	numb := dao.CountTag()
	c.HTML(http.StatusOK, "tag.html", gin.H{
		"user":      user,
		"blog":      blog,
		"count":     count,
		"numb":      numb,
		"name":      name,
		"tag":       tags,
		"page":      page,
		"active3":   "active",
		"pageTitle": "标签页面",
	})
}
