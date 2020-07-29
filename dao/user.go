package dao

import "blog/model"

func SelectUser() *model.User {
	var user model.User
	DB.Debug().Find(&user)
	user.Password = ""
	return &user
}
func UpdateUser(user *model.User) {
	DB.Debug().Where("id=?", user.Id).Save(user)
}
