package service

import (
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/mapper"
	"dou-xiao-yin/src/model"
	"dou-xiao-yin/src/utils"
	"errors"
	"fmt"
	"mime/multipart"
)

func PublishList(userId int, token string) ([]*json_model.Video, error) {
	// 用户权限的验证
	err := TokenVerify(userId, token)
	if err != nil {
		return nil, err
	}
	// 根据用户找到其发布的视频列表
	videoVos := make([]*json_model.Video, 0)
	videos, err := mapper.GetVideosByAuthorId(userId)
	if err != nil {
		return nil, errors.New("获取信息失败")
	}
	author, err := mapper.GetUserById(userId)
	if err != nil {
		return nil, errors.New("获取信息失败")
	}
	for _, video := range videos {
		author := json_model.User{
			Id:            author.Id,
			Username:      author.Username,
			FollowCount:   author.FollowerCount,
			FollowerCount: author.FollowCount,
			IsFollow:      mapper.IsFollow(author.Id, userId), //userId为当前登录的用户id
		}
		video := json_model.Video{
			Id:            video.Id,
			Author:        author,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    mapper.IsFavorite(video.Id, userId),
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
	user, err := mapper.GetUserByToken(token)
	if err != nil {
		fmt.Println(token)
		return errors.New("token 失效")
	}

	// 向数据库写入视频信息
	// TODO: implement 这个地方和下一步应该写成一个原子的操作
	video := &model.Video{FavoriteCount: 0, CommentCount: 0, AuthorId: user.Id, Title: title}
	err = mapper.AddVideo(video)
	if err != nil {
		return errors.New("保存视频到数据库失败")
	}

	// 上传视频文件到 oss
	// TODO: implement 添加上传封面的功能
	err = utils.UploadVideo(file, video)
	if err != nil {
		return err
	}
	return nil
}
