package service

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func TestAddComment(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	userId := 5
	videoId := 16
	content := "可爱狗狗"
	comment, err := AddComment(userId, videoId, content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment.Id)
}
