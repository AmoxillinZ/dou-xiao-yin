/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/22 20:39
 */

package service

import (
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/mapper"
	"dou-xiao-yin/src/model"
	"errors"
	"fmt"
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

func FavoriteList(userId int, loginId int) ([]*json_model.Video, error) {
	videos := make([]*json_model.Video, 0)
	videosOri := make([]*model.Video, 0)
	var err error
	videosOri, err = mapper.FavoriteList(userId)
	if len(videosOri) <= 0 { // 如果没有获取到点赞数据，直接返回空的数据
		return videos, nil
	}
	// 将model.Video赋值给json_model.Video,并计算对应的user信息
	for _, videoOri := range videosOri {
		authorOri, err := mapper.GetUserById(videoOri.AuthorId)
		if err != nil {
			fmt.Println(err)
		}
		// model.User -> json_model.User
		author := json_model.User{
			Id:            authorOri.Id,
			Username:      authorOri.Username,
			FollowCount:   authorOri.FollowCount,
			FollowerCount: authorOri.FollowerCount,
			IsFollow:      mapper.IsFollow(authorOri.Id, loginId),
			//loginId用于判断当前登录用户是否关注了其余用户
		}
		video := json_model.Video{
			Id:            videoOri.Id,
			Author:        author,
			FavoriteCount: videoOri.FavoriteCount,
			CommentCount:  videoOri.CommentCount,
			IsFavorite:    mapper.IsFavorite(videoOri.Id, userId),
			Title:         videoOri.Title,
		}
		videos = append(videos, &video)
	}
	// 处理播放和封面路由
	parseUrl(videos)
	return videos, err
}
