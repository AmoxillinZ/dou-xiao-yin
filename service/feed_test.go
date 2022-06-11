package service

import (
	"fmt"
	"testing"
)

func Test_getVideoList(t *testing.T) {
	list, nextTime := GetVideoList(1, "1653297826698")
	fmt.Println("len(list) = ", len(list))
	for _, video := range list {
		fmt.Printf("%v \n", video)
	}
	fmt.Println(nextTime)
}
