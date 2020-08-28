package page

import (
	"blog/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ArchivesHandler(c *gin.Context) {
	query := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(query)
	user := dao.SelectUser()
	numb := dao.Count()
	year,A := dao.ArchYear(page)
	count := dao.ArchYearCount()
	fmt.Println(count)
	fmt.Println(page)
	c.HTML(http.StatusOK, "archives.html", gin.H{
		"user":      user,
		"pageTitle": "文章归档",
		"numb":      numb,
		"year":      year,
		"title":      A,
		"page":      page,
		"count":     count,
		"active4":   "active",
	})
}
