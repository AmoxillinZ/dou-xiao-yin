/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/18 15:37
 */

package config

import (
	"fmt"
	"testing"
)

func TestGetConf(t *testing.T) {
	conf := GetConf()
	fmt.Println(conf.Database.Port)
}
