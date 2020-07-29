package dao

import (
	"blog/model"
)

func CheckUser(username string, password string) error {
	var user model.User
	err := DB.Where(map[string]interface{}{"user_name": username, "password": password}).First(&user).Error
	return err

}
