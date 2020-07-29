package page

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TypesHandler(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	//if page != 1 {
	//	c.Redirect(http.StatusMovedPermanently, "/types")
	//	return
	//}
	types := dao.SelectName()
	user := dao.SelectUser()
	blog, count := dao.BlogView(page)
	numb := dao.CountType()
	c.HTML(http.StatusOK, "types.html", gin.H{
		"user":      user,
		"blog":      blog,
		"count":     count,
		"numb":      numb,
		"type":      types,
		"page":      page,
		"active2":   "active",
		"pageTitle": "分类页面",
	})
}
func TypesNameHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	user := dao.SelectUser()
	name := c.Param("name")
	types := dao.SelectName()
	numb := dao.CountType()
	type1, count, err := dao.SelectNameAll(name, page)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}
	c.HTML(http.StatusOK, "type.html", gin.H{
		"user":      user,
		"count":     count,
		"type":      types,
		"types":     type1,
		"name":      name,
		"numb":      numb,
		"page":      page,
		"active2":   "active",
		"pageTitle": name + "分类页面",
	})
}
