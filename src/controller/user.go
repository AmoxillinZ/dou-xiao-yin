package controller

import (
	"dou-xiao-yin/src/model"
	"dou-xiao-yin/src/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserLoginResponse 用户登录响应体
type UserLoginResponse struct {
	service.Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserInfoResponse struct {
	service.Response
	User *model.User `json:"user"`
}

type UserRegisterResponse struct {
	service.Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// Login 响应用户登录
func Login(c *gin.Context) {
	// 在 service 根据用户名和密码查找用户
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		// 返回数据
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    user.Token,
		})
	}
}

// UserInfo 相应用户信息获取
func UserInfo(c *gin.Context) {
	//使用 token 来获取用户信息
	id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	user, err := service.GetUserInfo(id, token)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: service.Response{StatusCode: 0},
			User:     user,
		})
	}
}

// Register 响应用户注册
func Register(c *gin.Context) {
	// 前置检验合法性
	username := c.Query("username")
	password := c.Query("password")

	if len(username) > 32 || len(username) < 1 {
		err := errors.New("用户名长度不合法")
		c.JSON(http.StatusBadRequest, UserRegisterResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	if len(password) > 32 || len(password) < 5 {
		err := errors.New("密码长度不合法")
		c.JSON(http.StatusBadRequest, UserRegisterResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	user, err := service.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserRegisterResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: service.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    user.Token,
	})

}
