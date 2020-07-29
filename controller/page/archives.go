package page

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArchivesHandler(c *gin.Context) {

	user := dao.SelectUser()
	numb := dao.Count()
	year := dao.ArchYear()
	c.HTML(http.StatusOK, "archives.html", gin.H{
		"user":      user,
		"pageTitle": "文章归档",
		"numb":      numb,
		"year":      year,
		"active4":      "active",
	})
}
