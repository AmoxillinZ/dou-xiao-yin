package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 私钥
var jwtKey = []byte("yi-fan-feng-shun-er-long-xi-zhu-san-yang-kai-tai-si-wu-liu-qi-ba-jiu-shi")

// Claims jwt.StandardClaims 提供到期时间等字段
type Claims struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	jwt.StandardClaims
}

// SetToken 颁发 token
func SetToken(id int, username string, password string) (string, error) {
	// 设置过期时间（七天一登陆？是否添加 token 刷新？）
	expireTime := time.Now().Add(7 * 24 * time.Hour)

	// 封装
	claims := &Claims{
		Id:       id,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "dou-xiao-yin-backend", // 签名颁发者
			Subject:   "user token",           //签名主题
		},
	}

	// 设置加密的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成 tokenString
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("生成 token 失败")
	}
	return tokenString, nil
}

// ParseToken 解析 token，返回用户信息，按需使用
func ParseToken(tokenString string) (id int, username string, password string, e error) {

	if tokenString == "" {
		fmt.Println("token 为空")
		return 0, "", "", errors.New("用户权限不足")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("token 无法解析或 token 已过期")
		return 0, "", "", errors.New("用户权限不足")
	}

	return claims.Id, claims.Username, claims.Password, nil
}
