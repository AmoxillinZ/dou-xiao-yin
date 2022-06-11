/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 15:04
 */

package model

import (
	"dou-xiao-yin/config"
	"gorm.io/gorm"
)

func GetVideoById(id int) *Video {
	vs := &Video{}
	db := config.GetDefaultDb()
	db.Where("id = ?", id).Find(&vs)
	return vs
}

// GetVideos : latest_time缺省，返回最新视频
func GetVideos() []*Video {
	vs := make([]*Video, 0)
	db := config.GetDefaultDb()
	//按publish_time倒序返回
	db.Limit(30).Order("publish_time desc").Find(&vs)
	return vs
}

// GetVideosByTime : 限制投稿时间
func GetVideosByTime(latestTime int64) []*Video {
	vs := make([]*Video, 0)
	db := config.GetDefaultDb()
	//查找latest_time之前的视频，按publish_time倒序返回
	db.Where("publish_time < FROM_UNIXTIME(?)", latestTime/1000).Limit(30).Order("publish_time desc").Find(&vs)
	return vs
}

func GetVideosByAuthorId(authorId int) ([]*Video, error) {
	videos := make([]*Video, 0)
	db := config.GetDefaultDb()
	// 结果倒序返回
	result := db.Where("author_id = ?", authorId).Order("publish_time desc").Find(&videos)
	return videos, result.Error
}

func AddVideo(video *Video) error {
	db := config.GetDefaultDb()
	result := db.Create(&video)
	return result.Error
}

func GetAuthorIdByVideoId(videoId int) (int, error) {
	db := config.GetDefaultDb()
	video := &Video{Id: videoId}
	result := db.First(&video)
	return video.AuthorId, result.Error
}

func AddCommentCount(videoId int) error {
	db := config.GetDefaultDb()
	video := Video{Id: videoId}
	result := db.Model(&video).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	return result.Error
}

func DeleteCommentCount(videoId int) error {
	db := config.GetDefaultDb()
	video := Video{Id: videoId}
	result := db.Model(&video).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	return result.Error
}
