/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 15:04
 */

package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
)

func GetVideoById(id int) *model.Video {
	vs := &model.Video{}
	db := config.GetDefaultDb()
	db.Where("id = ?", id).Find(&vs)
	return vs
}

func GetVideos() []*model.Video {
	vs := make([]*model.Video, 0)
	db := config.GetDefaultDb()
	//按publish_time倒序返回
	db.Limit(10).Order("publish_time desc").Find(&vs)
	return vs
}
