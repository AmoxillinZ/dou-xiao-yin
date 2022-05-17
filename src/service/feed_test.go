package service

import (
	"fmt"
	"testing"
)

func Test_getVideoList(t *testing.T) {
	list := GetVideoList()
	fmt.Println("len(list) = ", len(list))
	for _, video := range list {
		fmt.Printf("%v \n", video)
	}
}
