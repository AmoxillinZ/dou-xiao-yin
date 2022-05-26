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
		return FollowAction(userId, toUserId)
	} else if actionType == 2 {
		return UnFollowAction(userId, toUserId)
	} else {
		return errors.New("actionType传入类型错误")
	}
}

// FollowAction TODO：relation添加记录、user的关注数+1、toUser的粉丝数+1，这三个操作是一个原子操作，封装事务？
// FollowAction : 关注操作
func FollowAction(userId int, toUserId int) error {
	// relation表中新增记录
	if err := mapper.AddToRelation(userId, toUserId); err != nil {
		return err
	}
	// toUserId对应的用户粉丝总数(follower_count)加一
	if err := mapper.IncreaseFollowerCount(toUserId); err != nil {
		return err
	}
	// userId对应的用户关注总数(follow_count)加一
	if err := mapper.IncreaseFollowCount(userId); err != nil {
		return err
	}
	return nil
}

// UnFollowAction : 取关操作
func UnFollowAction(userId int, toUserId int) error {
	// relation表中删除记录
	if err := mapper.DeleteFromRelation(userId, toUserId); err != nil {
		return err
	}
	// toUserId对应的用户粉丝总数(follower_count)减一
	if err := mapper.DecreaseFollowerCount(toUserId); err != nil {
		return err
	}
	// userId对应的用户关注总数(follow_count)减一
	if err := mapper.DecreaseFollowCount(userId); err != nil {
		return err
	}
	return nil
}
