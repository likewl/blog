package page

import (
	"blog/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	user := dao.SelectUser()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	blog, count := dao.BlogView(page)
	numb := dao.Count()
	fmt.Println(count)
	tag := dao.RandTag()
	types := dao.Rand()
	recommend := dao.RandBlog()
	c.HTML(http.StatusOK, "pages/index.html", gin.H{
		"user":      user,
		"blog":      blog,
		"count":     count,
		"numb":      numb,
		"type":      types,
		"tag":       tag,
		"page":      page,
		"recommend": recommend,
		"active1":   "active",
		"pageTitle": "一个致力于做JPG的小码农",
	})
}
func PagesHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	if page == 1 {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	user := dao.SelectUser()
	blog, count := dao.BlogView(page)
	numb := dao.Count()
	tag := dao.RandTag()
	types := dao.Rand()
	recommend := dao.RandBlog()
	c.HTML(http.StatusOK, "pages/index.html", gin.H{
		"user":      user,
		"blog":      blog,
		"count":     count,
		"numb":      numb,
		"type":      types,
		"tag":       tag,
		"page":      page,
		"recommend": recommend,
		"pageTitle": "一个致力于做JPG的小码农",
		"active1":   "active",
	})
}
