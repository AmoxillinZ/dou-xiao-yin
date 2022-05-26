/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 15:30
 */

package service

import (
	"dou-xiao-yin/src/config"
	"fmt"
	"testing"
)

func TestRelationAction(t *testing.T) {
	config.InitConf("../config/conf.yaml")
	config.InitDefaultDbEngine()
	err := RelationAction(6, 1, 2)
	fmt.Println(err)
}
