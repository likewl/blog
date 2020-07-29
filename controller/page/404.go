package page

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {

	user := dao.SelectUser()
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"user":      user,
		"pageTitle": "404 NOT FOUND",
	})
}
