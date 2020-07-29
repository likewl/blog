package page

import (
	"blog/dao"
	"blog/middleware"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentHandler(c *gin.Context) {
	var comment model.Comment
	atoi := c.PostForm("blog.id")
	BlogId, _ := strconv.Atoi(atoi)
	ReCommentId, _ := strconv.Atoi(c.PostForm("parentComment.id"))
	ReCommentName := c.PostForm("reCommentName")
	CommentName := c.PostForm("nickname")
	CommentEmail := c.PostForm("email")
	CommentContent := c.PostForm("content")
	typeName := c.PostForm("typeName")
	user := dao.SelectUser()
	one := dao.SelectCommentByReComment(CommentName, CommentEmail)
	//找博主
	if CommentName == user.Name && CommentEmail == user.UserEmail {
		fmt.Println("进入了1")
		comment.Img = user.UserImg
		//找相同评论人
	} else if CommentName == one.CommentName && CommentEmail == one.CommentEmail {
		fmt.Println("进入了2")
		comment.Img = one.Img
		//随机选头像
	} else {
		fmt.Println("进入了3")
		comment.Img = middleware.GetImg()
	}
	comment.BlogId = BlogId
	comment.ReCommentId = ReCommentId
	comment.CommentName = CommentName
	comment.CommentEmail = CommentEmail
	comment.CommentContent = CommentContent
	comment.ReCommentName = ReCommentName
	comment.CommentTime = middleware.GetTime()
	dao.CreatComment(&comment)
	c.Redirect(http.StatusMovedPermanently, "/types/"+typeName+"/"+atoi+"#comment-container")
}
