/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:22
 */

package model

import (
	config2 "dou-xiao-yin/config"
	"fmt"
	"testing"
)

func TestAddToRelation(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	err := AddToRelation(2, 1)
	fmt.Println(err)
}

func TestDeleteFromRelation(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	DeleteFromRelation(2, 1)
}

func TestFollowList(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	ans := FollowList(6)
	for _, value := range ans {
		fmt.Println(value)
	}
}

func TestFollowerList(t *testing.T) {
	config2.InitConf("../config/conf.yaml")
	config2.InitDefaultDbEngine()
	ans := FollowerList(6)
	for _, value := range ans {
		fmt.Println(value)
	}
}
