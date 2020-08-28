package dao

import (
	"blog/model"
	"fmt"
	"time"
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

//判断是否有子评论
func BoolReCommentByCommentId(id int) bool {
	var flag bool
	var count int
	DB.Debug().Table("comments").Where("re_comment_id = ?", id).Count(&count)
	if count != 0 {
		flag = true
	}
	return flag

}

//求文章浏览总数总和
func SUMBlogNUM() int {
	var count struct{
		Numb int
	}
	DB.Debug().Table("blogs").Select("SUM(numb) as numb").Scan(&count)
	fmt.Println(count.Numb)
	return count.Numb

}
//求文章个数总和
func SUMBlog() int {
	var count struct{
		Numb int
	}
	DB.Debug().Table("blogs").Select("count(*) as numb").Where("type_id != 16").Scan(&count)
	fmt.Println(count.Numb)
	return count.Numb

}
//求文章个数总和
func NewBlog() int {
	var count struct{
		Numb int
	}
	DB.Debug().Table("blogs").Select("count(*) as numb").Where("creat_time like ?",time.Now().Format("2006-01-02")+"%").Scan(&count)
	fmt.Println(count.Numb)
	return count.Numb

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
//最新留言
func SelectNewComment() *[]model.Comment {
	var comment []model.Comment
	DB.Debug().Limit(5).Order("id desc").Find(&comment)
	return &comment
}


