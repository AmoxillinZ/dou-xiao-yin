/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 21:47
 */

package model

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func Test_isFavorite(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	isFavorite := IsFavorite(13, 1)
	fmt.Println(isFavorite)
}

func TestFavoriteAction(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	UnFavoriteAction(1, 2)
	FavoriteAction(1, 2)
}

func TestFavoriteList(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	list, _ := FavoriteList(6)
	for _, v := range list {
		fmt.Println(v)
	}
}

func BenchmarkFavoriteList(b *testing.B) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	b.ResetTimer()
	FavoriteList(3)
	//for _, v := range list {
	//	fmt.Println(v)
	//}

}
