package admin

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TypeHandler(c *gin.Context) {
	user := dao.SelectUser()
	all := dao.SelectName()
	id, _ := strconv.Atoi(c.Query("id"))
	//删除分类名称
	if id != 0 {
		dao.DeleteType(id)
		all = dao.SelectName()
		c.HTML(http.StatusCreated, "admin/types.html", gin.H{
			"all":     *all,
			"msg":     "删除成功",
			"choice2": "active",
			"user":    user,
		})
		return
	}

	//fmt.Println(*all)
	c.HTML(http.StatusOK, "admin/types.html", gin.H{
		"all":     *all,
		"choice2": "active",
		"user":    user,
	})
}

func TypeEditHandler(c *gin.Context) {
	user := dao.SelectUser()
	id, _ := strconv.Atoi(c.Query("id"))
	name := c.Query("name")
	editname := c.PostForm("editname")
	editid, _ := strconv.Atoi(c.PostForm("id"))
	//修改分类名称
	if editname != "" {
		dao.UpdateType(editid, editname)
		c.Redirect(http.StatusMovedPermanently, "/admin/type")
		return
	}
	//选择要修改的标题
	c.HTML(http.StatusOK, "types-edit.html", gin.H{
		"id":      id,
		"name":    name,
		"choice2": "active",
		"user":    user,
	})
}

func TypeAddHandler(c *gin.Context) {
	user := dao.SelectUser()
	name := c.PostForm("name")
	//添加分类
	if name != "" {
		dao.CreatType(name)
		c.Redirect(http.StatusMovedPermanently, "/admin/type")
		return
	}
	c.HTML(http.StatusOK, "types-input.html", gin.H{
		"choice2": "active",
		"user":    user,
	})
}
