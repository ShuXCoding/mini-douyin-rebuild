// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	VideoID       int64     `gorm:"column:video_id;primaryKey;autoIncrement:true" json:"video_id"`            // 视频id
	ID            int64     `gorm:"column:id" json:"id"`                                                      // 用户id
	PlayURL       string    `gorm:"column:play_url" json:"play_url"`                                          // 视频URL
	CoverURL      string    `gorm:"column:cover_url" json:"cover_url"`                                        // 封面URL
	FavoriteCount int32     `gorm:"column:favorite_count" json:"favorite_count"`                              // 点赞总数
	CommentCount  int32     `gorm:"column:comment_count" json:"comment_count"`                                // 评论总数
	Title         string    `gorm:"column:title" json:"title"`                                                // 视频标题
	CreateDate    time.Time `gorm:"column:create_date;not null;default:CURRENT_TIMESTAMP" json:"create_date"` // 创建时间
	IsFavorite    int32     `gorm:"column:is_favorite;not null" json:"is_favorite"`                           // 是否点赞
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
