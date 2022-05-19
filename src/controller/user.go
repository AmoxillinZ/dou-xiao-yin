package controller

import (
	"dou-xiao-yin/src/model"
	"dou-xiao-yin/src/service"
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
