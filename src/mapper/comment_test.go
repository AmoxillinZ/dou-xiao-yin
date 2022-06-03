package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
	"fmt"
	"testing"
)

func TestGetCommentByIdComment(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	id := 1
	comment, err := GetCommentById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment)
}

func TestAddComment(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	newComment := &model.Comment{Content: "iuheria;lh"}
	id, err := AddComment(newComment)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func TestDeleteComment(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	comment := &model.Comment{Id: 2}
	err := DeleteComment(comment)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetCommentsByVideoId(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	comments, err := GetCommentsByVideoId(2)
	if err != nil {
		fmt.Print("err is ")
		fmt.Println(err)
	}
	for _, comment := range comments {
		fmt.Println(comment.Content)
	}
}
