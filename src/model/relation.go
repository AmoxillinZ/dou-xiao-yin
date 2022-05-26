/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 11:43
 */

package model

import (
	"time"
)

type Relation struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(11)" json:"user_id"`       // 关注者id
	ToUserId   int       `gorm:"column:to_user_id;type:int(11)" json:"to_user_id"` // 被关注者id
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
}

func (m *Relation) TableName() string {
	return "relation"
}
