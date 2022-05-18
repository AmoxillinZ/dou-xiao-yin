/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/12 10:02
 */

package model

type Video struct {
	Id int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	//PlayUrl       string `gorm:"column:play_url;type:varchar(255)" json:"play_url"`
	//CoverUrl      string `gorm:"column:cover_url;type:varchar(255)" json:"cover_url"`
	FavoriteCount int `gorm:"column:favorite_count;type:bigint(20)" json:"favorite_count"`
	CommentCount  int `gorm:"column:comment_count;type:bigint(20)" json:"comment_count"`
	AuthorId      int `gorm:"column:author_id;type:bigint(20)" json:"author_id"`
}

func (m *Video) TableName() string {
	return "video"
}
