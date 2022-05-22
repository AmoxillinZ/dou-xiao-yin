/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/22 20:39
 */

package service

import "dou-xiao-yin/src/mapper"

func FavoriteAction(videoId int, userId int) error {
	// TODO 原子操作
	//1、相应视频的点赞数+1； 2、生成点赞记录
	if err := mapper.IncreaseFavoriteCount(videoId); err != nil {
		return err
	}
	if err := mapper.FavoriteAction(videoId, userId); err != nil {
		return err
	}
	return nil
}

func DisFavoriteAction(videoId int, userId int) error {
	// TODO 原子操作
	//1、相应视频的点赞数-1； 2、删除点赞记录
	if err := mapper.DecreaseFavoriteCount(videoId); err != nil {
		return err
	}
	if err := mapper.DisFavoriteAction(videoId, userId); err != nil {
		return err
	}
	return nil
}
