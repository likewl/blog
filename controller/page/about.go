package page

import (
	"blog/dao"
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func AboutHandler(c *gin.Context) {

	user := dao.SelectUser()
	content1 := user.Content
	content := middleware.Markdown(content1)
	c.HTML(http.StatusOK, "about.html", gin.H{
		"user":      user,
		"pageTitle": "关于" + user.Name,
		"content": template.HTML(content),
		"active5": "active",
	})
}
