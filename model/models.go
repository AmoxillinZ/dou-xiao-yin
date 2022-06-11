package model

import (
	"database/sql"
	"time"
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

type Favorite struct {
	Id         int          `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VideoId    int          `gorm:"column:video_id;type:int(11)" json:"video_id"`
	UserId     int          `gorm:"column:user_id;type:int(11)" json:"user_id"`
	CreateTime sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
}

func (m *Favorite) TableName() string {
	return "favorite"
}

type Relation struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(11)" json:"user_id"`       // 关注者id
	ToUserId   int       `gorm:"column:to_user_id;type:int(11)" json:"to_user_id"` // 被关注者id
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
}

func (m *Relation) TableName() string {
	return "relation"
}

type User struct {
	Id            int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Username      string `gorm:"column:username;type:varchar(255);uniqueIndex" json:"name"`
	Password      string `gorm:"column:password;type:varchar(255)" json:"password"`
	FollowCount   int    `gorm:"column:follow_count;type:int(11)" json:"follow_count"`
	FollowerCount int    `gorm:"column:follower_count;type:int(11)" json:"follower_count"`
	Token         string `gorm:"column:token;type:varchar(255)" json:"token"`
}

func (m *User) TableName() string {
	return "user"
}

type Video struct {
	Id            int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	FavoriteCount int       `gorm:"column:favorite_count;type:bigint(20)" json:"favorite_count"`
	CommentCount  int       `gorm:"column:comment_count;type:bigint(20)" json:"comment_count"`
	AuthorId      int       `gorm:"column:author_id;type:bigint(20)" json:"author_id"`
	PublishTime   time.Time `gorm:"column:publish_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"publish_time"`
	Title         string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
}

func (m *Video) TableName() string {
	return "video"
}
