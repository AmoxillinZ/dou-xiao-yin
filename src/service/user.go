package service

import (
	"dou-xiao-yin/src/mapper"
	"dou-xiao-yin/src/model"
	"dou-xiao-yin/src/utils"
	"errors"
	"fmt"
)

func UserLogin(username string, password string) (*model.User, error) {
	// 根据用户名称查询用户
	user, err := mapper.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("没有查询到用户")
	}

	//验证密码
	if user.Password != password {
		return nil, errors.New("用户密码错误")
	}

	//为该用户生成 Token, 这里回头使用 jwt? 来生成 token
	token, err := utils.SetToken(user.Id, user.Username, user.Password)
	if err != nil {
		return nil, errors.New("生成 token 失败")
	}

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

func UserRegister(username string, password string) (*model.User, error) {
	// 注册新用户
	newUser := &model.User{Username: username, Password: password}
	err := mapper.AddUser(newUser)
	if err != nil {
		return nil, errors.New("无法添加用户")
	}

	// 获取用户 id 生成 token
	user, err := mapper.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("无法根据用户名获取用户信息")
	}

	id := user.Id
	token, err := utils.SetToken(id, username, password)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("生成 token 失败")
	}

	// 写入 Token
	err = mapper.UpdateUserToken(user, token)
	if err != nil {
		return nil, errors.New("更新用户 Token 失败")
	}

	// 重新获取用户
	user, err = mapper.GetUserById(id)
	if err != nil {
		return nil, errors.New("无法根据 id 获取用户信息")
	}

	return user, nil
}

// VerifyUser : 鉴权，校验传入的userId和token的一致性*/
func VerifyUser(userId int, token string) bool {
	//通过token校验的用户id与传入的相等，鉴权通过
	parsedId, _, _, _ := utils.ParseToken(token)
	if userId == parsedId {
		return true
	}
	return false
}
