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

func GetUserByToken(token string) (*model.User, error) {
	user := &model.User{}
	db := config.GetDefaultDb()
	result := db.Where("token = ?", token).Take(&user)
	return user, result.Error
}

// GetUserIdByToken : 根据token查询userId。
// 可以通过jwt_util直接获取,不需要查数据库
func GetUserIdByToken(token string) int {
	db := config.GetDefaultDb()
	userWithId := &model.User{}
	db.Select("id").Where("token = ?", token).Limit(1).Take(&userWithId)
	return userWithId.Id
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
