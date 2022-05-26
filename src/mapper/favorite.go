/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 21:34
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
)

// IsFavorite 判断某个用户是否给某个视频点赞了
func IsFavorite(videoId int, userId int) bool {
	// 如果userId == 0表示当前处于未登录状态，直接返回false
	if userId == 0 {
		return false
	}
	favoriteList := make([]*model.Favorite, 0)
	db := config.GetDefaultDb()
	db.Where("video_id = ? and user_id = ?", videoId, userId).Limit(1).Take(&favoriteList)
	if len(favoriteList) == 0 {
		return false
	} else {
		return true
	}
}

// FavoriteAction : 点赞操作，向点赞表中插入一条记录 */
// TODO 因为客户端不能正确区分是否点赞，所以不能正确请求action是点赞还是取消操作，因此可能会出现同一用户对同一视频多次点赞的情况
func FavoriteAction(videoId int, userId int) error {
	favorite := &model.Favorite{VideoId: videoId, UserId: userId}
	db := config.GetDefaultDb()
	result := db.Create(&favorite)
	return result.Error
}

// UnFavoriteAction : 取消点赞操作，在点赞表中删除相应记录 */
func UnFavoriteAction(videoId int, userId int) error {
	db := config.GetDefaultDb()
	result := db.Where("video_id = ? and user_id = ?", videoId, userId).Delete(model.Favorite{})
	return result.Error
}
