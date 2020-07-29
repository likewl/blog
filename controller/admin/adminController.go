package admin

import (
	"blog/dao"
	"blog/logic"
	"blog/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/admin/index")
}
func AdminIndexHandler(c *gin.Context) {
	user := dao.SelectUser()
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"user": user,
	})

}

func LoginHandler(c *gin.Context) {
	user := dao.SelectUser()
	var u middleware.U
	c.Bind(&u)
	err := logic.CheckUser(u.UserName, u.Password)
	fmt.Println(err, "err")
	if err == nil {
		middleware.SetUser(c, u.UserName)
		user := middleware.GetUser(c)
		fmt.Println(user, "登录了")
		c.Redirect(http.StatusMovedPermanently, "/admin")
		return
	} else {
		c.HTML(http.StatusFound, "login.html", gin.H{
			"msg":  "账号或密码错误",
			"user": user,
		})

		return

	}
}

func LogoutHandler(c *gin.Context) {
	user := middleware.GetUser(c)
	fmt.Println(user, "退出前")
	middleware.SetUser(c, "false")

	user = middleware.GetUser(c)
	middleware.SetUser(c, user)

	fmt.Println(user, "退出后")

	c.Redirect(http.StatusMovedPermanently, "/admin")

}
