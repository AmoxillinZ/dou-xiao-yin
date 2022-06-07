/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 21:47
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"fmt"
	"testing"
)

func Test_isFavorite(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	isFavorite := IsFavorite(13, 1)
	fmt.Println(isFavorite)
}

func TestFavoriteAction(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	UnFavoriteAction(1, 2)
	FavoriteAction(1, 2)
}

func TestFavoriteList(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	list, _ := FavoriteList(6)
	for _, v := range list {
		fmt.Println(v)
	}
}

func BenchmarkFavoriteList(b *testing.B) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	b.ResetTimer()
	FavoriteList(3)
	//for _, v := range list {
	//	fmt.Println(v)
	//}

}
