package models

import (
	"github.com/jinzhu/gorm"
)

// TweetInfo Tweetから情報を取得してきたやつ
type TweetInfo struct {
	gorm.Model
	Images ImageInfos
}

// ImageInfo Tweet情報内のImage情報
type ImageInfo struct {
	gorm.Model
	TweetInfoID uint
	ImageURL    string `gorm:"not null"`
}

// ImageInfos List Of ImageInfo
type ImageInfos []ImageInfo
