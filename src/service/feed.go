/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package service

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/mapper"
	"fmt"
)

func GetVideoList(userId int) []*Video {
	//传入userId用于判断当前用户是否给各个视频点了赞
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
			IsFavorite:    mapper.IsFavorite(videoOri.Id, userId),
		}
		videos = append(videos, &video)
	}
	// 处理播放和封面路由
	parseUrl(videos)
	return videos
}

/*为video生成视频和封面的静态资源路径*/
func parseUrl(videos []*Video) {
	ip := config.GetConf().Resource.Ip
	port := config.GetConf().Resource.Port
	for _, video := range videos {
		videoUrl := fmt.Sprintf("http://%s:%s/douyin/resources/video/%d/%d", ip, port, video.Author.Id, video.Id)
		coverUrl := fmt.Sprintf("http://%s:%s/douyin/resources/cover/%d/%d", ip, port, video.Author.Id, video.Id)
		video.PlayUrl = videoUrl
		video.CoverUrl = coverUrl
	}
}
