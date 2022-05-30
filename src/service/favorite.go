/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/22 20:39
 */

package service

import (
	"dou-xiao-yin/src/mapper"
	"errors"
)

func FavoriteAction(videoId int, userId int) error {
	// 先检验是否已经赞过
	if mapper.IsFavorite(videoId, userId) {
		return errors.New("已经点过赞")
	}
	return mapper.FavoriteAction(videoId, userId)
}

func UnFavoriteAction(videoId int, userId int) error {
	return mapper.UnFavoriteAction(videoId, userId)
}
