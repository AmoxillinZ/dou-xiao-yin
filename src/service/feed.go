/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package service

import (
	"dou-xiao-yin/src/mapper"
	"fmt"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int    `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int    `json:"favorite_count,omitempty"`
	CommentCount  int    `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int    `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int    `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func GetVideoList() []*Video {
	videos := make([]*Video, 0)
	// videos_ori：model.Video类型
	videosOri := mapper.GetVideos()
	for _, videoOri := range videosOri {
		// author_ori:model.User类型
		authorOri, err := mapper.GetUserById(videoOri.AuthorId)
		if err != nil {
			fmt.Println(err)
		}
		// model.User ->
		author := User{
			Id:            authorOri.Id,
			Username:      authorOri.Username,
			FollowCount:   authorOri.FollowerCount,
			FollowerCount: authorOri.FollowCount,
			IsFollow:      false, //待补充
		}
		video := Video{
			Id:            videoOri.Id,
			Author:        author,
			FavoriteCount: videoOri.FavoriteCount,
			CommentCount:  videoOri.CommentCount,
			IsFavorite:    false, //待补充
		}
		videos = append(videos, &video)
	}
	// 处理播放和封面路由
	parseUrl(videos)
	return videos
}

const (
	ip   = "172.20.6.84"
	port = 8080
)

func parseUrl(videos []*Video) {
	for _, video := range videos {
		videoUrl := fmt.Sprintf("http://%s:%d/douyin/resources/video/%d/%d", ip, port, video.Author.Id, video.Id)
		coverUrl := fmt.Sprintf("http://%s:%d/douyin/resources/cover/%d/%d", ip, port, video.Author.Id, video.Id)
		video.PlayUrl = videoUrl
		video.CoverUrl = coverUrl
	}

}
