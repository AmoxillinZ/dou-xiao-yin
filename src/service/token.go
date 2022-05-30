package service

import (
	"dou-xiao-yin/src/utils"
	"errors"
)

func TokenVerify(userId int, token string) error {
	if userId == 0 {
		return errors.New("请求中没有user_id")
	}
	verifiedUserId, err := utils.GetIdFromToken(token)
	if err != nil || verifiedUserId != userId {
		return errors.New("校验用户失败")
	}
	return nil
}
