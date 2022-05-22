package mapper

import (
	"dou-xiao-yin/src/config"
	"testing"
)

func TestUpdateFavoriteCount(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	IncreaseFavoriteCount(2)
}
