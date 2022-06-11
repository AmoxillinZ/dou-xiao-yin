/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 20:32
 */

package model

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func TestGetUserIdByToken(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6InpoaCIsInBhc3N3b3JkIjoiMjMzMzMzIiwiZXhwIjoxNjUzODIzNDc2LCJpYXQiOjE2NTMyMTg2NzYsImlzcyI6ImRvdS14aWFvLXlpbi1iYWNrZW5kIiwic3ViIjoidXNlciB0b2tlbiJ9.g1aakHj-VVOyE45ZjTxxjTlTPt89eCugZluqJvkTsVI"
	id := GetUserIdByToken(token)
	fmt.Println(id)
}

func TestGetUserByToken(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	users, err := GetUserByToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6InpoaCIsInBhc3N3b3JkIjoiMjMzMzMzIiwiZXhwIjoxNjUzNzQ3Mjk0LCJpYXQiOjE2NTMxNDI0OTQsImlzcyI6ImRvdS14aWFvLXlpbi1iYWNrZW5kIiwic3ViIjoidXNlciB0b2tlbiJ9.CunVLdzYet8Et46UL_2mxKKG4Zpz7eS9rZlctRV47rY")
	if err == nil {
		fmt.Println(users)
	} else {
		fmt.Println(err)
	}
}

func TestIncreaseFollowerCount(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	IncreaseFollowerCount(1)
	DecreaseFollowerCount(2)
}

func TestIncreaseFollowCount(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	IncreaseFollowCount(1)
	DecreaseFollowCount(2)
}
