package models

import (
	_ "time"

	"github.com/jinzhu/gorm"
)

type TweetImage struct {
	gorm.Model
	Images ImageInfos
}

type ImageInfo struct {
	gorm.Model
	ImageURL string `gorm:"not null"`
}

type ImageInfos []ImageInfo
