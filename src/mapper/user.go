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

func GetUserById(id int) *model.User {
	usr := &model.User{}
	db := config.GetDefaultDb()
	db.Where("id = ?", id).Find(&usr)
	return usr
}
