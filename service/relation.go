/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 11:24
 */

package service

import (
	"dou-xiao-yin/json_model"
	"dou-xiao-yin/model"
	"errors"
)

type RelationResponse struct {
	json_model.Response
	VideoList []*json_model.User `json:"user_list,omitempty"`
}

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
	if model.IsFollowed(userId, toUserId) {
		return errors.New("已经关注过")
	}
	return model.FollowAction(userId, toUserId)
}

// UnFollowAction : 取关操作
func UnFollowAction(userId int, toUserId int) error {
	return model.UnFollowAction(userId, toUserId)
}

// FollowList : 用户的关注列表
func FollowList(userId int, loginId int) []*json_model.User {
	// 获得目标用户的关注列表
	followListUser := model.FollowList(userId)
	// 如果当前未登录用户，则所有的isFollowed都为false，可直接返回
	if loginId == 0 {
		return followListUser
	}
	// 存到map中，便于判断isFollow
	mp := make(map[int]*json_model.User)
	for _, usr := range followListUser {
		mp[usr.Id] = usr
	}
	// 获得当前登录用户的关注列表
	followListLogin := model.FollowList(loginId)
	// 如果当前用户的关注列表中的某一个用户也在目标用户的列表中，就将其isFollow设置为true
	for _, usr := range followListLogin {
		if _, exists := mp[usr.Id]; exists {
			mp[usr.Id].IsFollow = true
		}
	}
	return followListUser
}

// FollowerList : 用户的粉丝列表
func FollowerList(userId int, loginId int) []*json_model.User {
	// 获得目标用户的粉丝列表
	followerListUser := model.FollowerList(userId)
	// 如果当前未登录用户，则所有的isFollowed都为false，可直接返回
	if loginId == 0 {
		return followerListUser
	}
	// 存到map中，便于判断isFollow
	mp := make(map[int]*json_model.User)
	for _, usr := range followerListUser {
		mp[usr.Id] = usr
	}
	// 获得当前登录用户的关注列表
	followListLogin := model.FollowList(loginId)
	// 如果当前用户的关注列表中的某一个用户也在目标用户的列表中，就将其isFollow设置为true
	for _, usr := range followListLogin {
		if _, exists := mp[usr.Id]; exists {
			mp[usr.Id].IsFollow = true
		}
	}
	return followerListUser
}
