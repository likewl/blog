package dao

import (
	"blog/model"
	"fmt"
)

var Comments []model.Comment

func CreatComment(comment *model.Comment) {
	DB.Debug().Create(comment)
}

//找到服评论
func SelectCommentByBlog(id int) *[]model.Comment {
	var comments []model.Comment
	DB.Debug().Where("re_comment_id = -1 AND blog_id = ?", id).Find(&comments)
	return &comments
}

//找对应的子评论
func SelectCommentByComment(id int, count int) *[]model.Comment {
	if count != 0 {
		fmt.Println("if判断1")
		var comments []model.Comment
		DB.Debug().Where("re_comment_id = ?", id).Find(&comments)
		count = len(comments)
		if count != 0 {
			fmt.Println("if判断2")
			for _, comment := range comments {
				fmt.Println("for判断1")
				Comments = append(Comments, comment)
				SelectCommentByComment(comment.Id, count)
			}
		}
	}
	fmt.Println(Comments)
	return &Comments

}

func ClearComments() bool {
	var comments []model.Comment
	Comments = comments
	return true
}

//查找符合要求的评论，为其复制头像
func SelectCommentByReComment(name string, email string) *model.Comment {
	var comment model.Comment
	DB.Debug().Where(" comment_name = ? AND comment_email = ?", name, email).First(&comment)
	return &comment
}
