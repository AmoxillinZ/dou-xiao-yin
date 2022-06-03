package model

import (
	"database/sql"
)

type Comment struct {
	Id         int          `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VideoId    int          `gorm:"column:video_id;type:int(11)" json:"video_id"`
	UserId     int          `gorm:"column:user_id;type:int(11)" json:"user_id"`
	Content    string       `gorm:"column:content;type:text" json:"content"`
	CreateTime sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	DeleteTime sql.NullTime `gorm:"column:delete_time;type:timestamp;" json:"delete_time"`
}

func (m *Comment) TableName() string {
	return "comment"
}
