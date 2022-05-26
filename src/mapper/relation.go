/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:06
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
)

// AddToRelation : 在relation表中新增一条记录
func AddToRelation(userId int, toUserId int) error {
	db := config.GetDefaultDb()
	relation := &model.Relation{UserId: userId, ToUserId: toUserId}
	result := db.Create(&relation)
	return result.Error
}

// DeleteFromRelation : 在relation表中删除一条记录
func DeleteFromRelation(userId int, toUserId int) error {
	db := config.GetDefaultDb()
	result := db.Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(model.Relation{})
	return result.Error
}
