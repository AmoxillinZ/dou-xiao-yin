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
