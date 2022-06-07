/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/12 10:35
 */

package model

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
