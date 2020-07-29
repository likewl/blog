package logic

import (
	"blog/dao"
)

func CheckUser(userName string, passWord string) error {
	err := dao.CheckUser(userName, passWord)
	return err
}
