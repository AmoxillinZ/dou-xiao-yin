/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:06
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/model"
	"gorm.io/gorm"
)

func IsFollowed(userId int, toUserId int) bool {
	// 如果userId == 0表示当前处于未登录状态，直接返回false
	if userId == 0 {
		return false
	}
	db := config.GetDefaultDb()
	// 如果记录已存在，返回true
	if result := db.Where("user_id = ? and to_user_id = ?", userId, toUserId).Find(&model.Relation{}); result.Error == nil && result.RowsAffected > 0 {
		return true
	}
	return false
}

// FollowAction : 关注的3个操作，合并为一个事务
func FollowAction(userId int, toUserId int) error {
	db := config.GetDefaultDb()
	// 执行事务，以下为一个原子操作
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1.在relation表中新增一条记录
		relation := &model.Relation{UserId: userId, ToUserId: toUserId}
		if err := db.Create(&relation).Error; err != nil {
			return err
		}
		// 2.用户toUser粉丝数+1
		toUser := &model.User{Id: toUserId}
		if err := db.Model(&toUser).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}
		// 3.用户user关注数+1
		user := &model.User{Id: userId}
		if err := db.Model(&user).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UnFollowAction : 取关的三个操作，合并为一个事务
func UnFollowAction(userId int, toUserId int) error {
	db := config.GetDefaultDb()
	// 执行事务，以下为一个原子操作
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1.在relation表中删除一条记录
		if err := db.Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(model.Relation{}).Error; err != nil {
			return err
		}
		// 2.用户toUser粉丝数-1
		toUser := &model.User{Id: toUserId}
		if err := db.Model(&toUser).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}
		// 3.用户user关注数-1
		user := &model.User{Id: userId}
		if err := db.Model(&user).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

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

// FollowList : 某个用户的关注列表
func FollowList(userId int) []*json_model.User {
	fl := make([]*json_model.User, 0)
	db := config.GetDefaultDb()
	db.Model(&model.Relation{}).Select("to_user_id id, username, follow_count, follower_count").
		Joins("inner join `user` on `user`.id = relation.to_user_id").
		Where("relation.user_id = ?", userId).
		Find(&fl)
	return fl
}

// FollowerList : 某个用户的粉丝列表
func FollowerList(userId int) []*json_model.User {
	fl := make([]*json_model.User, 0)
	db := config.GetDefaultDb()
	db.Model(&model.Relation{}).Select("user_id id, username, follow_count, follower_count").
		Joins("inner join `user` on `user`.id = relation.user_id").
		Where("relation.to_user_id = ?", userId).
		Find(&fl)
	return fl
}
