package models

import (
	_ "time"

	"github.com/jinzhu/gorm"
)

type TweetInfo struct {
	gorm.Model
	Images ImageInfos
}

type ImageInfo struct {
	gorm.Model
	TweetInfoID uint
	ImageURL    string `gorm:"not null"`
}

type ImageInfos []ImageInfo
