package admin

import (
	"blog/dao"
	"blog/middleware"
	"blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserEditHandler(c *gin.Context) {
	user := dao.SelectUser()
	var users model.User
	c.ShouldBind(&users)
	if users.Password != "" {
		users.UserType = user.UserType
		users.UserCreate = user.UserCreate
		users.UserUpdate = middleware.GetTime()
		dao.UpdateUser(&users)
		middleware.SetUser(c, "false")

		getUser := middleware.GetUser(c)
		middleware.SetUser(c, getUser)
		c.Redirect(http.StatusMovedPermanently, "/admin")
		return
	}
	c.HTML(http.StatusOK, "user-edit.html", gin.H{
		"user": user,
	})
}
