package service

import (
	"dou-xiao-yin/src/mapper"
	"errors"
)

func TokenVerify(userId int, token string) error {
	user, err := mapper.GetUserById(userId)
	if err != nil {
		return errors.New("没有找到用户")
	}
	if user.Token != token {
		return errors.New("用户令牌无效")
	}
	return nil
}
