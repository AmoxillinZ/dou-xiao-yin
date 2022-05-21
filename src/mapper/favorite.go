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
