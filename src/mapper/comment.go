package mapper

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
	"time"
)

func AddComment(comment *model.Comment) (commentId int, err error) {

	db := config.GetDefaultDb()
	result := db.Create(&comment)

	return comment.Id, result.Error
}

func GetCommentById(id int) (*model.Comment, error) {

	comment := &model.Comment{}
	db := config.GetDefaultDb()
	result := db.Where("id = ?", id).Find(&comment)

	return comment, result.Error
}

func DeleteComment(comment *model.Comment) error {

	db := config.GetDefaultDb()
	result := db.Model(&comment).Update("delete_time", time.Now().Format("2006-01-02 15:04:05"))

	return result.Error
}

func GetCommentsByVideoId(videoId int) ([]*model.Comment, error) {

	comments := make([]*model.Comment, 0)
	db := config.GetDefaultDb()
	result := db.Where("video_id = ?", videoId).Order("create_time desc").Find(&comments)

	return comments, result.Error
}
