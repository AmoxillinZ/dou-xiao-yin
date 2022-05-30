/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package service

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/mapper"
	"dou-xiao-yin/src/model"
	"fmt"
	"strconv"
	"time"
)

// GetVideoList ：传入userId用于判断当前用户是否给各个视频点了赞
func GetVideoList(userId int, latestTime string) ([]*Video, int64) {
	videos := make([]*Video, 0)
	videosOri := make([]*model.Video, 0)
	var nextTime int64
	// latest_time是否为空的两种情况，请求视频列表
	fmt.Println("^^latestTime = ", latestTime, "len(latestTime) = ", len(latestTime))
	if len(latestTime) == 0 {
		videosOri = mapper.GetVideos()
	} else { // 请求携带latest_time，查询latest_time之前的视频
		lt, _ := strconv.Atoi(latestTime)
		latestTimeUnix := int64(lt)
		videosOri = mapper.GetVideosByTime(latestTimeUnix)
	}
	if len(videosOri) <= 0 {
		return nil, time.Now().Unix()
	}
	//取本次返回的视频列表的最后一个的发布时间，并转换为时间戳赋值给nextTime
	nextTime = videosOri[len(videosOri)-1].PublishTime.UnixMilli()
	for _, videoOri := range videosOri {
		authorOri, err := mapper.GetUserById(videoOri.AuthorId)
		if err != nil {
			fmt.Println(err)
		}
		// model.User -> common.User
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
			Title:         videoOri.Title,
		}
		videos = append(videos, &video)
	}
	// 处理播放和封面路由
	parseUrl(videos)
	return videos, nextTime
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
