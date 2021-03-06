package service

import (
	"dou-xiao-yin/json_model"
	mapper2 "dou-xiao-yin/model"
	utils2 "dou-xiao-yin/utils"
	"errors"
	"fmt"
	"mime/multipart"
)

func PublishList(userId int, token string) ([]*json_model.Video, error) {
	// 获得当前登录用户
	loginId, _ := utils2.GetIdFromToken(token)
	// 根据用户找到其发布的视频列表
	videoVos := make([]*json_model.Video, 0)
	videos, err := mapper2.GetVideosByAuthorId(userId)
	if err != nil {
		return nil, errors.New("获取信息失败")
	}
	author, err := mapper2.GetUserById(userId)
	if err != nil {
		return nil, errors.New("获取信息失败")
	}
	for _, video := range videos {
		author := json_model.User{
			Id:            author.Id,
			Username:      author.Username,
			FollowCount:   author.FollowerCount,
			FollowerCount: author.FollowCount,
			IsFollow:      mapper2.IsFollow(author.Id, loginId), //判断当前登录用户是否关注了author
		}
		video := json_model.Video{
			Id:            video.Id,
			Author:        author,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    mapper2.IsFavorite(video.Id, loginId), //判断当前登录用户是否给该video点赞
			Title:         video.Title,
		}
		videoVos = append(videoVos, &video)
	}
	parseUrl(videoVos)
	// 返回视频列表
	return videoVos, nil
}

func PublishVideo(file *multipart.FileHeader, token string, title string) error {
	// 用户鉴权, 这里只有一个 token 该怎么鉴权呢? 先写成根据 token 查询用户
	user, err := mapper2.GetUserByToken(token)
	if err != nil {
		fmt.Println(token)
		return errors.New("token 失效")
	}

	// 向数据库写入视频信息
	// TODO: implement 这个地方和下一步应该写成一个原子的操作
	video := &mapper2.Video{FavoriteCount: 0, CommentCount: 0, AuthorId: user.Id, Title: title}
	err = mapper2.AddVideo(video)
	if err != nil {
		return errors.New("保存视频到数据库失败")
	}

	// 上传视频文件到 oss
	// TODO: implement 添加上传封面的功能
	var videoLink string
	err, videoLink = utils2.UploadVideo(file, video)
	videoLink = "http://dou-xiao-yin.oss-cn-beijing.aliyuncs.com:80/" + videoLink
	if err != nil {
		return err
	}
	// 获得刚刚上传的视频，截取封面并上传封面
	utils2.UploadCover(videoLink)
	return nil
}
