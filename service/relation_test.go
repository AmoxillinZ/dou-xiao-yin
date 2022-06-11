/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:30
 */

package service

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func TestRelationAction(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	err := RelationAction(6, 1, 2)
	fmt.Println(err)
}

func TestFollowList(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	list := FollowList(3, 6)
	for _, usr := range list {
		fmt.Println(usr)
	}
	fmt.Println(list)
}
