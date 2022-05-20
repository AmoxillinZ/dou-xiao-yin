/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 15:04
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
)

func GetUserById(id int) (*model.User, error) {
	usr := &model.User{}
	db := config.GetDefaultDb()
	result := db.Where("id = ?", id).Find(&usr)
	return usr, result.Error
}

func GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	db := config.GetDefaultDb()
	result := db.Where("username = ?", username).Take(&user)
	return user, result.Error
}

func UpdateUserToken(user *model.User, token string) error {
	db := config.GetDefaultDb()
	result := db.Model(&user).Update("token", token)
	return result.Error
}

func AddUser(user *model.User) error {
	db := config.GetDefaultDb()
	result := db.Create(&user)
	return result.Error
}
