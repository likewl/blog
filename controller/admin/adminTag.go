package admin

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//1是成功
//0是无操作
func TagHandler(c *gin.Context) {
	user := dao.SelectUser()
	//all := dao.SelectTag()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	all, count := dao.TagPage(page)
	id, _ := strconv.Atoi(c.Query("id"))
	msg := c.Query("msg")
	//空字符串是无信息
	//1是删除成功
	//2是修改成功
	//3是创建成功
	if msg == "" {
		msg = ""
	} else if msg == "1" {
		msg = "删除成功"
	} else if msg == "1" {
		msg = "修改成功"
	} else if msg == "3" {
		msg = "创建成功"
	}
	//删除标签名称
	if id != 0 {
		dao.DeleteTag(id)
		dao.DelBlogTag(id)
		pages := c.DefaultQuery("page", "1")
		url := "/admin/tag?msg=1&page=" + pages
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	c.HTML(http.StatusOK, "admin/tags.html", gin.H{
		"choice3": "active",
		"msg":     msg,
		"page":    page,
		"count":   count,
		"all":     *all,
		"user":    user,
	})

}

func TagEditHandler(c *gin.Context) {
	user := dao.SelectUser()
	id, _ := strconv.Atoi(c.Query("id"))
	page := c.DefaultQuery("page", "1")
	name := c.Query("name")
	editName := c.PostForm("editName")
	editId, _ := strconv.Atoi(c.PostForm("id"))
	msg := c.PostForm("msg")
	if msg == "" {
		msg = ""
	} else if msg == "1" {
		msg = "修改成功"
	}
	//修改标签名称
	if editName != "" {
		dao.UpdateTag(editId, editName)
		pages := c.DefaultPostForm("page", "1")
		url := "/admin/tag?msg=1&page=" + pages
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	//选择要修改的标题
	c.HTML(http.StatusOK, "tags-edit.html", gin.H{
		"page":    page,
		"id":      id,
		"name":    name,
		"choice3": "active",
		"user":    user,
	})
}

func TagAddHandler(c *gin.Context) {
	user := dao.SelectUser()
	page := c.DefaultQuery("page", "1")
	name := c.PostForm("name")
	msg := c.PostForm("msg")
	if msg == "" {
		msg = ""
	} else if msg == "1" {
		msg = "创建成功"
	}
	//添加标签
	if name != "" {
		dao.CreatTag(name)
		pages := c.DefaultPostForm("page", "1")
		url := "/admin/tag?msg=3&page=" + pages
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	c.HTML(http.StatusOK, "tags-input.html", gin.H{
		"page":    page,
		"choice3": "active",
		"user":    user,
	})
}
