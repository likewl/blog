package model

type User struct {
	Id         uint   `form:"id" json:"id"`
	BlogName   string `form:"blogName" json:"blogName"`
	Name       string `form:"name" json:"name"`
	UserName   string `form:"userName" json:"userName"`
	Password   string `form:"password" json:"password"`
	UserEmail  string `form:"userEmail" json:"userEmail"`
	UserType   int
	UserImg    string `form:"userImg" json:"userImg"`
	Content    string `gorm:"type:longtext" form:"content" json:"content"`
	Favorite   string `gorm:"type:text" form:"favorite" json:"favorite"`
	Language   string `gorm:"type:text" form:"language" json:"language"`
	Describe   string `gorm:"type:text" form:"describe" json:"describe"`
	UserCreate string ` json:"user_create,omitempty"`
	UserUpdate string ` json:"user_update,omitempty"`
	Blogs      []Blog
}
