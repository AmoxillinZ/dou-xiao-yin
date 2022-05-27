/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 11:24
 */

package service

import (
	"dou-xiao-yin/src/mapper"
	"errors"
)

func RelationAction(userId int, toUserId int, actionType int) error {
	if userId == toUserId {
		return errors.New("不能关注自己")
	}
	if actionType == 1 {
		err := FollowAction(userId, toUserId)
		return err
	} else if actionType == 2 {
		return UnFollowAction(userId, toUserId)
	} else {
		return errors.New("actionType传入类型错误")
	}
}

// FollowAction : 关注操作
func FollowAction(userId int, toUserId int) error {
	if mapper.IsFollowed(userId, toUserId) {
		return errors.New("已经关注过")
	}
	return mapper.FollowAction(userId, toUserId)
}

// UnFollowAction : 取关操作
func UnFollowAction(userId int, toUserId int) error {
	return mapper.UnFollowAction(userId, toUserId)
}
