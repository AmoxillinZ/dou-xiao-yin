/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 15:04
 */

package mapper

import (
	"MyProject/src/config"
	"MyProject/src/model"
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
	db.Limit(10).Find(&vs)
	return vs
}
