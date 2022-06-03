package service

import (
	"dou-xiao-yin/src/config"
	"fmt"
	"testing"
)

func TestAddComment(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	userId := 5
	videoId := 16
	content := "可爱狗狗"
	comment, err := AddComment(userId, videoId, content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment.Id)
}
