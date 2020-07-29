package model

type Comment struct {
	Id             int    //评论id
	BlogId         int    //博文id
	ReCommentId    int    //父评论id
	CommentName    string //评论人姓名
	ReCommentName  string //父评论人姓名
	CommentEmail   string
	Img            string
	CommentContent string `gorm:"type:text" form:"CommentContent" json:"CommentContent"` //评论内容
	CommentTime    string ` json:"comment_time,omitempty"`
	Parent         *Comment
	Comments       []Comment
}
