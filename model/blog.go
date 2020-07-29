package model

type Blog struct {
	Id           int
	Title        string    `form:"title" json:"title"`                          //标题
	Content      string    `gorm:"type:longtext" form:"content" json:"content"` //内容
	Image        string    `gorm:"type:text" form:"image" json:"image"`         //首图
	Numb         int       `form:"numb" json:"numb"`                            //次数
	Flag         string    `form:"flag" json:"flag"`                            //原创或者转载
	Appreciation string    `form:"appreciation" json:"appreciation"`            //打赏开关
	ShareInfo    string    `form:"shareInfo" json:"shareInfo"`                  //版权开关
	Comment      string    `form:"comment" json:"comment"`                      //评论开关
	Recommend    string    `form:"recommend" json:"recommend"`                  //推荐开关
	Description  string    `gorm:"type:text" form:"description" json:"description"`
	CreatTime    string    `form:"creatTime"  json:"creat_time,omitempty"`   //创建时间
	UpdateTime   string    `form:"updateTime"  json:"update_time,omitempty"` //更新时间
	Comments     []Comment `form:"comments" json:"comments"`
	TypeId       int       `form:"typeId" json:"typeId"`
	Type         Type      `form:"type" gorm:"ForeignKey:TypeId"`
	Year         int       `form:"year" json:"year"`
	Month        int       `form:"month" json:"month"`
	Day          int       `form:"day" json:"day"`
}
