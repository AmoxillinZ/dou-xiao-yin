/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:22
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"fmt"
	"testing"
)

func TestAddToRelation(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	err := AddToRelation(2, 1)
	fmt.Println(err)
}

func TestDeleteFromRelation(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	DeleteFromRelation(2, 1)
}

func TestFollowList(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	ans := FollowList(6)
	for _, value := range ans {
		fmt.Println(value)
	}
}

func TestFollowerList(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	ans := FollowerList(6)
	for _, value := range ans {
		fmt.Println(value)
	}
}
