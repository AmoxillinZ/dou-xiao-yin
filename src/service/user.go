package service

import (
	"dou-xiao-yin/src/mapper"
	"dou-xiao-yin/src/model"
	"errors"
	"fmt"
)

func UserLogin(username string, password string) (*model.User, error) {
	// 根据用户名称查询用户
	user, err := mapper.GetUserByUsername(username)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("没有查询到用户")
	}

	//验证密码
	if user.Password != password {
		return nil, errors.New("用户密码错误")
	}

	//为该用户生成 Token, 这里回头使用 jwt? 来生成 token
	token := username + password

	// 写入 Token
	err = mapper.UpdateUserToken(user, token)
	if err != nil {
		return nil, errors.New("更新用户 Token 失败")
	}

	// 重新获取用户
	user, err = mapper.GetUserById(user.Id)
	if err != nil {
		return nil, errors.New("获取用户失败")
	}
	return user, nil
}

func GetUserInfo(id int, token string) (*model.User, error) {
	// 根据 Id 得到数据
	user, err := mapper.GetUserById(id)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}

	// 进行 token 的权限认证
	if user.Token != token {
		return nil, errors.New("用户 token 认证失败")
	}

	// 返回数据
	return user, nil
}
