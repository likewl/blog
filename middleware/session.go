package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//session信息
type U struct {
	Id       uint
	Name     string
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

//将userID注入session
func SetUser(c *gin.Context, user interface{}) {
	session := sessions.Default(c)
	session.Set("user", user)
	session.Save()
	c.Next()
}

// 获取用户登录信息 userID
func GetUser(c *gin.Context) interface{} {
	session := sessions.Default(c)
	user := session.Get("user")
	return user
}

// Authorization 验证用户
func Authorization(c *gin.Context) {

	fmt.Println("中间件开始校验")
	user := GetUser(c)
	fmt.Println(user, "获取user")
	if user == nil || user == "false" {
		fmt.Println(user, "校验错误")
		c.HTML(http.StatusFound, "login.html", nil)
		c.Abort()
		return
	}
	fmt.Println(user, "校验正确")
	c.Next()

}
