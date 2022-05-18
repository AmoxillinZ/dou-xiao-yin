/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package service

import (
	"MyProject/src/mapper"
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
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func GetVideoList() []Video {
	videos := make([]Video, 0)
	// videos_ori：model.Video类型
	videos_ori := mapper.GetVideos()
	for _, video_ori := range videos_ori {
		// author_ori:model.User类型
		author_ori, err := mapper.GetUserById(video_ori.AuthorId)
		if err != nil {
			fmt.Println(err)
		}
		// model.User ->
		author := User{
			Id:            author_ori.Id,
			Name:          author_ori.Name,
			FollowCount:   author_ori.FollowerCount,
			FollowerCount: author_ori.FollowCount,
			IsFollow:      false, //待补充
		}
		video := Video{
			Id:            video_ori.Id,
			Author:        author,
			PlayUrl:       video_ori.PlayUrl,
			CoverUrl:      video_ori.CoverUrl,
			FavoriteCount: video_ori.FavoriteCount,
			CommentCount:  video_ori.CommentCount,
			IsFavorite:    false, //待补充
		}
		videos = append(videos, video)
	}
	return videos
}
