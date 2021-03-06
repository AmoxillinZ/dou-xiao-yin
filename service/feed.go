/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package service

import (
	"dou-xiao-yin/config"
	"dou-xiao-yin/json_model"
	mapper2 "dou-xiao-yin/model"
	"fmt"
	"strconv"
	"time"
)

// GetVideoList ：传入userId用于判断当前用户是否给各个视频点了赞
func GetVideoList(userId int, latestTime string) ([]*json_model.Video, int64) {
	videos := make([]*json_model.Video, 0)
	videosOri := make([]*mapper2.Video, 0)
	var nextTime int64
	// latest_time是否为空的两种情况，请求视频列表
	if len(latestTime) == 0 {
		videosOri = mapper2.GetVideos()
	} else { // 请求携带latest_time，查询latest_time之前的视频
		lt, _ := strconv.Atoi(latestTime)
		latestTimeUnix := int64(lt)
		videosOri = mapper2.GetVideosByTime(latestTimeUnix)
	}
	if len(videosOri) <= 0 {
		return nil, time.Now().Unix()
	}
	//取本次返回的视频列表的最后一个的发布时间，并转换为时间戳赋值给nextTime
	nextTime = videosOri[len(videosOri)-1].PublishTime.UnixMilli()
	for _, videoOri := range videosOri {
		authorOri, err := mapper2.GetUserById(videoOri.AuthorId)
		if err != nil {
			fmt.Println(err)
		}
		// model.User -> json_model.User
		author := json_model.User{
			Id:            authorOri.Id,
			Username:      authorOri.Username,
			FollowCount:   authorOri.FollowCount,
			FollowerCount: authorOri.FollowerCount,
			IsFollow:      mapper2.IsFollow(authorOri.Id, userId), //userId为当前登录的用户id
		}
		video := json_model.Video{
			Id:            videoOri.Id,
			Author:        author,
			FavoriteCount: videoOri.FavoriteCount,
			CommentCount:  videoOri.CommentCount,
			IsFavorite:    mapper2.IsFavorite(videoOri.Id, userId),
			Title:         videoOri.Title,
		}
		videos = append(videos, &video)
	}
	// 处理播放和封面路由
	parseUrl(videos)
	return videos, nextTime
}

/*为video生成视频和封面的静态资源路径*/
func parseUrl(videos []*json_model.Video) {
	ip := config.GetConf().Resource.Ip
	port := config.GetConf().Resource.Port
	for _, video := range videos {
		videoUrl := fmt.Sprintf("http://%s:%s/douyin/resources/video/%d/%d", ip, port, video.Author.Id, video.Id)
		coverUrl := fmt.Sprintf("http://%s:%s/douyin/resources/cover/%d/%d", ip, port, video.Author.Id, video.Id)
		video.PlayUrl = videoUrl
		video.CoverUrl = coverUrl
	}
}
