package model

import (
	"dou-xiao-yin/config"
	"time"
)

func AddComment(comment *Comment) (commentId int, err error) {

	db := config.GetDefaultDb()
	result := db.Create(&comment)

	return comment.Id, result.Error
}

func GetCommentById(id int) (*Comment, error) {

	comment := &Comment{}
	db := config.GetDefaultDb()
	result := db.Where("id = ?", id).Find(&comment)

	return comment, result.Error
}

func DeleteComment(comment *Comment) error {

	db := config.GetDefaultDb()
	result := db.Model(&comment).Update("delete_time", time.Now().Format("2006-01-02 15:04:05"))

	return result.Error
}

func GetCommentsByVideoId(videoId int) ([]*Comment, error) {

	comments := make([]*Comment, 0)
	db := config.GetDefaultDb()
	result := db.Where("video_id = ?", videoId).Order("create_time desc").Find(&comments)

	return comments, result.Error
}
