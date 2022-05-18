/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/17 20:15
 */

package utils

import (
	"fmt"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ip, _ := GetLocalIP()
	fmt.Println(ip)
}
