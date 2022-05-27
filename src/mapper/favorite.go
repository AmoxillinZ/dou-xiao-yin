/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 21:34
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
	"gorm.io/gorm"
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

// FavoriteAction : 点赞原子操作：向点赞表中插入一条记录、video点赞数+1
// TODO 因为客户端不能正确区分是否点赞，所以不能正确请求action是点赞还是取消操作，因此可能会出现同一用户对同一视频多次点赞的情况
func FavoriteAction(videoId int, userId int) error {
	db := config.GetDefaultDb()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1.生成点赞记录
		favorite := &model.Favorite{VideoId: videoId, UserId: userId}
		if err := db.Create(&favorite).Error; err != nil {
			return err
		}
		// 2.video的点赞数+1
		video := model.Video{Id: videoId}
		if err := db.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UnFavoriteAction : 取消赞原子操作：向点赞表中删除相应记录、video点赞数-1
func UnFavoriteAction(videoId int, userId int) error {
	db := config.GetDefaultDb()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1.删除点赞记录
		if err := db.Where("video_id = ? and user_id = ?", videoId, userId).Delete(model.Favorite{}).Error; err != nil {
			return err
		}
		// 2.video的点赞数-1
		video := model.Video{Id: videoId}
		if err := db.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
