package mapper

import (
	"dou-xiao-yin/src/config"
	"fmt"
	"testing"
)

func TestUpdateFavoriteCount(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	IncreaseFavoriteCount(2)
}

func TestGetAuthorIdByVideoId(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	uid, err := GetAuthorIdByVideoId(2)
	if err != nil {
		return
	}
	fmt.Println(uid)
}

func TestAddCommentCount(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	err := AddCommentCount(2)
	if err != nil {
		return
	}
}
