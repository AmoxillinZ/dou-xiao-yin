package model

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func TestGetCommentByIdComment(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	id := 1
	comment, err := GetCommentById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment)
}

func TestAddComment(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	newComment := &Comment{Content: "iuheria;lh"}
	id, err := AddComment(newComment)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func TestDeleteComment(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	comment := &Comment{Id: 2}
	err := DeleteComment(comment)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetCommentsByVideoId(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	comments, err := GetCommentsByVideoId(2)
	if err != nil {
		fmt.Print("err is ")
		fmt.Println(err)
	}
	for _, comment := range comments {
		fmt.Println(comment.Content)
	}
}
