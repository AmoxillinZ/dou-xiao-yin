/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 15:04
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
	"gorm.io/gorm"
)

func GetVideoById(id int) *model.Video {
	vs := &model.Video{}
	db := config.GetDefaultDb()
	db.Where("id = ?", id).Find(&vs)
	return vs
}

// GetVideos : latest_time缺省，返回最新视频
func GetVideos() []*model.Video {
	vs := make([]*model.Video, 0)
	db := config.GetDefaultDb()
	//按publish_time倒序返回
	db.Limit(30).Order("publish_time desc").Find(&vs)
	return vs
}

// GetVideosByTime : 限制投稿时间
func GetVideosByTime(latestTime int64) []*model.Video {
	vs := make([]*model.Video, 0)
	db := config.GetDefaultDb()
	//查找latest_time之前的视频，按publish_time倒序返回
	db.Where("publish_time < FROM_UNIXTIME(?)", latestTime/1000).Limit(30).Order("publish_time desc").Find(&vs)
	return vs
}

func GetVideosByAuthorId(authorId int) ([]*model.Video, error) {
	videos := make([]*model.Video, 0)
	db := config.GetDefaultDb()
	result := db.Where("author_id = ?", authorId).Find(&videos)
	return videos, result.Error
}

func AddVideo(video *model.Video) error {
	db := config.GetDefaultDb()
	result := db.Create(&video)
	return result.Error
}

// IncreaseFavoriteCount : 相应video的点赞数+1
func IncreaseFavoriteCount(videoId int) error {
	video := model.Video{Id: videoId}
	db := config.GetDefaultDb()
	result := db.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
	return result.Error
}

// DecreaseFavoriteCount : 相应video的点赞数-1
func DecreaseFavoriteCount(videoId int) error {
	video := model.Video{Id: videoId}
	db := config.GetDefaultDb()
	result := db.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
	return result.Error
}

func GetAuthorIdByVideoId(videoId int) (int, error) {
	db := config.GetDefaultDb()
	video := &model.Video{Id: videoId}
	result := db.First(&video)
	return video.AuthorId, result.Error
}
