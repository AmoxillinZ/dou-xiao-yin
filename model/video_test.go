package model

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func TestUpdateFavoriteCount(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	IncreaseFavoriteCount(2)
}

func TestGetAuthorIdByVideoId(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	uid, err := GetAuthorIdByVideoId(2)
	if err != nil {
		return
	}
	fmt.Println(uid)
}

func TestAddCommentCount(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	err := AddCommentCount(2)
	if err != nil {
		return
	}
}
