/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 20:32
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"fmt"
	"testing"
)

func TestGetUserIdByToken(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	users, err := GetUserIdByToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6InpoaCIsInBhc3N3b3JkIjoiMjMzMzMzIiwiZXhwIjoxNjUzNzQ3Mjk0LCJpYXQiOjE2NTMxNDI0OTQsImlzcyI6ImRvdS14aWFvLXlpbi1iYWNrZW5kIiwic3ViIjoidXNlciB0b2tlbiJ9.CunVLdzYet8Et46UL_2mxKKG4Zpz7eS9rZlctRV47rY")
	if err == nil {
		fmt.Println(users)
	} else {
		fmt.Println(err)
	}
}

func TestGetUserByToken(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	//users, err := GetUserByToken("wjpassword")
	users, err := GetUserByToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6InpoaCIsInBhc3N3b3JkIjoiMjMzMzMzIiwiZXhwIjoxNjUzNzQ3Mjk0LCJpYXQiOjE2NTMxNDI0OTQsImlzcyI6ImRvdS14aWFvLXlpbi1iYWNrZW5kIiwic3ViIjoidXNlciB0b2tlbiJ9.CunVLdzYet8Et46UL_2mxKKG4Zpz7eS9rZlctRV47rY")
	if err == nil {
		fmt.Println(users)
	} else {
		fmt.Println(err)
	}
}