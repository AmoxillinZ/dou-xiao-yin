/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/21 21:32
 */

package model

import "database/sql"

type Favorite struct {
	Id         int          `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VideoId    int          `gorm:"column:video_id;type:int(11)" json:"video_id"`
	UserId     int          `gorm:"column:user_id;type:int(11)" json:"user_id"`
	CreateTime sql.NullTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
}

func (m *Favorite) TableName() string {
	return "favorite"
}
